package handlers

import (
	"bytes"
	"math"
	"net/http"
	"sort"
	"strings"

	"landing/backend/ent"
	"landing/backend/ent/blog"
	"landing/backend/internal/ai/embeddings"
	"landing/backend/internal/config"
	"landing/backend/internal/db"

	"github.com/gofiber/fiber/v2"
	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/net/html"
)

// CreateBlogRequest is the payload for creating a blog.
// swagger:model
type CreateBlogRequest struct {
	Category string `json:"category"`
	Text     string `json:"text"`
	Path     string `json:"path"`
}

// sanitizeAndExtractBody takes an incoming HTML string, extracts the inner HTML of the <body>
// if present, and sanitizes it using a safe allowlist policy. This prevents scripts and unsafe
// attributes while allowing common content formatting for blog posts.
func sanitizeAndExtractBody(in string) string {
	trimmed := strings.TrimSpace(in)
	if trimmed == "" {
		return ""
	}

	// Attempt to parse and extract <body> inner HTML. If no <body> exists, use the input as-is.
	n, err := html.Parse(strings.NewReader(trimmed))
	var content string
	if err == nil && n != nil {
		// Find the first <body> node.
		var body *html.Node
		var f func(*html.Node)
		f = func(node *html.Node) {
			if node.Type == html.ElementNode && node.Data == "body" {
				body = node
				return
			}
			for c := node.FirstChild; c != nil && body == nil; c = c.NextSibling {
				f(c)
			}
		}
		f(n)

		if body != nil {
			var buf bytes.Buffer
			for c := body.FirstChild; c != nil; c = c.NextSibling {
				html.Render(&buf, c)
			}
			content = buf.String()
		}
	}
	if content == "" {
		content = trimmed
	}

	// Sanitize with a UGC policy (allows common formatting tags but strips scripts, iframes, etc.).
	p := bluemonday.UGCPolicy()
	// Allow structural/article-related elements commonly used in blog posts.
	p.AllowElements("article", "section", "figure", "figcaption", "footer", "time")
	// Allow useful attributes on common elements.
	p.AllowAttrs("class", "id").OnElements("div", "span", "p", "article", "section", "figure", "figcaption", "h1", "h2", "h3", "h4", "ul", "ol", "li")
	// Microdata attributes for SEO snippets if present in content
	p.AllowAttrs("itemprop", "itemscope", "itemtype").OnElements("article", "div", "span", "time")
	// Images: allow common safe attributes
	p.AllowAttrs("src", "alt", "title", "width", "height", "loading", "decoding").OnElements("img")
	// Links: keep defaults, ensure target/rel are permitted (UGCPolicy already handles href). Add target _blank for external links.
	p.AddTargetBlankToFullyQualifiedLinks(true)
	// If in the future you need to allow specific iframes (e.g., YouTube), whitelist here explicitly.
	return p.Sanitize(content)
}

// ListBlogsHandler returns blogs, optionally filtered by category via query param `category`.
// @Summary List blogs
// @Tags blogs
// @Produce json
// @Param category query string false "Filter by category"
// @Success 200 {array} ent.Blog
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /blogs [get]
func ListBlogsHandler(c *fiber.Ctx) error {
	client := db.ClientFromCtx(c)
	if client == nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "database client missing"})
	}

	category := c.Query("category")
	q := client.Blog.Query()
	if category != "" {
		q = q.Where(blog.CategoryEQ(category))
	}
	items, err := q.Order(ent.Asc(blog.FieldPath)).All(c.UserContext())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(items)
}

// GetBlogByPathHandler returns a single blog by its path param.
// @Summary Get blog by path
// @Tags blogs
// @Produce json
// @Param path path string true "Blog path"
// @Success 200 {object} ent.Blog
// @Failure 404 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /blogs/{path} [get]
func GetBlogByPathHandler(c *fiber.Ctx) error {
	client := db.ClientFromCtx(c)
	if client == nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "database client missing"})
	}

	p := c.Params("path")
	if p == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "missing path"})
	}
	item, err := client.Blog.Query().Where(blog.PathEQ(p)).Only(c.UserContext())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "blog not found"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Generate offline embedding for the requested blog and persist (handles legacy mismatched dims)
	cfg := config.Load()
	genA, genErr := embeddings.GenerateEmbedding(c.UserContext(), cfg, item.Text)
	var a []float32
	if genErr == nil && genA != nil && len(genA) > 0 {
		// Best-effort persist latest embedding
		if _, uerr := client.Blog.UpdateOneID(item.ID).SetEmbedding(genA).Save(c.UserContext()); uerr == nil {
			item.Embedding = genA
		}
		a = genA
	} else if len(item.Embedding) > 0 {
		// Fallback to stored embedding if generation failed for some reason
		a = item.Embedding
	}
	if len(a) == 0 {
		return c.JSON(fiber.Map{"blog": item, "similar": []any{}})
	}

	// Fetch candidates and compute cosine similarity in-app.
	others, err := client.Blog.Query().Where(blog.PathNEQ(p)).All(c.UserContext())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	type scored struct {
		b   *ent.Blog
		sim float64
	}
	scores := make([]scored, 0, len(others))
	for _, b := range others {
		// Prefer stored embedding; generate and persist if missing
		var be []float32
		if b.Embedding != nil && len(b.Embedding) == len(a) {
			be = b.Embedding
		} else {
			gen, berr := embeddings.GenerateEmbedding(c.UserContext(), cfg, b.Text)
			if berr != nil || gen == nil || len(gen) != len(a) {
				continue
			}
			// Best-effort persist, ignore errors
			_, _ = client.Blog.UpdateOneID(b.ID).SetEmbedding(gen).Save(c.UserContext())
			be = gen
		}
		s := cosine(a, be)
		if !math.IsNaN(s) && !math.IsInf(s, 0) {
			scores = append(scores, scored{b: b, sim: s})
		}
	}
	sort.Slice(scores, func(i, j int) bool { return scores[i].sim > scores[j].sim })
	// Take top 5
	n := 5
	if len(scores) < n {
		n = len(scores)
	}
	similar := make([]*ent.Blog, 0, n)
	for i := 0; i < n; i++ {
		similar = append(similar, scores[i].b)
	}
	return c.JSON(fiber.Map{"blog": item, "similar": similar})
}

// CreateBlogHandler creates a new blog post.
// @Summary Create a blog post
// @Tags blogs
// @Accept json
// @Produce json
// @Param data body CreateBlogRequest true "Blog payload"
// @Success 201 {object} ent.Blog
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /blogs [post]
func CreateBlogHandler(c *fiber.Ctx) error {
	client := db.ClientFromCtx(c)
	if client == nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "database client missing"})
	}

	var req CreateBlogRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid JSON body"})
	}
	// Normalize inputs.
	req.Category = strings.TrimSpace(req.Category)
	req.Path = strings.TrimSpace(req.Path)
	if req.Category == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "category is required"})
	}
	if req.Path == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "path is required"})
	}

	// Allow full HTML documents by extracting <body> and sanitizing content before storing.
	processed := sanitizeAndExtractBody(req.Text)

	// Generate offline embedding for the content (best-effort)
	var emb []float32
	if e, err := embeddings.GenerateEmbedding(c.UserContext(), config.Load(), processed); err == nil && e != nil {
		emb = e
	}

	builder := client.Blog.Create().
		SetCategory(req.Category).
		SetText(processed).
		SetPath(req.Path)
	if len(emb) > 0 {
		builder = builder.SetEmbedding(emb)
	}
	created, err := builder.Save(c.UserContext())
	if err != nil {
		if ent.IsConstraintError(err) {
			return c.Status(http.StatusConflict).JSON(fiber.Map{"error": "blog with this path already exists"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(created)
}

// cosine computes cosine similarity between two equal-length float32 vectors.
func cosine(a, b []float32) float64 {
	if len(a) == 0 || len(a) != len(b) {
		return 0
	}
	var dot float64
	var na float64
	var nb float64
	for i := 0; i < len(a); i++ {
		va := float64(a[i])
		vb := float64(b[i])
		dot += va * vb
		na += va * va
		nb += vb * vb
	}
	denom := math.Sqrt(na) * math.Sqrt(nb)
	if denom == 0 {
		return 0
	}
	return dot / denom
}

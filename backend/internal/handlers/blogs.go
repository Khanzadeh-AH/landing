package handlers

import (
	"net/http"

	"landing/backend/ent"
	"landing/backend/ent/blog"
	"landing/backend/internal/db"

	"github.com/gofiber/fiber/v2"
)

// ListBlogsHandler returns blogs, optionally filtered by category via query param `category`.
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
	return c.JSON(item)
}

// CreateBlogHandler creates a new blog post.
// Expects JSON body: {"category": string, "text": string, "path": string}
func CreateBlogHandler(c *fiber.Ctx) error {
	client := db.ClientFromCtx(c)
	if client == nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "database client missing"})
	}

	var req struct {
		Category string `json:"category"`
		Text     string `json:"text"`
		Path     string `json:"path"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid JSON body"})
	}
	if req.Category == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "category is required"})
	}
	if req.Path == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "path is required"})
	}

	created, err := client.Blog.Create().
		SetCategory(req.Category).
		SetText(req.Text).
		SetPath(req.Path).
		Save(c.UserContext())
	if err != nil {
		if ent.IsConstraintError(err) {
			return c.Status(http.StatusConflict).JSON(fiber.Map{"error": "blog with this path already exists"})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(created)
}

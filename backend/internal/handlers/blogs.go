package handlers

import (
	"net/http"

	"landing/backend/ent"
	"landing/backend/ent/blog"
	"landing/backend/internal/db"

	"github.com/gofiber/fiber/v2"
)

// CreateBlogRequest is the payload for creating a blog.
// swagger:model
type CreateBlogRequest struct {
	Category string `json:"category"`
	Text     string `json:"text"`
	Path     string `json:"path"`
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
	return c.JSON(item)
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

package db

import (
	"context"
	"log"
	"strings"
	"time"

	"landing/backend/ent"
	"landing/backend/internal/config"
	"landing/backend/internal/sanitize"
)

// SeedDev inserts or updates development seed data.
func SeedDev(ctx context.Context, client *ent.Client, cfg config.Config) {
	// Base sample items
	blogs := []struct {
		Category string
		Path     string
		Text     string
	}{
		{
			Category: "ai",
			Path:     "what-is-rag",
			Text:     `<h1>RAG چیست؟</h1><p>RAG (Retrieval-Augmented Generation) روشی برای غنی‌سازی پاسخ‌های مدل‌های زبانی با جست‌وجو در پایگاه دانش است.</p>`,
		},
		{
			Category: "dev",
			Path:     "go-fiber-ent-setup",
			Text:     `<h1>راه‌اندازی Go Fiber + Ent</h1><p>در این مقاله یک بک‌اند سریع با Fiber و Ent می‌سازیم.</p>`,
		},
		{
			Category: "news",
			Path:     "welcome-to-tehranbot",
			Text:     `<h1>خوش‌آمدید به تهران‌بات</h1><p>اخبار و مطالب تیم ما را اینجا دنبال کنید.</p>`,
		},
	}

	// Append the provided full-HTML sample article (with simple placeholder replacements)
	rep := strings.NewReplacer(
		"{SITE_NAME}", cfg.AppName,
		"{AUTHOR}", "تهران‌بات",
		"{AUTHOR_BIO}", "تیم هوش مصنوعی و نرم‌افزار تهران‌بات",
		"{PUBLISH_DATE}", time.Now().Format(time.RFC3339),
		"{PUBLISH_DATE_FORMATTED}", time.Now().Format("2006-01-02"),
		"{MODIFIED_DATE}", time.Now().Format(time.RFC3339),
		"{MODIFIED_DATE_FORMATTED}", time.Now().Format("2006-01-02"),
		"{READING_TIME}", "8",
	)
	blogs = append(blogs, struct {
		Category string
		Path     string
		Text     string
	}{
		Category: "ai",
		Path:     "ai-industry-fa",
		Text:     rep.Replace(SampleAIIndustryFA),
	})

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	inserted := 0
	for _, b := range blogs {
		// Sanitize/normalize HTML before storing
		safe := sanitize.SanitizeBlogHTML(b.Text)
		_, err := client.Blog.Create().
			SetCategory(b.Category).
			SetPath(b.Path).
			SetText(safe).
			Save(ctx)
		if err != nil {
			if ent.IsConstraintError(err) {
				// Already exists; skip
				continue
			}
			log.Printf("seed: create blog '%s' failed: %v", b.Path, err)
			continue
		}
		inserted++
	}
	if inserted > 0 {
		log.Printf("seed: inserted %d blogs", inserted)
	}
}

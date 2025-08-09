package db

import (
	"context"
	"log"
	"time"

	"landing/backend/ent"
)

// SeedDev inserts development seed data if none exists.
func SeedDev(ctx context.Context, client *ent.Client) {
	// Seed blogs
	count, err := client.Blog.Query().Count(ctx)
	if err != nil {
		log.Printf("seed: count blogs failed: %v", err)
		return
	}
	if count > 0 {
		return
	}

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

	batch := make([]*ent.BlogCreate, 0, len(blogs))
	for _, b := range blogs {
		batch = append(batch, client.Blog.Create().SetCategory(b.Category).SetPath(b.Path).SetText(b.Text))
	}
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if _, err := client.Blog.CreateBulk(batch...).Save(ctx); err != nil {
		log.Printf("seed: create blogs failed: %v", err)
		return
	}
	log.Printf("seed: inserted %d blogs", len(blogs))
}

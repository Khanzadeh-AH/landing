package sanitize

import (
	"bytes"
	"regexp"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/net/html"
)

// SanitizeBlogHTML extracts <body> inner HTML if present, then sanitizes using a
// safe allowlist suitable for blog content. It allows common structural and
// formatting tags while stripping scripts/unsafe content.
func SanitizeBlogHTML(in string) string {
	trimmed := strings.TrimSpace(in)
	if trimmed == "" {
		return ""
	}

	// Attempt to parse and extract <body> inner HTML. If no <body> exists, use input as-is.
	n, err := html.Parse(strings.NewReader(trimmed))
	var content string
	if err == nil && n != nil {
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

	p := bluemonday.UGCPolicy()
	p.AllowElements("article", "section", "figure", "figcaption", "footer", "time")
	p.AllowAttrs("class", "id").OnElements("div", "span", "p", "article", "section", "figure", "figcaption", "h1", "h2", "h3", "h4", "ul", "ol", "li")
	p.AllowAttrs("itemprop", "itemscope", "itemtype").OnElements("article", "div", "span", "time")
	p.AllowAttrs("src", "alt", "title", "width", "height", "loading", "decoding").OnElements("img")
	p.AddTargetBlankToFullyQualifiedLinks(true)
	// Allow JSON-LD scripts specifically (but no other scripts)
	p.AllowElements("script")
	p.AllowAttrs("type").Matching(regexp.MustCompile(`(?i)^application/ld\+json$`)).OnElements("script")
	return p.Sanitize(content)
}

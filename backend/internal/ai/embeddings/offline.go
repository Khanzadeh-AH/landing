package embeddings

import (
    "bytes"
    "context"
    "hash/fnv"
    "math"
    "strings"
    "unicode"

    "golang.org/x/net/html"
    "landing/backend/internal/config"
)

// GenerateEmbedding creates a simple offline embedding using the hashing trick.
// Steps:
// 1) Strip HTML to text
// 2) Tokenize on unicode letter/digit boundaries, lowercase
// 3) Hash tokens into a fixed-size bag-of-words vector (dimension D)
// 4) L2-normalize the vector
// Returns nil if no tokens are found.
func GenerateEmbedding(_ context.Context, _ config.Config, input string) ([]float32, error) {
    text := strings.TrimSpace(stripHTML(input))
    if text == "" {
        return nil, nil
    }
    const D = 256 // embedding dimension
    vec := make([]float64, D)

    // Simple tokenizer: accumulate letters/digits, split on others
    var tok strings.Builder
    flush := func() {
        if tok.Len() == 0 { return }
        t := strings.ToLower(tok.String())
        // Hash token into [0, D)
        idx := hashToBucket(t, D)
        vec[idx] += 1.0
        tok.Reset()
    }
    for _, r := range text {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            tok.WriteRune(r)
        } else {
            flush()
        }
    }
    flush()

    // L2 normalize
    var norm float64
    for i := range vec { norm += vec[i] * vec[i] }
    norm = math.Sqrt(norm)
    if norm == 0 {
        return nil, nil
    }
    out := make([]float32, D)
    for i := range vec {
        out[i] = float32(vec[i] / norm)
    }
    return out, nil
}

func stripHTML(in string) string {
    n, err := html.Parse(strings.NewReader(in))
    if err != nil || n == nil {
        // fallback: remove tags crudely
        return removeTags(in)
    }
    var buf bytes.Buffer
    var f func(*html.Node)
    f = func(node *html.Node) {
        if node.Type == html.TextNode {
            buf.WriteString(node.Data)
        }
        for c := node.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
    f(n)
    return buf.String()
}

func removeTags(s string) string {
    // lightweight tag stripper: remove substrings between '<' and '>'
    b := strings.Builder{}
    inTag := false
    for _, r := range s {
        switch r {
        case '<':
            inTag = true
        case '>':
            inTag = false
        default:
            if !inTag { b.WriteRune(r) }
        }
    }
    return b.String()
}

func hashToBucket(token string, mod int) int {
    h := fnv.New32a()
    _, _ = h.Write([]byte(token))
    return int(h.Sum32() % uint32(mod))
}

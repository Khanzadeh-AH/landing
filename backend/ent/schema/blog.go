package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Blog holds the schema definition for the Blog entity.
type Blog struct{ ent.Schema }

// Fields of the Blog.
func (Blog) Fields() []ent.Field {
	return []ent.Field{
		field.String("category").NotEmpty(),
		field.Text("text"),
		field.String("path").NotEmpty().Unique(),
		// Embedding stores a vector representation for similarity search (offline-generated).
        // Note: Nillable() is not supported for JSON in this Ent version; Optional() suffices.
        field.JSON("embedding", []float32{}).Optional(),
	}
}

// Edges of the Blog.
func (Blog) Edges() []ent.Edge { return nil }

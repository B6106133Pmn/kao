package schema

import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/field"
    "github.com/facebookincubator/ent/schema/edge"
)

// Nametitle holds the schema definition for the Nametitle entity.
type Nametitle struct {
	ent.Schema
}

// Fields of the Nametitle.
func (Nametitle) Fields() []ent.Field {
	return []ent.Field{
		field.String("NameTitle").NotEmpty(),
		
	}	
}

// Edges of the Nametitle.
func (Nametitle) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("doctor",Doctor.Type).StorageKey(edge.Column("doctor_id")),	
		}
}

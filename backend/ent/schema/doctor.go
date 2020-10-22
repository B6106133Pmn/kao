package schema

import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/field"
    "github.com/facebookincubator/ent/schema/edge"
)

// Doctor holds the schema definition for the Doctor entity.
type Doctor struct {
	ent.Schema
}

// Fields of the Doctor.
func (Doctor) Fields() []ent.Field {
	return []ent.Field{
		field.String("Email").NotEmpty(),
		field.String("Password").NotEmpty(),
		field.String("Name").NotEmpty(),
		field.String("Tel").NotEmpty(),
    }
}

// Edges of the Doctor.
func (Doctor) Edges() []ent.Edge {
	return []ent.Edge{
        edge.From("department",Department.Type).Ref("doctor").Unique(),   
        edge.From("degree",Degree.Type).Ref("doctor").Unique(),
        edge.From("nametitle",Nametitle.Type).Ref("doctor").Unique(),
	}
}

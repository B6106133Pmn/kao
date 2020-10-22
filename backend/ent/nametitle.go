// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/pmn-kao/app/ent/nametitle"
)

// Nametitle is the model entity for the Nametitle schema.
type Nametitle struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// NameTitle holds the value of the "NameTitle" field.
	NameTitle string `json:"NameTitle,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NametitleQuery when eager-loading is set.
	Edges NametitleEdges `json:"edges"`
}

// NametitleEdges holds the relations/edges for other nodes in the graph.
type NametitleEdges struct {
	// Doctor holds the value of the doctor edge.
	Doctor []*Doctor
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DoctorOrErr returns the Doctor value or an error if the edge
// was not loaded in eager-loading.
func (e NametitleEdges) DoctorOrErr() ([]*Doctor, error) {
	if e.loadedTypes[0] {
		return e.Doctor, nil
	}
	return nil, &NotLoadedError{edge: "doctor"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Nametitle) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // NameTitle
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Nametitle fields.
func (n *Nametitle) assignValues(values ...interface{}) error {
	if m, n := len(values), len(nametitle.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	n.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field NameTitle", values[0])
	} else if value.Valid {
		n.NameTitle = value.String
	}
	return nil
}

// QueryDoctor queries the doctor edge of the Nametitle.
func (n *Nametitle) QueryDoctor() *DoctorQuery {
	return (&NametitleClient{config: n.config}).QueryDoctor(n)
}

// Update returns a builder for updating this Nametitle.
// Note that, you need to call Nametitle.Unwrap() before calling this method, if this Nametitle
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Nametitle) Update() *NametitleUpdateOne {
	return (&NametitleClient{config: n.config}).UpdateOne(n)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (n *Nametitle) Unwrap() *Nametitle {
	tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Nametitle is not a transactional entity")
	}
	n.config.driver = tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Nametitle) String() string {
	var builder strings.Builder
	builder.WriteString("Nametitle(")
	builder.WriteString(fmt.Sprintf("id=%v", n.ID))
	builder.WriteString(", NameTitle=")
	builder.WriteString(n.NameTitle)
	builder.WriteByte(')')
	return builder.String()
}

// Nametitles is a parsable slice of Nametitle.
type Nametitles []*Nametitle

func (n Nametitles) config(cfg config) {
	for _i := range n {
		n[_i].config = cfg
	}
}
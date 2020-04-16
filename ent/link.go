// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/kcarretto/paragon/ent/link"
)

// Link is the model entity for the Link schema.
type Link struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Alias holds the value of the "Alias" field.
	Alias string `json:"Alias,omitempty"`
	// ExpirationTime holds the value of the "ExpirationTime" field.
	ExpirationTime time.Time `json:"ExpirationTime,omitempty"`
	// Clicks holds the value of the "Clicks" field.
	Clicks int `json:"Clicks,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LinkQuery when eager-loading is set.
	Edges struct {
		// File holds the value of the file edge.
		File *File
	} `json:"edges"`
	file_id *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Link) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Alias
		&sql.NullTime{},   // ExpirationTime
		&sql.NullInt64{},  // Clicks
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Link) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // file_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Link fields.
func (l *Link) assignValues(values ...interface{}) error {
	if m, n := len(values), len(link.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	l.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Alias", values[0])
	} else if value.Valid {
		l.Alias = value.String
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field ExpirationTime", values[1])
	} else if value.Valid {
		l.ExpirationTime = value.Time
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field Clicks", values[2])
	} else if value.Valid {
		l.Clicks = int(value.Int64)
	}
	values = values[3:]
	if len(values) == len(link.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field file_id", value)
		} else if value.Valid {
			l.file_id = new(int)
			*l.file_id = int(value.Int64)
		}
	}
	return nil
}

// QueryFile queries the file edge of the Link.
func (l *Link) QueryFile() *FileQuery {
	return (&LinkClient{l.config}).QueryFile(l)
}

// Update returns a builder for updating this Link.
// Note that, you need to call Link.Unwrap() before calling this method, if this Link
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Link) Update() *LinkUpdateOne {
	return (&LinkClient{l.config}).UpdateOne(l)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (l *Link) Unwrap() *Link {
	tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Link is not a transactional entity")
	}
	l.config.driver = tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Link) String() string {
	var builder strings.Builder
	builder.WriteString("Link(")
	builder.WriteString(fmt.Sprintf("id=%v", l.ID))
	builder.WriteString(", Alias=")
	builder.WriteString(l.Alias)
	builder.WriteString(", ExpirationTime=")
	builder.WriteString(l.ExpirationTime.Format(time.ANSIC))
	builder.WriteString(", Clicks=")
	builder.WriteString(fmt.Sprintf("%v", l.Clicks))
	builder.WriteByte(')')
	return builder.String()
}

// Links is a parsable slice of Link.
type Links []*Link

func (l Links) config(cfg config) {
	for _i := range l {
		l[_i].config = cfg
	}
}

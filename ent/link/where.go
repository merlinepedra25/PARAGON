// Code generated by entc, DO NOT EDIT.

package link

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/kcarretto/paragon/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	},
	)
}

// Alias applies equality check predicate on the "Alias" field. It's identical to AliasEQ.
func Alias(v string) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAlias), v))
	},
	)
}

// ExpirationTime applies equality check predicate on the "ExpirationTime" field. It's identical to ExpirationTimeEQ.
func ExpirationTime(v time.Time) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpirationTime), v))
	},
	)
}

// Clicks applies equality check predicate on the "Clicks" field. It's identical to ClicksEQ.
func Clicks(v int) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClicks), v))
	},
	)
}

// AliasEQ applies the EQ predicate on the "Alias" field.
func AliasEQ(v string) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAlias), v))
	},
	)
}

// AliasNEQ applies the NEQ predicate on the "Alias" field.
func AliasNEQ(v string) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAlias), v))
	},
	)
}

// AliasIn applies the In predicate on the "Alias" field.
func AliasIn(vs ...string) predicate.Link {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Link(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAlias), v...))
	},
	)
}

// AliasNotIn applies the NotIn predicate on the "Alias" field.
func AliasNotIn(vs ...string) predicate.Link {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Link(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAlias), v...))
	},
	)
}

// AliasGT applies the GT predicate on the "Alias" field.
func AliasGT(v string) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAlias), v))
	},
	)
}

// AliasGTE applies the GTE predicate on the "Alias" field.
func AliasGTE(v string) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAlias), v))
	},
	)
}

// AliasLT applies the LT predicate on the "Alias" field.
func AliasLT(v string) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAlias), v))
	},
	)
}

// AliasLTE applies the LTE predicate on the "Alias" field.
func AliasLTE(v string) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAlias), v))
	},
	)
}

// AliasContains applies the Contains predicate on the "Alias" field.
func AliasContains(v string) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAlias), v))
	},
	)
}

// AliasHasPrefix applies the HasPrefix predicate on the "Alias" field.
func AliasHasPrefix(v string) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAlias), v))
	},
	)
}

// AliasHasSuffix applies the HasSuffix predicate on the "Alias" field.
func AliasHasSuffix(v string) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAlias), v))
	},
	)
}

// AliasEqualFold applies the EqualFold predicate on the "Alias" field.
func AliasEqualFold(v string) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAlias), v))
	},
	)
}

// AliasContainsFold applies the ContainsFold predicate on the "Alias" field.
func AliasContainsFold(v string) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAlias), v))
	},
	)
}

// ExpirationTimeEQ applies the EQ predicate on the "ExpirationTime" field.
func ExpirationTimeEQ(v time.Time) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpirationTime), v))
	},
	)
}

// ExpirationTimeNEQ applies the NEQ predicate on the "ExpirationTime" field.
func ExpirationTimeNEQ(v time.Time) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExpirationTime), v))
	},
	)
}

// ExpirationTimeIn applies the In predicate on the "ExpirationTime" field.
func ExpirationTimeIn(vs ...time.Time) predicate.Link {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Link(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExpirationTime), v...))
	},
	)
}

// ExpirationTimeNotIn applies the NotIn predicate on the "ExpirationTime" field.
func ExpirationTimeNotIn(vs ...time.Time) predicate.Link {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Link(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExpirationTime), v...))
	},
	)
}

// ExpirationTimeGT applies the GT predicate on the "ExpirationTime" field.
func ExpirationTimeGT(v time.Time) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExpirationTime), v))
	},
	)
}

// ExpirationTimeGTE applies the GTE predicate on the "ExpirationTime" field.
func ExpirationTimeGTE(v time.Time) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExpirationTime), v))
	},
	)
}

// ExpirationTimeLT applies the LT predicate on the "ExpirationTime" field.
func ExpirationTimeLT(v time.Time) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExpirationTime), v))
	},
	)
}

// ExpirationTimeLTE applies the LTE predicate on the "ExpirationTime" field.
func ExpirationTimeLTE(v time.Time) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExpirationTime), v))
	},
	)
}

// ExpirationTimeIsNil applies the IsNil predicate on the "ExpirationTime" field.
func ExpirationTimeIsNil() predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldExpirationTime)))
	},
	)
}

// ExpirationTimeNotNil applies the NotNil predicate on the "ExpirationTime" field.
func ExpirationTimeNotNil() predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldExpirationTime)))
	},
	)
}

// ClicksEQ applies the EQ predicate on the "Clicks" field.
func ClicksEQ(v int) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClicks), v))
	},
	)
}

// ClicksNEQ applies the NEQ predicate on the "Clicks" field.
func ClicksNEQ(v int) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldClicks), v))
	},
	)
}

// ClicksIn applies the In predicate on the "Clicks" field.
func ClicksIn(vs ...int) predicate.Link {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Link(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldClicks), v...))
	},
	)
}

// ClicksNotIn applies the NotIn predicate on the "Clicks" field.
func ClicksNotIn(vs ...int) predicate.Link {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Link(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldClicks), v...))
	},
	)
}

// ClicksGT applies the GT predicate on the "Clicks" field.
func ClicksGT(v int) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldClicks), v))
	},
	)
}

// ClicksGTE applies the GTE predicate on the "Clicks" field.
func ClicksGTE(v int) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldClicks), v))
	},
	)
}

// ClicksLT applies the LT predicate on the "Clicks" field.
func ClicksLT(v int) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldClicks), v))
	},
	)
}

// ClicksLTE applies the LTE predicate on the "Clicks" field.
func ClicksLTE(v int) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldClicks), v))
	},
	)
}

// HasFile applies the HasEdge predicate on the "file" edge.
func HasFile() predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FileTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FileTable, FileColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	},
	)
}

// HasFileWith applies the HasEdge predicate on the "file" edge with a given conditions (other predicates).
func HasFileWith(preds ...predicate.File) predicate.Link {
	return predicate.Link(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FileInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FileTable, FileColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Link) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for _, p := range predicates {
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Link) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for i, p := range predicates {
				if i > 0 {
					s1.Or()
				}
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Link) predicate.Link {
	return predicate.Link(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}

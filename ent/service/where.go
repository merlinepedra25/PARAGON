// Code generated by entc, DO NOT EDIT.

package service

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/kcarretto/paragon/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
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
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
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
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "Name" field. It's identical to NameEQ.
func Name(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// PubKey applies equality check predicate on the "PubKey" field. It's identical to PubKeyEQ.
func PubKey(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPubKey), v))
	})
}

// Config applies equality check predicate on the "Config" field. It's identical to ConfigEQ.
func Config(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConfig), v))
	})
}

// IsActivated applies equality check predicate on the "IsActivated" field. It's identical to IsActivatedEQ.
func IsActivated(v bool) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsActivated), v))
	})
}

// NameEQ applies the EQ predicate on the "Name" field.
func NameEQ(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "Name" field.
func NameNEQ(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "Name" field.
func NameIn(vs ...string) predicate.Service {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Service(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "Name" field.
func NameNotIn(vs ...string) predicate.Service {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Service(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "Name" field.
func NameGT(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "Name" field.
func NameGTE(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "Name" field.
func NameLT(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "Name" field.
func NameLTE(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "Name" field.
func NameContains(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "Name" field.
func NameHasPrefix(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "Name" field.
func NameHasSuffix(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "Name" field.
func NameEqualFold(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "Name" field.
func NameContainsFold(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// PubKeyEQ applies the EQ predicate on the "PubKey" field.
func PubKeyEQ(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPubKey), v))
	})
}

// PubKeyNEQ applies the NEQ predicate on the "PubKey" field.
func PubKeyNEQ(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPubKey), v))
	})
}

// PubKeyIn applies the In predicate on the "PubKey" field.
func PubKeyIn(vs ...string) predicate.Service {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Service(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPubKey), v...))
	})
}

// PubKeyNotIn applies the NotIn predicate on the "PubKey" field.
func PubKeyNotIn(vs ...string) predicate.Service {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Service(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPubKey), v...))
	})
}

// PubKeyGT applies the GT predicate on the "PubKey" field.
func PubKeyGT(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPubKey), v))
	})
}

// PubKeyGTE applies the GTE predicate on the "PubKey" field.
func PubKeyGTE(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPubKey), v))
	})
}

// PubKeyLT applies the LT predicate on the "PubKey" field.
func PubKeyLT(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPubKey), v))
	})
}

// PubKeyLTE applies the LTE predicate on the "PubKey" field.
func PubKeyLTE(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPubKey), v))
	})
}

// PubKeyContains applies the Contains predicate on the "PubKey" field.
func PubKeyContains(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPubKey), v))
	})
}

// PubKeyHasPrefix applies the HasPrefix predicate on the "PubKey" field.
func PubKeyHasPrefix(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPubKey), v))
	})
}

// PubKeyHasSuffix applies the HasSuffix predicate on the "PubKey" field.
func PubKeyHasSuffix(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPubKey), v))
	})
}

// PubKeyEqualFold applies the EqualFold predicate on the "PubKey" field.
func PubKeyEqualFold(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPubKey), v))
	})
}

// PubKeyContainsFold applies the ContainsFold predicate on the "PubKey" field.
func PubKeyContainsFold(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPubKey), v))
	})
}

// ConfigEQ applies the EQ predicate on the "Config" field.
func ConfigEQ(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConfig), v))
	})
}

// ConfigNEQ applies the NEQ predicate on the "Config" field.
func ConfigNEQ(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldConfig), v))
	})
}

// ConfigIn applies the In predicate on the "Config" field.
func ConfigIn(vs ...string) predicate.Service {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Service(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldConfig), v...))
	})
}

// ConfigNotIn applies the NotIn predicate on the "Config" field.
func ConfigNotIn(vs ...string) predicate.Service {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Service(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldConfig), v...))
	})
}

// ConfigGT applies the GT predicate on the "Config" field.
func ConfigGT(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldConfig), v))
	})
}

// ConfigGTE applies the GTE predicate on the "Config" field.
func ConfigGTE(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldConfig), v))
	})
}

// ConfigLT applies the LT predicate on the "Config" field.
func ConfigLT(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldConfig), v))
	})
}

// ConfigLTE applies the LTE predicate on the "Config" field.
func ConfigLTE(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldConfig), v))
	})
}

// ConfigContains applies the Contains predicate on the "Config" field.
func ConfigContains(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldConfig), v))
	})
}

// ConfigHasPrefix applies the HasPrefix predicate on the "Config" field.
func ConfigHasPrefix(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldConfig), v))
	})
}

// ConfigHasSuffix applies the HasSuffix predicate on the "Config" field.
func ConfigHasSuffix(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldConfig), v))
	})
}

// ConfigEqualFold applies the EqualFold predicate on the "Config" field.
func ConfigEqualFold(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldConfig), v))
	})
}

// ConfigContainsFold applies the ContainsFold predicate on the "Config" field.
func ConfigContainsFold(v string) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldConfig), v))
	})
}

// IsActivatedEQ applies the EQ predicate on the "IsActivated" field.
func IsActivatedEQ(v bool) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsActivated), v))
	})
}

// IsActivatedNEQ applies the NEQ predicate on the "IsActivated" field.
func IsActivatedNEQ(v bool) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsActivated), v))
	})
}

// HasTag applies the HasEdge predicate on the "tag" edge.
func HasTag() predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TagTable, TagColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagWith applies the HasEdge predicate on the "tag" edge with a given conditions (other predicates).
func HasTagWith(preds ...predicate.Tag) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TagTable, TagColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEvents applies the HasEdge predicate on the "events" edge.
func HasEvents() predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EventsTable, EventsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventsWith applies the HasEdge predicate on the "events" edge with a given conditions (other predicates).
func HasEventsWith(preds ...predicate.Event) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EventsTable, EventsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Service) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Service) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Service) predicate.Service {
	return predicate.Service(func(s *sql.Selector) {
		p(s.Not())
	})
}

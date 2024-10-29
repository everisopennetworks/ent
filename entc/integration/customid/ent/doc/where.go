// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package doc

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/entc/integration/customid/ent/schema"
)

// ID filters vertices based on their ID field.
func ID(id schema.DocID) predicate.Doc {
	return predicate.Doc(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id schema.DocID) predicate.Doc {
	return predicate.Doc(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id schema.DocID) predicate.Doc {
	return predicate.Doc(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...schema.DocID) predicate.Doc {
	return predicate.Doc(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...schema.DocID) predicate.Doc {
	return predicate.Doc(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id schema.DocID) predicate.Doc {
	return predicate.Doc(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id schema.DocID) predicate.Doc {
	return predicate.Doc(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id schema.DocID) predicate.Doc {
	return predicate.Doc(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id schema.DocID) predicate.Doc {
	return predicate.Doc(sql.FieldLTE(FieldID, id))
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.Doc {
	return predicate.Doc(sql.FieldEQ(FieldText, v))
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.Doc {
	return predicate.Doc(sql.FieldEQ(FieldText, v))
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.Doc {
	return predicate.Doc(sql.FieldNEQ(FieldText, v))
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.Doc {
	return predicate.Doc(sql.FieldIn(FieldText, vs...))
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.Doc {
	return predicate.Doc(sql.FieldNotIn(FieldText, vs...))
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.Doc {
	return predicate.Doc(sql.FieldGT(FieldText, v))
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.Doc {
	return predicate.Doc(sql.FieldGTE(FieldText, v))
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.Doc {
	return predicate.Doc(sql.FieldLT(FieldText, v))
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.Doc {
	return predicate.Doc(sql.FieldLTE(FieldText, v))
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.Doc {
	return predicate.Doc(sql.FieldContains(FieldText, v))
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.Doc {
	return predicate.Doc(sql.FieldHasPrefix(FieldText, v))
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.Doc {
	return predicate.Doc(sql.FieldHasSuffix(FieldText, v))
}

// TextRegex applies the Regex predicate on the "text" field.
func TextRegex(v string) predicate.Doc {
	return predicate.Doc(sql.FieldRegex(FieldText, v))
}

// TextIsNil applies the IsNil predicate on the "text" field.
func TextIsNil() predicate.Doc {
	return predicate.Doc(sql.FieldIsNull(FieldText))
}

// TextNotNil applies the NotNil predicate on the "text" field.
func TextNotNil() predicate.Doc {
	return predicate.Doc(sql.FieldNotNull(FieldText))
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.Doc {
	return predicate.Doc(sql.FieldEqualFold(FieldText, v))
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.Doc {
	return predicate.Doc(sql.FieldContainsFold(FieldText, v))
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Doc {
	return predicate.Doc(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Doc) predicate.Doc {
	return predicate.Doc(func(s *sql.Selector) {
		step := newParentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.Doc {
	return predicate.Doc(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.Doc) predicate.Doc {
	return predicate.Doc(func(s *sql.Selector) {
		step := newChildrenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRelated applies the HasEdge predicate on the "related" edge.
func HasRelated() predicate.Doc {
	return predicate.Doc(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, RelatedTable, RelatedPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRelatedWith applies the HasEdge predicate on the "related" edge with a given conditions (other predicates).
func HasRelatedWith(preds ...predicate.Doc) predicate.Doc {
	return predicate.Doc(func(s *sql.Selector) {
		step := newRelatedStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Doc) predicate.Doc {
	return predicate.Doc(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Doc) predicate.Doc {
	return predicate.Doc(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Doc) predicate.Doc {
	return predicate.Doc(sql.NotPredicates(p))
}

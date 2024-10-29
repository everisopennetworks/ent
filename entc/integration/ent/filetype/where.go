// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package filetype

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.FileType {
	return predicate.FileType(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.FileType {
	return predicate.FileType(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.FileType {
	return predicate.FileType(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.FileType {
	return predicate.FileType(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.FileType {
	return predicate.FileType(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.FileType {
	return predicate.FileType(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.FileType {
	return predicate.FileType(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.FileType {
	return predicate.FileType(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.FileType {
	return predicate.FileType(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.FileType {
	return predicate.FileType(sql.FieldEQ(FieldName, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.FileType {
	return predicate.FileType(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.FileType {
	return predicate.FileType(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.FileType {
	return predicate.FileType(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.FileType {
	return predicate.FileType(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.FileType {
	return predicate.FileType(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.FileType {
	return predicate.FileType(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.FileType {
	return predicate.FileType(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.FileType {
	return predicate.FileType(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.FileType {
	return predicate.FileType(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.FileType {
	return predicate.FileType(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.FileType {
	return predicate.FileType(sql.FieldHasSuffix(FieldName, v))
}

// NameRegex applies the Regex predicate on the "name" field.
func NameRegex(v string) predicate.FileType {
	return predicate.FileType(sql.FieldRegex(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.FileType {
	return predicate.FileType(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.FileType {
	return predicate.FileType(sql.FieldContainsFold(FieldName, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.FileType {
	return predicate.FileType(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.FileType {
	return predicate.FileType(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.FileType {
	return predicate.FileType(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.FileType {
	return predicate.FileType(sql.FieldNotIn(FieldType, vs...))
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v State) predicate.FileType {
	return predicate.FileType(sql.FieldEQ(FieldState, v))
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v State) predicate.FileType {
	return predicate.FileType(sql.FieldNEQ(FieldState, v))
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...State) predicate.FileType {
	return predicate.FileType(sql.FieldIn(FieldState, vs...))
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...State) predicate.FileType {
	return predicate.FileType(sql.FieldNotIn(FieldState, vs...))
}

// HasFiles applies the HasEdge predicate on the "files" edge.
func HasFiles() predicate.FileType {
	return predicate.FileType(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FilesTable, FilesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFilesWith applies the HasEdge predicate on the "files" edge with a given conditions (other predicates).
func HasFilesWith(preds ...predicate.File) predicate.FileType {
	return predicate.FileType(func(s *sql.Selector) {
		step := newFilesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FileType) predicate.FileType {
	return predicate.FileType(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FileType) predicate.FileType {
	return predicate.FileType(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.FileType) predicate.FileType {
	return predicate.FileType(sql.NotPredicates(p))
}

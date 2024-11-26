// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package blog

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the blog type in the database.
	Label = "blog"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOid holds the string denoting the oid field in the database.
	FieldOid = "oid"
	// EdgeAdmins holds the string denoting the admins edge name in mutations.
	EdgeAdmins = "admins"
	// UserFieldID holds the string denoting the ID field of the User.
	UserFieldID = "oid"
	// Table holds the table name of the blog in the database.
	Table = "blogs"
	// AdminsTable is the table that holds the admins relation/edge.
	AdminsTable = "users"
	// AdminsInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AdminsInverseTable = "users"
	// AdminsColumn is the table column denoting the admins relation/edge.
	AdminsColumn = "blog_admins"
)

// Columns holds all SQL columns for blog fields.
var Columns = []string{
	FieldID,
	FieldOid,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Blog queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOid orders the results by the oid field.
func ByOid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOid, opts...).ToFunc()
}

// ByAdminsField orders the results by admins field.
func ByAdminsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAdminsStep(), sql.OrderByField(field, opts...))
	}
}

// ByAdminsCount orders the results by admins count.
func ByAdminsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAdminsStep(), opts...)
	}
}

// ByAdmins orders the results by admins terms.
func ByAdmins(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAdminsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAdminsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AdminsInverseTable, UserFieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AdminsTable, AdminsColumn),
	)
}

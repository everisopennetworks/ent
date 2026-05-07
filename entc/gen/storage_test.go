// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gen

import (
	"testing"

	"entgo.io/ent/schema/field"
	"github.com/stretchr/testify/require"
)

func TestSQLStorageJSONContainsFold(t *testing.T) {
	storage, err := NewStorage("sql")
	require.NoError(t, err)

	jsonField := &Field{Type: &field.TypeInfo{Type: field.TypeJSON}}
	require.Equal(t, []Op{ContainsFold}, storage.Ops(jsonField))
	require.Equal(t, "ContainsJSONFold", storage.FieldOpCode(jsonField, ContainsFold))

	stringField := &Field{Type: &field.TypeInfo{Type: field.TypeString}}
	require.Equal(t, []Op{EqualFold, ContainsFold}, storage.Ops(stringField))
	require.Equal(t, "ContainsFold", storage.FieldOpCode(stringField, ContainsFold))
}

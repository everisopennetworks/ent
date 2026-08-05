// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"testing"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

// Doris (and other MySQL-wire-compatible engines) return extra columns for
// "SHOW VARIABLES LIKE 'version'" (e.g. Default_Value, Changed) on top of
// the Variable_name/Value pair that vanilla MySQL/MariaDB return.
func TestMySQL_InitVersion_ExtraColumns(t *testing.T) {
	db, mk, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mk.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
		WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value", "Default_Value", "Changed"}).
			AddRow("version", "8.0.18", "8.0.18", "0"))

	d := &MySQL{Driver: sql.OpenDB(dialect.MySQL, db)}
	require.NoError(t, d.init(context.Background()))
	require.Equal(t, "8.0.18", d.version)
	require.NoError(t, mk.ExpectationsWereMet())
}

func TestMySQL_InitVersion_StandardColumns(t *testing.T) {
	db, mk, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mk.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
		WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).
			AddRow("version", "8.0.31"))

	d := &MySQL{Driver: sql.OpenDB(dialect.MySQL, db)}
	require.NoError(t, d.init(context.Background()))
	require.Equal(t, "8.0.31", d.version)
	require.NoError(t, mk.ExpectationsWereMet())
}

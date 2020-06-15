// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/entc/integration/ent/schema"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/fieldtype"
)

// FieldTypeCreate is the builder for creating a FieldType entity.
type FieldTypeCreate struct {
	config
	mutation *FieldTypeMutation
	hooks    []Hook
}

// SetInt sets the int field.
func (ftc *FieldTypeCreate) SetInt(i int) *FieldTypeCreate {
	ftc.mutation.SetInt(i)
	return ftc
}

// SetInt8 sets the int8 field.
func (ftc *FieldTypeCreate) SetInt8(i int8) *FieldTypeCreate {
	ftc.mutation.SetInt8(i)
	return ftc
}

// SetInt16 sets the int16 field.
func (ftc *FieldTypeCreate) SetInt16(i int16) *FieldTypeCreate {
	ftc.mutation.SetInt16(i)
	return ftc
}

// SetInt32 sets the int32 field.
func (ftc *FieldTypeCreate) SetInt32(i int32) *FieldTypeCreate {
	ftc.mutation.SetInt32(i)
	return ftc
}

// SetInt64 sets the int64 field.
func (ftc *FieldTypeCreate) SetInt64(i int64) *FieldTypeCreate {
	ftc.mutation.SetInt64(i)
	return ftc
}

// SetOptionalInt sets the optional_int field.
func (ftc *FieldTypeCreate) SetOptionalInt(i int) *FieldTypeCreate {
	ftc.mutation.SetOptionalInt(i)
	return ftc
}

// SetNillableOptionalInt sets the optional_int field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt(i *int) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt(*i)
	}
	return ftc
}

// SetOptionalInt8 sets the optional_int8 field.
func (ftc *FieldTypeCreate) SetOptionalInt8(i int8) *FieldTypeCreate {
	ftc.mutation.SetOptionalInt8(i)
	return ftc
}

// SetNillableOptionalInt8 sets the optional_int8 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt8(i *int8) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt8(*i)
	}
	return ftc
}

// SetOptionalInt16 sets the optional_int16 field.
func (ftc *FieldTypeCreate) SetOptionalInt16(i int16) *FieldTypeCreate {
	ftc.mutation.SetOptionalInt16(i)
	return ftc
}

// SetNillableOptionalInt16 sets the optional_int16 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt16(i *int16) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt16(*i)
	}
	return ftc
}

// SetOptionalInt32 sets the optional_int32 field.
func (ftc *FieldTypeCreate) SetOptionalInt32(i int32) *FieldTypeCreate {
	ftc.mutation.SetOptionalInt32(i)
	return ftc
}

// SetNillableOptionalInt32 sets the optional_int32 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt32(i *int32) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt32(*i)
	}
	return ftc
}

// SetOptionalInt64 sets the optional_int64 field.
func (ftc *FieldTypeCreate) SetOptionalInt64(i int64) *FieldTypeCreate {
	ftc.mutation.SetOptionalInt64(i)
	return ftc
}

// SetNillableOptionalInt64 sets the optional_int64 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt64(i *int64) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt64(*i)
	}
	return ftc
}

// SetNillableInt sets the nillable_int field.
func (ftc *FieldTypeCreate) SetNillableInt(i int) *FieldTypeCreate {
	ftc.mutation.SetNillableInt(i)
	return ftc
}

// SetNillableNillableInt sets the nillable_int field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt(i *int) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt(*i)
	}
	return ftc
}

// SetNillableInt8 sets the nillable_int8 field.
func (ftc *FieldTypeCreate) SetNillableInt8(i int8) *FieldTypeCreate {
	ftc.mutation.SetNillableInt8(i)
	return ftc
}

// SetNillableNillableInt8 sets the nillable_int8 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt8(i *int8) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt8(*i)
	}
	return ftc
}

// SetNillableInt16 sets the nillable_int16 field.
func (ftc *FieldTypeCreate) SetNillableInt16(i int16) *FieldTypeCreate {
	ftc.mutation.SetNillableInt16(i)
	return ftc
}

// SetNillableNillableInt16 sets the nillable_int16 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt16(i *int16) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt16(*i)
	}
	return ftc
}

// SetNillableInt32 sets the nillable_int32 field.
func (ftc *FieldTypeCreate) SetNillableInt32(i int32) *FieldTypeCreate {
	ftc.mutation.SetNillableInt32(i)
	return ftc
}

// SetNillableNillableInt32 sets the nillable_int32 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt32(i *int32) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt32(*i)
	}
	return ftc
}

// SetNillableInt64 sets the nillable_int64 field.
func (ftc *FieldTypeCreate) SetNillableInt64(i int64) *FieldTypeCreate {
	ftc.mutation.SetNillableInt64(i)
	return ftc
}

// SetNillableNillableInt64 sets the nillable_int64 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt64(i *int64) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt64(*i)
	}
	return ftc
}

// SetValidateOptionalInt32 sets the validate_optional_int32 field.
func (ftc *FieldTypeCreate) SetValidateOptionalInt32(i int32) *FieldTypeCreate {
	ftc.mutation.SetValidateOptionalInt32(i)
	return ftc
}

// SetNillableValidateOptionalInt32 sets the validate_optional_int32 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableValidateOptionalInt32(i *int32) *FieldTypeCreate {
	if i != nil {
		ftc.SetValidateOptionalInt32(*i)
	}
	return ftc
}

// SetOptionalUint sets the optional_uint field.
func (ftc *FieldTypeCreate) SetOptionalUint(u uint) *FieldTypeCreate {
	ftc.mutation.SetOptionalUint(u)
	return ftc
}

// SetNillableOptionalUint sets the optional_uint field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUint(u *uint) *FieldTypeCreate {
	if u != nil {
		ftc.SetOptionalUint(*u)
	}
	return ftc
}

// SetOptionalUint8 sets the optional_uint8 field.
func (ftc *FieldTypeCreate) SetOptionalUint8(u uint8) *FieldTypeCreate {
	ftc.mutation.SetOptionalUint8(u)
	return ftc
}

// SetNillableOptionalUint8 sets the optional_uint8 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUint8(u *uint8) *FieldTypeCreate {
	if u != nil {
		ftc.SetOptionalUint8(*u)
	}
	return ftc
}

// SetOptionalUint16 sets the optional_uint16 field.
func (ftc *FieldTypeCreate) SetOptionalUint16(u uint16) *FieldTypeCreate {
	ftc.mutation.SetOptionalUint16(u)
	return ftc
}

// SetNillableOptionalUint16 sets the optional_uint16 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUint16(u *uint16) *FieldTypeCreate {
	if u != nil {
		ftc.SetOptionalUint16(*u)
	}
	return ftc
}

// SetOptionalUint32 sets the optional_uint32 field.
func (ftc *FieldTypeCreate) SetOptionalUint32(u uint32) *FieldTypeCreate {
	ftc.mutation.SetOptionalUint32(u)
	return ftc
}

// SetNillableOptionalUint32 sets the optional_uint32 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUint32(u *uint32) *FieldTypeCreate {
	if u != nil {
		ftc.SetOptionalUint32(*u)
	}
	return ftc
}

// SetOptionalUint64 sets the optional_uint64 field.
func (ftc *FieldTypeCreate) SetOptionalUint64(u uint64) *FieldTypeCreate {
	ftc.mutation.SetOptionalUint64(u)
	return ftc
}

// SetNillableOptionalUint64 sets the optional_uint64 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalUint64(u *uint64) *FieldTypeCreate {
	if u != nil {
		ftc.SetOptionalUint64(*u)
	}
	return ftc
}

// SetState sets the state field.
func (ftc *FieldTypeCreate) SetState(f fieldtype.State) *FieldTypeCreate {
	ftc.mutation.SetState(f)
	return ftc
}

// SetNillableState sets the state field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableState(f *fieldtype.State) *FieldTypeCreate {
	if f != nil {
		ftc.SetState(*f)
	}
	return ftc
}

// SetOptionalFloat sets the optional_float field.
func (ftc *FieldTypeCreate) SetOptionalFloat(f float64) *FieldTypeCreate {
	ftc.mutation.SetOptionalFloat(f)
	return ftc
}

// SetNillableOptionalFloat sets the optional_float field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalFloat(f *float64) *FieldTypeCreate {
	if f != nil {
		ftc.SetOptionalFloat(*f)
	}
	return ftc
}

// SetOptionalFloat32 sets the optional_float32 field.
func (ftc *FieldTypeCreate) SetOptionalFloat32(f float32) *FieldTypeCreate {
	ftc.mutation.SetOptionalFloat32(f)
	return ftc
}

// SetNillableOptionalFloat32 sets the optional_float32 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalFloat32(f *float32) *FieldTypeCreate {
	if f != nil {
		ftc.SetOptionalFloat32(*f)
	}
	return ftc
}

// SetDatetime sets the datetime field.
func (ftc *FieldTypeCreate) SetDatetime(t time.Time) *FieldTypeCreate {
	ftc.mutation.SetDatetime(t)
	return ftc
}

// SetNillableDatetime sets the datetime field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableDatetime(t *time.Time) *FieldTypeCreate {
	if t != nil {
		ftc.SetDatetime(*t)
	}
	return ftc
}

// SetDecimal sets the decimal field.
func (ftc *FieldTypeCreate) SetDecimal(f float64) *FieldTypeCreate {
	ftc.mutation.SetDecimal(f)
	return ftc
}

// SetNillableDecimal sets the decimal field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableDecimal(f *float64) *FieldTypeCreate {
	if f != nil {
		ftc.SetDecimal(*f)
	}
	return ftc
}

// SetDir sets the dir field.
func (ftc *FieldTypeCreate) SetDir(h http.Dir) *FieldTypeCreate {
	ftc.mutation.SetDir(h)
	return ftc
}

// SetNillableDir sets the dir field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableDir(h *http.Dir) *FieldTypeCreate {
	if h != nil {
		ftc.SetDir(*h)
	}
	return ftc
}

// SetNdir sets the ndir field.
func (ftc *FieldTypeCreate) SetNdir(h http.Dir) *FieldTypeCreate {
	ftc.mutation.SetNdir(h)
	return ftc
}

// SetNillableNdir sets the ndir field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNdir(h *http.Dir) *FieldTypeCreate {
	if h != nil {
		ftc.SetNdir(*h)
	}
	return ftc
}

// SetStr sets the str field.
func (ftc *FieldTypeCreate) SetStr(ss sql.NullString) *FieldTypeCreate {
	ftc.mutation.SetStr(ss)
	return ftc
}

// SetNullStr sets the null_str field.
func (ftc *FieldTypeCreate) SetNullStr(ss sql.NullString) *FieldTypeCreate {
	ftc.mutation.SetNullStr(ss)
	return ftc
}

// SetLink sets the link field.
func (ftc *FieldTypeCreate) SetLink(s schema.Link) *FieldTypeCreate {
	ftc.mutation.SetLink(s)
	return ftc
}

// SetNullLink sets the null_link field.
func (ftc *FieldTypeCreate) SetNullLink(s schema.Link) *FieldTypeCreate {
	ftc.mutation.SetNullLink(s)
	return ftc
}

// SetActive sets the active field.
func (ftc *FieldTypeCreate) SetActive(s schema.Status) *FieldTypeCreate {
	ftc.mutation.SetActive(s)
	return ftc
}

// SetNillableActive sets the active field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableActive(s *schema.Status) *FieldTypeCreate {
	if s != nil {
		ftc.SetActive(*s)
	}
	return ftc
}

// SetNullActive sets the null_active field.
func (ftc *FieldTypeCreate) SetNullActive(s schema.Status) *FieldTypeCreate {
	ftc.mutation.SetNullActive(s)
	return ftc
}

// SetNillableNullActive sets the null_active field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNullActive(s *schema.Status) *FieldTypeCreate {
	if s != nil {
		ftc.SetNullActive(*s)
	}
	return ftc
}

// SetDeleted sets the deleted field.
func (ftc *FieldTypeCreate) SetDeleted(sb sql.NullBool) *FieldTypeCreate {
	ftc.mutation.SetDeleted(sb)
	return ftc
}

// SetDeletedAt sets the deleted_at field.
func (ftc *FieldTypeCreate) SetDeletedAt(st sql.NullTime) *FieldTypeCreate {
	ftc.mutation.SetDeletedAt(st)
	return ftc
}

// SetIP sets the ip field.
func (ftc *FieldTypeCreate) SetIP(n net.IP) *FieldTypeCreate {
	ftc.mutation.SetIP(n)
	return ftc
}

// SetNullInt64 sets the null_int64 field.
func (ftc *FieldTypeCreate) SetNullInt64(si sql.NullInt64) *FieldTypeCreate {
	ftc.mutation.SetNullInt64(si)
	return ftc
}

// SetSchemaInt sets the schema_int field.
func (ftc *FieldTypeCreate) SetSchemaInt(s schema.Int) *FieldTypeCreate {
	ftc.mutation.SetSchemaInt(s)
	return ftc
}

// SetNillableSchemaInt sets the schema_int field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableSchemaInt(s *schema.Int) *FieldTypeCreate {
	if s != nil {
		ftc.SetSchemaInt(*s)
	}
	return ftc
}

// SetSchemaInt8 sets the schema_int8 field.
func (ftc *FieldTypeCreate) SetSchemaInt8(s schema.Int8) *FieldTypeCreate {
	ftc.mutation.SetSchemaInt8(s)
	return ftc
}

// SetNillableSchemaInt8 sets the schema_int8 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableSchemaInt8(s *schema.Int8) *FieldTypeCreate {
	if s != nil {
		ftc.SetSchemaInt8(*s)
	}
	return ftc
}

// SetSchemaInt64 sets the schema_int64 field.
func (ftc *FieldTypeCreate) SetSchemaInt64(s schema.Int64) *FieldTypeCreate {
	ftc.mutation.SetSchemaInt64(s)
	return ftc
}

// SetNillableSchemaInt64 sets the schema_int64 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableSchemaInt64(s *schema.Int64) *FieldTypeCreate {
	if s != nil {
		ftc.SetSchemaInt64(*s)
	}
	return ftc
}

// SetSchemaFloat sets the schema_float field.
func (ftc *FieldTypeCreate) SetSchemaFloat(s schema.Float64) *FieldTypeCreate {
	ftc.mutation.SetSchemaFloat(s)
	return ftc
}

// SetNillableSchemaFloat sets the schema_float field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableSchemaFloat(s *schema.Float64) *FieldTypeCreate {
	if s != nil {
		ftc.SetSchemaFloat(*s)
	}
	return ftc
}

// SetSchemaFloat32 sets the schema_float32 field.
func (ftc *FieldTypeCreate) SetSchemaFloat32(s schema.Float32) *FieldTypeCreate {
	ftc.mutation.SetSchemaFloat32(s)
	return ftc
}

// SetNillableSchemaFloat32 sets the schema_float32 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableSchemaFloat32(s *schema.Float32) *FieldTypeCreate {
	if s != nil {
		ftc.SetSchemaFloat32(*s)
	}
	return ftc
}

// SetNullFloat sets the null_float field.
func (ftc *FieldTypeCreate) SetNullFloat(sf sql.NullFloat64) *FieldTypeCreate {
	ftc.mutation.SetNullFloat(sf)
	return ftc
}

// Mutation returns the FieldTypeMutation object of the builder.
func (ftc *FieldTypeCreate) Mutation() *FieldTypeMutation {
	return ftc.mutation
}

// Save creates the FieldType in the database.
func (ftc *FieldTypeCreate) Save(ctx context.Context) (*FieldType, error) {
	if _, ok := ftc.mutation.Int(); !ok {
		return nil, errors.New("ent: missing required field \"int\"")
	}
	if _, ok := ftc.mutation.Int8(); !ok {
		return nil, errors.New("ent: missing required field \"int8\"")
	}
	if _, ok := ftc.mutation.Int16(); !ok {
		return nil, errors.New("ent: missing required field \"int16\"")
	}
	if _, ok := ftc.mutation.Int32(); !ok {
		return nil, errors.New("ent: missing required field \"int32\"")
	}
	if _, ok := ftc.mutation.Int64(); !ok {
		return nil, errors.New("ent: missing required field \"int64\"")
	}
	if v, ok := ftc.mutation.ValidateOptionalInt32(); ok {
		if err := fieldtype.ValidateOptionalInt32Validator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"validate_optional_int32\": %w", err)
		}
	}
	if v, ok := ftc.mutation.State(); ok {
		if err := fieldtype.StateValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"state\": %w", err)
		}
	}
	var (
		err  error
		node *FieldType
	)
	if len(ftc.hooks) == 0 {
		node, err = ftc.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FieldTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ftc.mutation = mutation
			node, err = ftc.gremlinSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ftc.hooks) - 1; i >= 0; i-- {
			mut = ftc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ftc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ftc *FieldTypeCreate) SaveX(ctx context.Context) *FieldType {
	v, err := ftc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ftc *FieldTypeCreate) gremlinSave(ctx context.Context) (*FieldType, error) {
	res := &gremlin.Response{}
	query, bindings := ftc.gremlin().Query()
	if err := ftc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	ft := &FieldType{config: ftc.config}
	if err := ft.FromResponse(res); err != nil {
		return nil, err
	}
	return ft, nil
}

func (ftc *FieldTypeCreate) gremlin() *dsl.Traversal {
	v := g.AddV(fieldtype.Label)
	if value, ok := ftc.mutation.Int(); ok {
		v.Property(dsl.Single, fieldtype.FieldInt, value)
	}
	if value, ok := ftc.mutation.Int8(); ok {
		v.Property(dsl.Single, fieldtype.FieldInt8, value)
	}
	if value, ok := ftc.mutation.Int16(); ok {
		v.Property(dsl.Single, fieldtype.FieldInt16, value)
	}
	if value, ok := ftc.mutation.Int32(); ok {
		v.Property(dsl.Single, fieldtype.FieldInt32, value)
	}
	if value, ok := ftc.mutation.Int64(); ok {
		v.Property(dsl.Single, fieldtype.FieldInt64, value)
	}
	if value, ok := ftc.mutation.OptionalInt(); ok {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt, value)
	}
	if value, ok := ftc.mutation.OptionalInt8(); ok {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt8, value)
	}
	if value, ok := ftc.mutation.OptionalInt16(); ok {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt16, value)
	}
	if value, ok := ftc.mutation.OptionalInt32(); ok {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt32, value)
	}
	if value, ok := ftc.mutation.OptionalInt64(); ok {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt64, value)
	}
	if value, ok := ftc.mutation.NillableInt(); ok {
		v.Property(dsl.Single, fieldtype.FieldNillableInt, value)
	}
	if value, ok := ftc.mutation.NillableInt8(); ok {
		v.Property(dsl.Single, fieldtype.FieldNillableInt8, value)
	}
	if value, ok := ftc.mutation.NillableInt16(); ok {
		v.Property(dsl.Single, fieldtype.FieldNillableInt16, value)
	}
	if value, ok := ftc.mutation.NillableInt32(); ok {
		v.Property(dsl.Single, fieldtype.FieldNillableInt32, value)
	}
	if value, ok := ftc.mutation.NillableInt64(); ok {
		v.Property(dsl.Single, fieldtype.FieldNillableInt64, value)
	}
	if value, ok := ftc.mutation.ValidateOptionalInt32(); ok {
		v.Property(dsl.Single, fieldtype.FieldValidateOptionalInt32, value)
	}
	if value, ok := ftc.mutation.OptionalUint(); ok {
		v.Property(dsl.Single, fieldtype.FieldOptionalUint, value)
	}
	if value, ok := ftc.mutation.OptionalUint8(); ok {
		v.Property(dsl.Single, fieldtype.FieldOptionalUint8, value)
	}
	if value, ok := ftc.mutation.OptionalUint16(); ok {
		v.Property(dsl.Single, fieldtype.FieldOptionalUint16, value)
	}
	if value, ok := ftc.mutation.OptionalUint32(); ok {
		v.Property(dsl.Single, fieldtype.FieldOptionalUint32, value)
	}
	if value, ok := ftc.mutation.OptionalUint64(); ok {
		v.Property(dsl.Single, fieldtype.FieldOptionalUint64, value)
	}
	if value, ok := ftc.mutation.State(); ok {
		v.Property(dsl.Single, fieldtype.FieldState, value)
	}
	if value, ok := ftc.mutation.OptionalFloat(); ok {
		v.Property(dsl.Single, fieldtype.FieldOptionalFloat, value)
	}
	if value, ok := ftc.mutation.OptionalFloat32(); ok {
		v.Property(dsl.Single, fieldtype.FieldOptionalFloat32, value)
	}
	if value, ok := ftc.mutation.Datetime(); ok {
		v.Property(dsl.Single, fieldtype.FieldDatetime, value)
	}
	if value, ok := ftc.mutation.Decimal(); ok {
		v.Property(dsl.Single, fieldtype.FieldDecimal, value)
	}
	if value, ok := ftc.mutation.Dir(); ok {
		v.Property(dsl.Single, fieldtype.FieldDir, value)
	}
	if value, ok := ftc.mutation.Ndir(); ok {
		v.Property(dsl.Single, fieldtype.FieldNdir, value)
	}
	if value, ok := ftc.mutation.Str(); ok {
		v.Property(dsl.Single, fieldtype.FieldStr, value)
	}
	if value, ok := ftc.mutation.NullStr(); ok {
		v.Property(dsl.Single, fieldtype.FieldNullStr, value)
	}
	if value, ok := ftc.mutation.Link(); ok {
		v.Property(dsl.Single, fieldtype.FieldLink, value)
	}
	if value, ok := ftc.mutation.NullLink(); ok {
		v.Property(dsl.Single, fieldtype.FieldNullLink, value)
	}
	if value, ok := ftc.mutation.Active(); ok {
		v.Property(dsl.Single, fieldtype.FieldActive, value)
	}
	if value, ok := ftc.mutation.NullActive(); ok {
		v.Property(dsl.Single, fieldtype.FieldNullActive, value)
	}
	if value, ok := ftc.mutation.Deleted(); ok {
		v.Property(dsl.Single, fieldtype.FieldDeleted, value)
	}
	if value, ok := ftc.mutation.DeletedAt(); ok {
		v.Property(dsl.Single, fieldtype.FieldDeletedAt, value)
	}
	if value, ok := ftc.mutation.IP(); ok {
		v.Property(dsl.Single, fieldtype.FieldIP, value)
	}
	if value, ok := ftc.mutation.NullInt64(); ok {
		v.Property(dsl.Single, fieldtype.FieldNullInt64, value)
	}
	if value, ok := ftc.mutation.SchemaInt(); ok {
		v.Property(dsl.Single, fieldtype.FieldSchemaInt, value)
	}
	if value, ok := ftc.mutation.SchemaInt8(); ok {
		v.Property(dsl.Single, fieldtype.FieldSchemaInt8, value)
	}
	if value, ok := ftc.mutation.SchemaInt64(); ok {
		v.Property(dsl.Single, fieldtype.FieldSchemaInt64, value)
	}
	if value, ok := ftc.mutation.SchemaFloat(); ok {
		v.Property(dsl.Single, fieldtype.FieldSchemaFloat, value)
	}
	if value, ok := ftc.mutation.SchemaFloat32(); ok {
		v.Property(dsl.Single, fieldtype.FieldSchemaFloat32, value)
	}
	if value, ok := ftc.mutation.NullFloat(); ok {
		v.Property(dsl.Single, fieldtype.FieldNullFloat, value)
	}
	return v.ValueMap(true)
}

// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package database_models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type vkgroupsTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *vkgroupsTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("vkgroups").
func (v *vkgroupsTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *vkgroupsTableType) Columns() []string {
	return []string{
		"id",
		"name",
		"category",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *vkgroupsTableType) NewStruct() reform.Struct {
	return new(Vkgroups)
}

// NewRecord makes a new record for that table.
func (v *vkgroupsTableType) NewRecord() reform.Record {
	return new(Vkgroups)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *vkgroupsTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// VkgroupsTable represents vkgroups view or table in SQL database.
var VkgroupsTable = &vkgroupsTableType{
	s: parse.StructInfo{
		Type:    "Vkgroups",
		SQLName: "vkgroups",
		Fields: []parse.FieldInfo{
			{Name: "ID", Type: "int32", Column: "id"},
			{Name: "Name", Type: "string", Column: "name"},
			{Name: "Category", Type: "*string", Column: "category"},
		},
		PKFieldIndex: 0,
	},
	z: new(Vkgroups).Values(),
}

// String returns a string representation of this struct or record.
func (s Vkgroups) String() string {
	res := make([]string, 3)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Name: " + reform.Inspect(s.Name, true)
	res[2] = "Category: " + reform.Inspect(s.Category, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Vkgroups) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Name,
		s.Category,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Vkgroups) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Name,
		&s.Category,
	}
}

// View returns View object for that struct.
func (s *Vkgroups) View() reform.View {
	return VkgroupsTable
}

// Table returns Table object for that record.
func (s *Vkgroups) Table() reform.Table {
	return VkgroupsTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Vkgroups) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Vkgroups) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Vkgroups) HasPK() bool {
	return s.ID != VkgroupsTable.z[VkgroupsTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.ID = pk.
func (s *Vkgroups) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = VkgroupsTable
	_ reform.Struct = (*Vkgroups)(nil)
	_ reform.Table  = VkgroupsTable
	_ reform.Record = (*Vkgroups)(nil)
	_ fmt.Stringer  = (*Vkgroups)(nil)
)

func init() {
	parse.AssertUpToDate(&VkgroupsTable.s, new(Vkgroups))
}
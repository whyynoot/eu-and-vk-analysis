// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package databaseModels

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type marksTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *marksTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("marks").
func (v *marksTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *marksTableType) Columns() []string {
	return []string{
		"student_id",
		"credit_1",
		"credit_2",
		"credit_3",
		"credit_4",
		"credit_5",
		"credit_6",
		"credit_7",
		"credit_8",
		"credit_9",
		"credit_10",
		"exam_1",
		"exam_2",
		"exam_3",
		"exam_4",
		"exam_5",
		"exam_6",
		"exam_7",
		"exam_8",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *marksTableType) NewStruct() reform.Struct {
	return new(Marks)
}

// NewRecord makes a new record for that table.
func (v *marksTableType) NewRecord() reform.Record {
	return new(Marks)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *marksTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// MarksTable represents marks view or table in SQL database.
var MarksTable = &marksTableType{
	s: parse.StructInfo{
		Type:    "Marks",
		SQLName: "marks",
		Fields: []parse.FieldInfo{
			{Name: "StudentID", Type: "int32", Column: "student_id"},
			{Name: "Credit1", Type: "*int32", Column: "credit_1"},
			{Name: "Credit2", Type: "*int32", Column: "credit_2"},
			{Name: "Credit3", Type: "*int32", Column: "credit_3"},
			{Name: "Credit4", Type: "*int32", Column: "credit_4"},
			{Name: "Credit5", Type: "*int32", Column: "credit_5"},
			{Name: "Credit6", Type: "*int32", Column: "credit_6"},
			{Name: "Credit7", Type: "*int32", Column: "credit_7"},
			{Name: "Credit8", Type: "*int32", Column: "credit_8"},
			{Name: "Credit9", Type: "*int32", Column: "credit_9"},
			{Name: "Credit10", Type: "*int32", Column: "credit_10"},
			{Name: "Exam1", Type: "*int32", Column: "exam_1"},
			{Name: "Exam2", Type: "*int32", Column: "exam_2"},
			{Name: "Exam3", Type: "*int32", Column: "exam_3"},
			{Name: "Exam4", Type: "*int32", Column: "exam_4"},
			{Name: "Exam5", Type: "*int32", Column: "exam_5"},
			{Name: "Exam6", Type: "*int32", Column: "exam_6"},
			{Name: "Exam7", Type: "*int32", Column: "exam_7"},
			{Name: "Exam8", Type: "*int32", Column: "exam_8"},
		},
		PKFieldIndex: 0,
	},
	z: new(Marks).Values(),
}

// String returns a string representation of this struct or record.
func (s Marks) String() string {
	res := make([]string, 19)
	res[0] = "StudentID: " + reform.Inspect(s.StudentID, true)
	res[1] = "Credit1: " + reform.Inspect(s.Credit1, true)
	res[2] = "Credit2: " + reform.Inspect(s.Credit2, true)
	res[3] = "Credit3: " + reform.Inspect(s.Credit3, true)
	res[4] = "Credit4: " + reform.Inspect(s.Credit4, true)
	res[5] = "Credit5: " + reform.Inspect(s.Credit5, true)
	res[6] = "Credit6: " + reform.Inspect(s.Credit6, true)
	res[7] = "Credit7: " + reform.Inspect(s.Credit7, true)
	res[8] = "Credit8: " + reform.Inspect(s.Credit8, true)
	res[9] = "Credit9: " + reform.Inspect(s.Credit9, true)
	res[10] = "Credit10: " + reform.Inspect(s.Credit10, true)
	res[11] = "Exam1: " + reform.Inspect(s.Exam1, true)
	res[12] = "Exam2: " + reform.Inspect(s.Exam2, true)
	res[13] = "Exam3: " + reform.Inspect(s.Exam3, true)
	res[14] = "Exam4: " + reform.Inspect(s.Exam4, true)
	res[15] = "Exam5: " + reform.Inspect(s.Exam5, true)
	res[16] = "Exam6: " + reform.Inspect(s.Exam6, true)
	res[17] = "Exam7: " + reform.Inspect(s.Exam7, true)
	res[18] = "Exam8: " + reform.Inspect(s.Exam8, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Marks) Values() []interface{} {
	return []interface{}{
		s.StudentID,
		s.Credit1,
		s.Credit2,
		s.Credit3,
		s.Credit4,
		s.Credit5,
		s.Credit6,
		s.Credit7,
		s.Credit8,
		s.Credit9,
		s.Credit10,
		s.Exam1,
		s.Exam2,
		s.Exam3,
		s.Exam4,
		s.Exam5,
		s.Exam6,
		s.Exam7,
		s.Exam8,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Marks) Pointers() []interface{} {
	return []interface{}{
		&s.StudentID,
		&s.Credit1,
		&s.Credit2,
		&s.Credit3,
		&s.Credit4,
		&s.Credit5,
		&s.Credit6,
		&s.Credit7,
		&s.Credit8,
		&s.Credit9,
		&s.Credit10,
		&s.Exam1,
		&s.Exam2,
		&s.Exam3,
		&s.Exam4,
		&s.Exam5,
		&s.Exam6,
		&s.Exam7,
		&s.Exam8,
	}
}

// View returns View object for that struct.
func (s *Marks) View() reform.View {
	return MarksTable
}

// Table returns Table object for that record.
func (s *Marks) Table() reform.Table {
	return MarksTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Marks) PKValue() interface{} {
	return s.StudentID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Marks) PKPointer() interface{} {
	return &s.StudentID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Marks) HasPK() bool {
	return s.StudentID != MarksTable.z[MarksTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.StudentID = pk.
func (s *Marks) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = MarksTable
	_ reform.Struct = (*Marks)(nil)
	_ reform.Table  = MarksTable
	_ reform.Record = (*Marks)(nil)
	_ fmt.Stringer  = (*Marks)(nil)
)

func init() {
	parse.AssertUpToDate(&MarksTable.s, new(Marks))
}

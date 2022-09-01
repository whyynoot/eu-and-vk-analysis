package databaseModels

//go:generate reform

// Groupsstudents represents a row in groupsstudents table.
//reform:groupsstudents
type Groupsstudents struct {
	GroupID   int32 `reform:"group_id"`
	StudentID int32 `reform:"student_id"`
}

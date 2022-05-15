package database_models

//go:generate reform

// Students represents a row in students table.
//reform:students
type Students struct {
	ID           int32   `reform:"id,pk"`
	Name         string  `reform:"name"`
	VkLink       *string `reform:"vk_link"`
	StudentGroup string  `reform:"student_group"`
}

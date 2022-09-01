package databaseModels

//go:generate reform

// Vkgroups represents a row in vkgroups table.
//reform:vkgroups
type Vkgroups struct {
	ID       int32   `reform:"id,pk"`
	Name     string  `reform:"name"`
	Category *string `reform:"category"`
}

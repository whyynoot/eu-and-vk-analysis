package models

// is id (vklink) of student needed ?
type Student struct {
	ID       int32
	Marks    Marks
	VKGroups []VKGroup
}

package models

// is id (vklink) of student needed ?
type Student struct {
	Id int32
	Marks Marks
	VKGroups []VKGroup
}
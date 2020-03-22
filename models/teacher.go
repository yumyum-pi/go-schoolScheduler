package models

// TeacherID is a unique identifier of the teacher
type TeacherID struct {
	Year   [4]byte `json:"yr"`
	JoinNo [4]byte `json:"jNo"`
}

// Teacher is a struct to store teacher data
type Teacher struct {
	Name             Name      `json:"name"`
	ID               TeacherID `json:"id"`
	Subject          Subject   `json:"subject"`
	Capacity         int       `json:"capacity"`
	FreePeriodPerDay int       `json:"freePeriodPerWeek"`
	Classes          []Class   `json:"classes"`
}

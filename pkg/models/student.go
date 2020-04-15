package models

// StudentID is a unique identifier of the student
type StudentID struct {
	Branch  [3]byte `json:"brcnh"`
	Year    [4]byte `json:"yr"`
	AdminNo [3]byte `json:"aNo"`
}

// Student struct hold's student data
type Student struct {
	ID      StudentID `json:"id"`      //Unique identifier for the student
	Name    Name      `json:"name"`    // Name object of the student
	ClassID ClassID   `json:"classID"` // Unique identifier the students class
}

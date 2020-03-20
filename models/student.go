package models

// Student struct hold's student data
type Student struct {
	// TODO make a proper ID just like classID
	ID      int     `json:"id"`      //Unique identifier for the student
	Name    Name    `json:"name"`    // Name object of the student
	ClassID ClassID `json:"classID"` // Unique identifier the students class
}

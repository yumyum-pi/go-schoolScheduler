package models

import "fmt"

// SubjectID is a unique identifier for a subject
type SubjectID struct {
	Standard [StandardBS]byte `json:"stnID"` // ID for the standard of the subject
	Type     [TypeBS]byte     `json:"type"`  // type of subject. Example English, Hindi
}

// Bytes return the byte value from Standard and Type values
func (id *SubjectID) Bytes() (b [SubjectIDBS]byte) {
	// loop through each byte
	for i := 0; i < SubjectIDBS; i++ {
		if i < StandardBS {
			// add standard info to the struct
			b[i] = (*id).Standard[i] // assign the bytes

		} else {
			// add type info to the struct
			j := i - StandardBS  // calculate the offest index of standard
			b[i] = (*id).Type[j] // assign the bytes
		}
	}
	return
}

// Init adds value to the SubjectID
func (id *SubjectID) Init(b [SubjectIDBS]byte) {
	// loop through each byte
	for i := 0; i < SubjectIDBS; i++ {
		if i < StandardBS {
			// add standard info to the byte
			(*id).Standard[i] = b[i] // assign the bytes
		} else {
			// add type info to the byte
			j := i - StandardBS  // calculate the offest index of standard
			(*id).Type[j] = b[i] // assign the bytes
		}
	}
}

// Subject is a struct to store subject data
type Subject struct {
	ID        SubjectID `json:"id"`         // unique identifier for a subject
	TeacherID TeacherID `json:"teacherID"`  // teacher Assigned
	Req       int       `json:"reqClasses"` // required classes per week
}

// Print writes class values to the console
func (s *Subject) Print() {
	fmt.Printf("sID=%v\tReq=%v\ttID=%v\n", s.ID.Bytes(), s.Req, s.TeacherID.Bytes())
}

// IsAssigned return bool of weather the teacher is assigned or not
func (s *Subject) IsAssigned() bool {
	// check is the teacherID and subjectID is assigned
	return (*s).TeacherID != (TeacherID{}) && (*s).ID != (SubjectID{})
}

// SubjectL is an array of Subject. It has the following methods:
// - FindByID
type SubjectL []Subject

// FindByID return's the index of the given id. It return's -1 if no id is found.
func (sl *SubjectL) FindByID(id [SubjectIDBS]byte) int {
	index := -1
	// loop through all the subjects in the list
	for i, s := range *sl {
		// match id
		if s.ID.Bytes() == id {
			return i // return the index of the matched id
		}
	}
	return index
}

// PrintSubjectL write the subjet list data to console
func PrintSubjectL(ss []Subject) {
	for _, s := range ss {
		s.Print()
	}
}

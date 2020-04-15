package models

// SubjectID is a unique identifier for a subject
type SubjectID struct {
	Standerd [StanderdBS]byte `json:"stnID"` // ID for the standerd of the subject
	Type     [TypeBS]byte     `json:"type"`  // type of subject. Example English, Hindi
}

// Bytes return the byte value from Standerd and Type values
func (id *SubjectID) Bytes() (b [SubjectIDBS]byte) {
	for i := 0; i < SubjectIDBS; i++ {
		if i < StanderdBS {
			b[i] = (*id).Standerd[i]
		} else {
			j := i - StanderdBS
			b[i] = (*id).Type[j]
		}
	}
	return
}

// Init adds value to the SubjectID
//TODO write test
func (id *SubjectID) Init(b [6]byte) {
	for i := 0; i < SubjectIDBS; i++ {
		if i < StanderdBS {
			(*id).Standerd[i] = b[i]
		} else {
			j := i - StanderdBS
			(*id).Type[j] = b[i]
		}
	}
}

// SubjectIDs is a slice to hold subject ids
type SubjectIDs []SubjectID

// SubTypeE is a struct to hold subject type and its frequency
type SubTypeE struct {
	Type [4]byte
	freq int
}

// SubType is a slice made up of SubTypeE
type SubType []SubTypeE

// Find return the index no. of the given subjectID
//TODO write test
func (s *SubType) Find(sub SubjectID) int {
	// loop to check each item in the slice
	for i, subTE := range *s {
		// check if ids are same
		if subTE.Type == sub.Type {
			return i // return if found a match
		}
	}
	return -1
}

// Add func add the subject id
//TODO write test
func (s *SubType) Add(sub SubjectID) {
	// check if the given id is in the slice
	i := (*s).Find(sub)

	// check if match is found
	if i != -1 {
		// match as found
		(*s)[i].freq++ // add to the frequency
	} else {
		// match not found
		(*s) = append((*s), SubTypeE{sub.Type, 1}) // add new element to the slice
	}

}

// Types return an slice of types of subject in the slice
//TODO write test
func (ids *SubjectIDs) Types() SubType {
	var types SubType
	for _, id := range *ids {
		types.Add(id)
	}

	return types
}

// Subject is a struct to store subject data
type Subject struct {
	ID         SubjectID `json:"id"`         // unique identifier for a subject
	ReqClasses int       `json:"reqClasses"` // required classes per week
	TeacherID  TeacherID `json:"teacherID"`  // teacher Assigned
}

// IsAssigned return bool of weather the teacher is assigned or not
//TODO write test
func (s *Subject) IsAssigned() bool {
	// check is the teacher is assigned
	return (*s).TeacherID != (TeacherID{})
}

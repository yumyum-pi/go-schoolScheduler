package models

const (
	// TeacherFreePeriod no of free period allowed in an weel
	TeacherFreePeriod = 5

	// TeacherCap calculate the max no of periods that a teacher has in a day
	TeacherCap = MaxCap - TeacherFreePeriod
)

// TeacherID is a unique identifier of the teacher
type TeacherID struct {
	Year   [YearBS]byte   `json:"yr"`  // Year at which the teacher joined
	JoinNo [JoinNoBS]byte `json:"jNo"` // the joining no.
}

// Bytes return combined bytes of the ID
func (id *TeacherID) Bytes() (b [TeacherIDBS]byte) {
	// loop through byte length
	for i := 0; i < TeacherIDBS; i++ {
		if i < YearBS {
			// add year info to the struct
			b[i] = (*id).Year[i] // assign the bytes
		} else {
			// add join no info to the struct
			j := i - YearBS        // calculate the offest index of join no
			b[i] = (*id).JoinNo[j] // assign the bytes
		}
	}
	return
}

// Init adds value to the TeacherID
func (id *TeacherID) Init(b [TeacherIDBS]byte) {
	// loop through byte length
	for i := 0; i < TeacherIDBS; i++ {
		if i < YearBS {
			// add year into to the byte
			(*id).Year[i] = b[i] // assign byte
		} else {
			// add join no to the byte
			j := i - YearBS        // calculate the offest index of join no
			(*id).JoinNo[j] = b[i] // assign byte
		}
	}
}

// ClassAssigned is a struct to store data of classed assigned to a teacher
type ClassAssigned struct {
	SubjectID SubjectID // Unique identifier for the subject
	ClassID   ClassID   // Unique identifier for the class
	Assigned  int       // No. of periods required
}

// Init assign the proper value of the struct
func (ca *ClassAssigned) Init(sID [SubjectIDBS]byte, cID [ClassIDBS]byte, a int) {
	(*ca).SubjectID.Init(sID) // assign subjectID
	(*ca).ClassID.Init(cID)   // assign classID
	(*ca).Assigned = a        // no. of assigned periods

}

// Teacher is a struct to store teacher data. Element includes
type Teacher struct {
	ID        TeacherID       `json:"id"`    // Unique identifier
	SubjectCT []SubjectID     `json:"subCT"` // List of subjects can teach
	AClassL   []ClassAssigned `json:"aCls"`  // Classes asssigned to the teacher
	Capacity  int             `json:"cap"`   // Max no. of periods free can take per week
}

// Init adds value to the teacher struct. ID and capacity
func (t *Teacher) Init(tID [TeacherIDBS]byte) {
	(*t).ID.Init(tID)          // assign teacherID
	(*t).Capacity = TeacherCap // set the capacity to Teacher cap
}

// CanTeach check if the teacher can teach the subject with the given subjectID
func (t *Teacher) CanTeach(sID SubjectID) bool {
	// loop through all the subject the teacher can teach
	for _, s := range (*t).SubjectCT {
		// match the given subjectID
		if s == sID {
			// found the subject
			// this teacher can teache the subject
			return true
		}
	}

	// not match found
	return false
}

// AssignClass will assign class to the teacher struct. It returns the diffrence between the capacity
// and the requirement. If the diff is < 0 then the teacher is not assigned. Paramerters :-
//  cID - ClassID of the class to teach
//  sID - StubjectID of the subject to teach
//  r - number to periods  required by the class
func (t *Teacher) AssignClass(cID ClassID, sID SubjectID, r int) (diff int) {
	// check if the teacher has the capacity to take the no. of periods required by the class
	diff = (*t).Capacity - r // calculate the difference

	// check if the teach has capacity to teach another class
	if diff >= 0 {
		// assign
		cAssign := ClassAssigned{sID, cID, r}
		(*t).Capacity = diff                         // reduce the current capacity
		(*t).AClassL = append((*t).AClassL, cAssign) // add the assign class to the list
	}
	return diff
}

// Teachers is a slice of Teacher with the following methords:
// -Add -FindIndex -FindBySubject
type Teachers []Teacher

// FindIndex return the index of the teacher
// of the given given ID
func (ts *Teachers) FindIndex(tID TeacherID) int {
	// Check if Teacher List is !empty
	if len(*ts) == 0 {
		return -1 // Element not found
	}

	// Loop through the slice and find the index
	for i, t := range *ts {
		if t.ID == tID {
			return i // Element found
		}
	}

	return -1 // Element not found
}

// FindBySub return a slice of teacherID with the subject given in parameter
func (ts *Teachers) FindBySub(sID SubjectID) (index []int) {
	// Check if the list is empty
	if len(*ts) == 0 {
		return // Element not found
	}

	// Loop through the slice and find the index
	for i, t := range *ts {
		// check if the techer can teacher this subject
		c := t.CanTeach(sID)
		if c {
			index = append(index, i) // Add the teacherID to the id slice
		}
	}

	return // Element not found
}

// FindBySubType return a slice of teacherID with the given subjectID
func (ts *Teachers) FindBySubType(sID SubjectID) (index []int) {
	// Check if the list is empty
	if len(*ts) == 0 {
		return // Element not found
	}

	// Loop through the slice and find the index
	for i, t := range *ts {
		// Loop through subejcts
		for _, s := range t.SubjectCT {
			// Check for a matching subjectID
			if s.Type == sID.Type {
				index = append(index, i) // Add the teacherID to the id slice
				break                    // exit the loop
			}
		}
	}
	return // Element not found
}

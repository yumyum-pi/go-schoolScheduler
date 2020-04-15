package models

// TeacherID is a unique identifier of the teacher
type TeacherID struct {
	Year   [YearBS]byte   `json:"yr"`
	JoinNo [JoinNoBS]byte `json:"jNo"`
}

// TeacherFreePeriod no of free period allowed in an weel
const TeacherFreePeriod = 5

// TeacherCap calculate the max no of periods that a teacher has in a day
const TeacherCap = MaxCap - TeacherFreePeriod

// Bytes return combined bytes of the ID
func (id *TeacherID) Bytes() (b [TeacherIDBS]byte) {
	for i := 0; i < TeacherIDBS; i++ {
		if i < YearBS {
			b[i] = (*id).Year[i]
		} else {
			j := i - YearBS
			b[i] = (*id).JoinNo[j]
		}
	}
	return
}

// Init adds value to the TeacherID
//TODO write test
func (id *TeacherID) Init(b [TeacherIDBS]byte) {
	for i := 0; i < TeacherIDBS; i++ {
		if i < YearBS {
			(*id).Year[i] = b[i]
		} else {
			j := i - YearBS
			(*id).JoinNo[j] = b[i]
		}
	}
}

// ClassAssigned is a struct to store data of a
// classed assigned to a teacher
type ClassAssigned struct {
	ClassID   ClassID   // Unique identifier for the class
	SubjectID SubjectID // Unique identifier for the subject
	Req       int       // No. of period required
}

// Teacher is a struct to store teacher data
type Teacher struct {
	Name            Name            `json:"name"`   // Name of the teacher
	ID              TeacherID       `json:"id"`     // Unique identifier of the teacher
	SubjectCT       []SubjectID     `json:"subCT"`  // Subjects and Standers that the teacher can teach
	Capacity        int             `json:"cap"`    // Max no. of periods the teacher can take per week
	ClassesAssigned []ClassAssigned `json:"cAsgnd"` // Classes asssigned to the teacher with the no. of period per week
}

// AssignClass will assign class to the teacher struct
func (t *Teacher) AssignClass(class ClassID, subID SubjectID, req int) (diff int) {
	diff = (*t).Capacity - req
	// check if the teach has capacity to teach another class
	if diff >= 0 {
		// assign
		cAssign := ClassAssigned{class, subID, req}
		(*t).Capacity = diff
		(*t).ClassesAssigned = append((*t).ClassesAssigned, cAssign)
	}
	return diff
}

// Teachers is a slice of Teacher with the following methords:
// -Add -FindIndex -FindBySubject
type Teachers []Teacher

// FindIndex return the index of the teacher
// of the given given ID
func (ts *Teachers) FindIndex(id *TeacherID) int {
	// Check if Teacher List is !empty
	if len(*ts) == 0 {
		return -1 // Element not found
	}

	// Loop through the slice and find the index
	for i, t := range *ts {
		if t.ID == *id {
			return i // Element found
		}
	}

	return -1 // Element not found
}

// FindBySub return a slice of teacherID with the subject given in parameter
func (ts *Teachers) FindBySub(subID *SubjectID) (index []int) {
	// Check if the list is empty
	if len(*ts) == 0 {
		return // Element not found
	}

	// Loop through the slice and find the index
	for i, t := range *ts {
		// Loop through subejcts
		for _, s := range t.SubjectCT {
			// Check for a matching subjectID
			if s == *subID {
				index = append(index, i) // Add the teacherID to the id slice
			}
		}
	}

	return // Element not found
}

// FindBySubType return a slice of teacherID with the subject given in parameter
func (ts *Teachers) FindBySubType(subID *SubjectID) (index []int) {
	// Check if the list is empty
	if len(*ts) == 0 {
		return // Element not found
	}

	// Loop through the slice and find the index
	for i, t := range *ts {
		// Loop through subejcts
		for _, s := range t.SubjectCT {
			// Check for a matching subjectID
			if s.Type == subID.Type {
				index = append(index, i) // Add the teacherID to the id slice
			}
		}
	}

	return // Element not found
}

// ClassAssigned returns the index of teach that has been assigned the given subject and class
func (ts *Teachers) ClassAssigned(subID *SubjectID, classID *ClassID) (int, int) {
	// loop through each teacher
	for i, t := range *ts {
		// check if the assigned class !empty
		if len(t.ClassesAssigned) != 0 {
			// loop though each classes assigned
			for j, ca := range t.ClassesAssigned {
				if subID == &ca.SubjectID && classID == &ca.ClassID {
					return i, j
				}
			}
		}

	}

	return -1, -1
}

//

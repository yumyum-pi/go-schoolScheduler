package models

import "github.com/yumyum-pi/go-schoolScheduler/utils"

// Class id has should have the following components
// - class	: (1 - 12)	= requires 2 byte of data
// - section: (1 - n)	= requires 2 byte of data
// - group	: (1 - n)	= requires 2 byte of data

// ClassID is struct to store class id
type ClassID struct {
	Standerd [2]byte `json:"stn"`
	Section  [2]byte `json:"sec"`
	Group    [2]byte `json:"grp"`
	Year     [4]byte `json:"year"`
}

// Bytes return combined bytes of the ID
func (i *ClassID) Bytes() [10]byte {
	return [10]byte{
		i.Year[0],
		i.Year[1],
		i.Year[2],
		i.Year[3],
		i.Standerd[0],
		i.Standerd[1],
		i.Section[0],
		i.Section[1],
		i.Group[0],
		i.Group[1],
	}
}

// Create return a new ClassID
func (i ClassID) Create(st, sec, grp [2]byte, yr [4]byte) ClassID {
	return ClassID{
		Standerd: st,
		Section:  sec,
		Group:    grp,
		Year:     yr,
	}
}

// Class in an struct to hold class data
type Class struct {
	ID          ClassID     `json:"id"`          // Unique identifier for a class
	StudentID   []StudentID `json:"student"`     // List of students in this class
	Subjects    []Subject   `json:"subjects"`    // List of subject to be studied in this class
	NFreePeriod int         `json:"nFreePeriod"` // Number of free period this class has
}

// Init initialize the class struct
func (c *Class) Init(id ClassID) {
	(*c).ID = id              // assign ID to class
	(*c).NFreePeriod = MaxCap // assign the default max cap to class
}

// AddSubject adds the subejct to the class
func (c *Class) AddSubject(sub Subject) {
	(*c).Subjects = append((*c).Subjects, sub) // add the new subject to the array
}

// AssignTeacher adds the subejct to the class
func (c *Class) AssignTeacher(subID SubjectID, tID TeacherID) {
	// find the subject
	for i := range (*c).Subjects {
		if (*c).Subjects[i].ID == subID {
			(*c).Subjects[i].TeacherID = tID
			(*c).NFreePeriod -= (*c).Subjects[i].ReqClasses // reduce the no. of remaining capacity by the no. of period required by the subject
		}
	}
}

// CalRemCap calculate and update the remaining capacity of the class
func (c *Class) CalRemCap() {
	(*c).NFreePeriod = MaxCap // assign the default max cap to class
	for _, s := range (*c).Subjects {
		// check if the teacherID is not !empty
		if s.IsAssigned() {
			(*c).NFreePeriod -= s.ReqClasses // reduce the no. of remaining capacity by the no. of period required by the subject
		}
	}
}

// Classes in an slice of class with extra function
type Classes []Class

// AssignTeachers assigns teacher to the class
// TODO Write test for this function
func (cs *Classes) AssignTeachers(t *Teachers) (emptySubs []SubjectID) {
	// loop through classes
	for _, cls := range *cs {
		// - loop through subjects
		for _, s := range cls.Subjects {
			// -- for each subject select a teacher
			tMatchIndex := (*t).FindBySub(&s.ID)

			// check if the matchs slice is !empty
			if len(tMatchIndex) != 0 {
				ifAssigned, checked := false, false

				// loop until assigned or check all element in tMatchIndex
				for !ifAssigned && !checked {
					// get the current length of tMatchIndex
					tMatchLength := len(tMatchIndex)

					// check if tMatchLength is 0
					if tMatchLength == 0 {
						// not remain matches are left
						// through error
						// to emptySubs
						emptySubs = append(emptySubs, s.ID)
						checked = true // exit the loop
					} else {
						// generate a random number
						i := utils.GenerateRandomInt(tMatchLength, 10)

						// -- assign the teacher to the subject of the class
						diff := (*t)[tMatchIndex[i]].AssignClass(cls.ID, s.ID, s.ReqClasses)
						// -- check capacity
						// check if diff is negative
						// means that the requested period was greater than the capacity

						if diff >= 0 {
							ifAssigned = true // exit the loop
							// -- assign the class  and subject to the teacher
							cls.AssignTeacher(s.ID, (*t)[tMatchIndex[i]].ID)
						} else {
							// delete the element from tMatchIndex
							tMatchIndex = append(tMatchIndex[:i], tMatchIndex[i+1:]...)
						}
					}
				}

			} else {
				// through error
				// to emptySubs
				emptySubs = append(emptySubs, s.ID)
			}
		}
	}

	return
}

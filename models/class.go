package models

import (
	"github.com/yumyum-pi/go-schoolScheduler/utils"
)

// ClassID is struct to store class id
type ClassID struct {
	Year     [YearBS]byte     `json:"yr"`
	Standerd [StanderdBS]byte `json:"stn"`
	Section  [SectionBS]byte  `json:"sec"`
	Group    [GroupBS]byte    `json:"grp"`
}

// Bytes return combined bytes of the ID
func (id *ClassID) Bytes() (b [ClassIDBS]byte) {
	ys := StanderdBS + YearBS
	yss := SectionBS + StanderdBS + YearBS
	for i := 0; i < ClassIDBS; i++ {
		if i < YearBS {
			b[i] = (*id).Year[i]
		} else if i >= YearBS && i < ys {
			j := i - YearBS
			b[i] = (*id).Standerd[j]
		} else if i >= StanderdBS && i < yss {
			j := i - ys
			b[i] = (*id).Section[j]
		} else {
			j := i - yss
			b[i] = (*id).Group[j]
		}
	}
	return
}

// Create return a new ClassID
func (id *ClassID) Create(yr [YearBS]byte, st [StanderdBS]byte, sec [SectionBS]byte, grp [GroupBS]byte) {
	(*id).Year = yr
	(*id).Standerd = st
	(*id).Section = sec
	(*id).Group = grp
}

// Init return adds values to  ClassID
func (id *ClassID) Init(cID [ClassIDBS]byte) {
	ys := StanderdBS + YearBS
	yss := SectionBS + StanderdBS + YearBS
	for i := 0; i < ClassIDBS; i++ {
		if i < YearBS {
			(*id).Year[i] = cID[i]
		} else if i >= YearBS && i < ys {
			j := i - YearBS
			(*id).Standerd[j] = cID[i]
		} else if i >= StanderdBS && i < yss {
			j := i - ys
			(*id).Section[j] = cID[i]
		} else {
			j := i - yss
			(*id).Group[j] = cID[i]
		}
	}
}

// Class in an struct to hold class data
type Class struct {
	ID          ClassID     `json:"id"`          // Unique identifier for a class
	StudentID   []StudentID `json:"student"`     // List of students in this class
	Subjects    []Subject   `json:"subjects"`    // List of subject to be studied in this class
	NFreePeriod int         `json:"nFreePeriod"` // Number of free period this class has
}

// Create assigns given classID and defults
// TODO Write Test
func (c *Class) Create(id ClassID) {
	(*c).ID = id              // assign ID to class
	(*c).NFreePeriod = MaxCap // assign the default max cap to class
}

// AddSubject adds the subejct to the class
// TODO Write Test
func (c *Class) AddSubject(sub Subject) {
	(*c).Subjects = append((*c).Subjects, sub) // add the new subject to the array
}

// AssignTeacher adds the subejct to the class
// TODO Write Test
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
// TODO Write Test
func (c *Class) CalRemCap() {
	(*c).NFreePeriod = MaxCap // assign the default max cap to class
	for _, s := range (*c).Subjects {
		// check if the teacherID is not !empty
		if s.IsAssigned() {
			(*c).NFreePeriod -= s.ReqClasses // reduce the no. of remaining capacity by the no. of period required by the subject
		}
	}
}

// CalCap calculate and update the remaining capacity of the class
// TODO Write Test
func (c *Class) CalCap() {
	(*c).NFreePeriod = MaxCap // assign the default max cap to class
	for _, s := range (*c).Subjects {
		// check if the teacherID is not !empty
		(*c).NFreePeriod -= s.ReqClasses // reduce the no. of remaining capacity by the no. of period required by the subject

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
						//fmt.Println(tMatchLength)
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

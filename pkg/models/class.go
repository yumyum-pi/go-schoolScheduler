package models

import (
	"fmt"
)

// ClassID is struct to store class id
type ClassID struct {
	Year     [YearBS]byte `json:"yr"`
	Standard byte         `json:"stn"`
	Section  byte         `json:"sec"`
	Group    byte         `json:"grp"`
}

// Bytes return combined bytes of the ID
func (id *ClassID) Bytes() (b ClassIDB) {
	b[0] = (*id).Year[0]
	b[1] = (*id).Year[1]
	b[2] = (*id).Standard
	b[3] = (*id).Section
	b[4] = (*id).Group
	return
}

// Init return adds values to  ClassID
func (id *ClassID) Init(b ClassIDB) {
	(*id).Year[0] = b[0]
	(*id).Year[1] = b[1]
	(*id).Standard = b[2]
	(*id).Section = b[3]
	(*id).Group = b[4]
}

// Class in an struct to hold class data
type Class struct {
	ID       ClassID  `json:"id"`          // Unique identifier for a class
	Subjects SubjectL `json:"subjects"`    // List of subject to be studied in this class
	Capacity int      `json:"nFreePeriod"` // No of free periods
}

// Print writes class values to the console
func (c *Class) Print() {
	fmt.Printf("> cID=%v\tcapacity=%v\n", (*c).ID.Bytes(), (*c).Capacity)
	PrintSubjectL((*c).Subjects)
}

// Init assigns given classID and default
func (c *Class) Init(id ClassIDB) {
	(*c).ID.Init(id)       // create class id from bytes
	(*c).Capacity = MaxCap // assign the default max cap to class
}

// AddSubject adds the subject to the class, reduce the class capacity and
// return the difference between class capacity and subject required periods
func (c *Class) AddSubject(sub Subject) (e error) {
	if sub == (Subject{}) {
		return fmt.Errorf(`> Error: given subject is nil`)
	}
	// calculate difference between class capacity and subject required
	d := (*c).Capacity - sub.Req
	// don't assign and new subject is difference is < 0
	if d < 0 {
		return fmt.Errorf(`> Error: Not enough capacity. requested=%v have=%v`, sub.Req, (*c).Capacity)

	}

	// check if already exist
	if i := (*c).Subjects.FindByID(sub.ID.Bytes()); i != -1 {
		return fmt.Errorf(`> Error: subject already exist`)
	}
	// add the subject to the class's subject slice
	(*c).Subjects = append((*c).Subjects, sub) // add the new subject to the array
	(*c).Capacity -= sub.Req                   // decrease the capacity
	return
}

// AssignTeacher adds the subject to the class
func (c *Class) AssignTeacher(sID SubjectID, tID TeacherID) (e error) {
	assigned := false

	// find the subject
	for i := range (*c).Subjects {
		// match the ID
		if (*c).Subjects[i].ID == sID {
			// check if assigned
			if (*c).Subjects[i].IsAssigned() {
				return fmt.Errorf("> Error: subject has already been assigned. sID=%v", sID)
			}
			// assign the teacher
			assigned = true
			(*c).Subjects[i].TeacherID = tID // add the teacher id to the subject
		}
	}
	// no subject found with the given subjectID
	if !assigned {
		return fmt.Errorf("> Error: subject id not found. sID=%v", sID)
	}
	return
}

// CalRemCap calculate and return the no of periods not assigned
func (c *Class) CalRemCap() (notAssigned int) {
	notAssigned = MaxCap // assign the default max
	for _, s := range (*c).Subjects {
		// check if the teacherID is not !empty
		if s.IsAssigned() {
			notAssigned -= s.Req // reduce the no. of remaining capacity by the no. of period required by the subject
		}
	}
	return
}

// CalCap calculate and update the remaining capacity of the class
func (c *Class) CalCap() {
	(*c).Capacity = MaxCap // assign the default max cap to class
	for _, s := range (*c).Subjects {
		// check if subjectID is not empty
		if s.ID != (SubjectID{}) {
			(*c).Capacity -= s.Req // reduce the no. of remaining capacity by the no. of period required by the subject
		}
	}
}

// PrintClassL writes the class data one by one on the console
func PrintClassL(cs []Class) {
	for _, c := range cs {
		c.Print()
	}
}

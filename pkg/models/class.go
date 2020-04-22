package models

import (
	"fmt"
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
	// calculate the offest
	secOffset := StanderdBS + YearBS   // calculate offect for standerd
	grpOffset := SectionBS + secOffset // calculate offect for group

	// loop through byte length
	for i := 0; i < ClassIDBS; i++ {
		if i < YearBS {
			// add year info to the struct
			b[i] = (*id).Year[i] // assign the bytes
		} else if i >= YearBS && i < secOffset {
			// add standerd info to the struct
			j := i - YearBS          // calculate the offset index of standerd
			b[i] = (*id).Standerd[j] // assign byte
		} else if i >= StanderdBS && i < grpOffset {
			// add section info to the struct
			j := i - secOffset      // calculate the offest index of section
			b[i] = (*id).Section[j] // assign byte
		} else {
			// add group info to the struct
			j := i - grpOffset    // calculate the offest index of type
			b[i] = (*id).Group[j] // assign byte
		}
	}
	return
}

// Init return adds values to  ClassID
func (id *ClassID) Init(cID [ClassIDBS]byte) {
	// calculate the offest
	secOffset := StanderdBS + YearBS   // calculate offect for standerd
	grpOffset := SectionBS + secOffset // calculate offect for group

	// loop through byte length
	for i := 0; i < ClassIDBS; i++ {
		if i < YearBS {
			// add year info to the byte
			(*id).Year[i] = cID[i] // assign the byte
		} else if i >= YearBS && i < secOffset {
			// add standerd info to the byte
			j := i - YearBS            //calculate the offect index of standerf
			(*id).Standerd[j] = cID[i] // assign the byte
		} else if i >= StanderdBS && i < grpOffset {
			// add section info to the byte
			j := i - secOffset        // calculate the offect index of section
			(*id).Section[j] = cID[i] // assign the byte
		} else {
			// add group info th=o the byte
			j := i - grpOffset      // calculate the offect indec of group
			(*id).Group[j] = cID[i] // assgin the byte
		}
	}
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
	for _, s := range (*c).Subjects {
		fmt.Printf("sID=%v\tReq=%v\ttID=%v\n", s.ID.Bytes(), s.Req, s.TeacherID.Bytes())
	}
}

// Init assigns given classID and defults
func (c *Class) Init(id [ClassIDBS]byte) {
	(*c).ID.Init(id)       // create class id from bytes
	(*c).Capacity = MaxCap // assign the default max cap to class
}

// AddSubject adds the subejct to the class, reduce the class capacity and
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

// AssignTeacher adds the subejct to the class
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

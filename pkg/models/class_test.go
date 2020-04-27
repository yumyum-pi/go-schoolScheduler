package models

import (
	"fmt"
	"testing"

	"github.com/yumyum-pi/go-schoolScheduler/internal/utils"
)

// wrong index create wrong index from the current index
// l = length of the array
// i = current index
func wrongIndex(l, i int) (j int) {
	j = i + 1 // make wrong index

	// check of out of range index
	if j >= l {
		// loop the value to be begining
		j = j - l // subtract the length to loop over
	}

	return
}

func tClassIDBytes(i int) error {
	j := wrongIndex(len(TClassIDBL), i) // create an incorrect index

	cID := TClassIDL[i] // get classID from the list
	b := cID.Bytes()    // create byte

	cBytes := TClassIDBL[i] // get correct bytes
	iBytes := TClassIDBL[j] // get incorrect class bytes

	// bytes match check
	p := (b == cBytes) // check with correct index
	f := (b == iBytes) // check with incorrect index

	// correct index should not be false
	if !p {
		return fmt.Errorf(
			"> Error: cID.Bytes=%v cBytes=%v, should match at=%v",
			b,
			cBytes,
			i,
		)
	}

	// incorrect index should not be true
	if f {
		return fmt.Errorf(
			"> Error: cID.Bytes=%v iBytes=%v, should not match at=%v",
			b,
			iBytes,
			i,
		)
	}
	return nil
}

func TestClassID_Bytes(t *testing.T) {
	l := len(TClassIDL) // get classID list length

	for i := 0; i < l; i++ {
		e := tClassIDBytes(i) // run the test function
		if e != nil {
			t.Error(e)
		}
	}
}

func tClassIDInit(i int) error {
	j := wrongIndex(len(TClassIDBL), i) // create an incorrect index

	var cID ClassID         // create a new classID
	cID.Init(TClassIDBL[i]) // assign value to the id

	cCID := TClassIDL[i] // correct classID
	iCID := TClassIDL[j] // incorrect classID

	// check if the classIDs match
	p := (cID == cCID) // this is the correct index
	f := (cID == iCID) // this is the incorrect index

	// check correct index
	// p should be true
	if !p {
		return fmt.Errorf(
			"> Error: cID=%v cCID=%v, should match at i=%v",
			cID,
			cCID,
			i,
		)
	}

	// check incorrect
	// f should not match
	if f {
		return fmt.Errorf(
			"> Error: cID=%v iCID=%v, should not match at i=%v",
			cID,
			iCID,
			i,
		)
	}
	return nil
}

func TestClassID_Init(t *testing.T) {
	l := len(TClassIDL) // get classID list length

	for i := 0; i < l; i++ {
		e := tClassIDInit(i) // run the test function
		if e != nil {
			t.Error(e)
		}
	}
}

func tClassInit(i int) error {
	j := wrongIndex(len(TClassIDL), i)

	// get ids
	cCID := TClassIDL[i] // correct index
	iCID := TClassIDL[j] // incorrect index

	var c Class          // create a new class
	c.Init(cCID.Bytes()) // initiate class with correct id

	// check correct index
	// class id should match
	if c.ID != cCID {
		return fmt.Errorf(
			"> Error: cID=%v cCID=%v, should match at i=%v",
			c.ID,
			cCID,
			i,
		)
	}

	// class subjects list should be 0
	if len(c.Subjects) != 0 {
		return fmt.Errorf(
			"> Error: cID=%v c.Subjects=%v, subject list should be 0 at i=%v",
			c.ID,
			c.Subjects,
			i,
		)
	}

	// class capacity should be full
	if c.Capacity != MaxCap {
		return fmt.Errorf(
			"> Error: cID=%v c.Capacity=%v, capacity should be %v at i=%v",
			c.ID,
			c.Capacity,
			MaxCap,
			i,
		)
	}

	// check incorrect index
	// class id should not match
	if c.ID == iCID {
		return fmt.Errorf(
			"> Error: cID=%v iCID=%v, should not match j=%v at i=%v",
			c.ID,
			iCID,
			j,
			i,
		)
	}
	return nil
}

func TestClass_Init(t *testing.T) {
	l := len(TClassIDL) // get length of the classID list

	for i := 0; i < l; i++ {
		e := tClassInit(i) // run the test function
		if e != nil {
			t.Error(e)
		}
	}
}

func TestClass_AddSubject(t *testing.T) {
	cID := TClassIDBL[0]
	// create a new class
	var c Class
	c.Init(cID)

	// loop through all the subjects
	for _, s := range TSubjectL {
		err := c.AddSubject(s) // add subject to the class
		if err != nil {
			t.Error(err)
		}
	}

	// length should be equal
	if len(c.Subjects) != len(TSubjectL) {
		t.Errorf(
			"> Error: c.Subject length is %v, length should be %v",
			len(c.Subjects),
			len(TSubjectL),
		)
	}

	// capacity should be 0
	if c.Capacity != 0 {
		t.Errorf("> Error: c.Capacity=%v, should be 0", c.Capacity)
	}

	// loop through all the subjects
	for i, s := range TSubjectL {
		// try to add another subject
		err := c.AddSubject(s)
		// check if successful
		if err == nil {
			t.Errorf(`> Error: added to class with no capacity`)
			return
		}

		// added subject with existing id
		j := wrongIndex(len(TSubjectL), i) // create wrong index of subject
		iS := c.Subjects[j]                // get the incorrect subject
		c.Subjects[i] = (Subject{})        // un assign a subject
		c.Capacity = s.Req                 // add capacity
		err = c.AddSubject(iS)             // assign subject that already exist
		if err == nil {
			t.Errorf(`> Error: added to class with the same id`)
			return
		}
		//check for nil error
		err = c.AddSubject((Subject{})) // assign a nil subject
		if err == nil {
			t.Errorf(`> Error: added a nil subject`)
			return
		}
		// assign the correct subject
		err = c.AddSubject(s) // assign subject that already exist
		if err != nil {
			t.Error(err)
			t.Errorf(`> Error: unable to add subject to the class with adequate capacity`)
			return
		}

	}

	// capacity should be 0
	if c.Capacity != 0 {
		t.Errorf("> Error: c.Capacity=%v, should be 0", c.Capacity)
	}

}

func TestClass_AssignTeacher(t *testing.T) {
	// generate a random numbers
	is := utils.GenerateRandomInt(len(TSubjectL), 10)   // for subject index
	it := utils.GenerateRandomInt(len(TTeacherIDL), 10) // for teacher index

	var c Class // create a teacherID
	c.Init(TClassIDL[0].Bytes())

	// assign all subjects
	for _, s := range TSubjectL {
		c.AddSubject(s)
	}

	// store id
	sID := c.Subjects[is].ID
	tID := TTeacherIDL[it]

	// assign a new teacher to a class with no capacity
	// expect to get error
	if e := c.AssignTeacher(sID, tID); e == nil {
		// not got on error
		t.Errorf("> Error: assigned a new teacher to a class with no capacity")
	}
	// un assign the teacher at the subject index
	c.Subjects[is].TeacherID = TeacherID{}

	// assign a new teacher to a class with capacity
	// expects no error
	if e := c.AssignTeacher(sID, tID); e != nil {
		t.Error(e)
	}

	// check is the teacher is assigned
	if c.Subjects[is].TeacherID != tID {
		t.Errorf(
			"> Error: TeacherID=%v,SubjectsID=%v TTeacherIDL[i]=%v",
			c.Subjects[is].TeacherID,
			c.Subjects[is].ID,
			TTeacherIDL[it],
		)
	}

	// check for unknown subject id
	// Remove the element at index "is" from class.Subjects.
	c.Subjects[is] = c.Subjects[len(c.Subjects)-1] // Copy last element to index "is".
	c.Subjects = c.Subjects[:len(c.Subjects)-1]    // Truncate slice.

	// assign a new teacher to a subject that does not exist
	// expect to get error
	if e := c.AssignTeacher(sID, tID); e == nil {
		// not got on error
		t.Errorf("> Error: assigned a new teacher to a subject that does not exist")
	}
}

func TestClass_CalRemCap(t *testing.T) {
	// generate two random no
	i := utils.GenerateRandomInt(len(TSubjectL), 10)
	j := wrongIndex(len(TSubjectL), i) // get different index from i

	// create a new class with subjects
	var c Class
	c.Init(TClassIDL[i].Bytes())                 // assign id and capacity
	c.Subjects = make([]Subject, len(TSubjectL)) // make the subject slice
	copy(c.Subjects, TSubjectL)                  // copy the TSubjectL slice

	up := c.CalRemCap() // calculate the no. of unassigned periods

	// check no of unassigned periods
	if up != 0 {
		t.Errorf("> Error: unassigned periods=%v", up)
	}

	// un assign a subjects
	c.Subjects[i].TeacherID = TeacherID{} // empty teacherID
	c.Subjects[j].ID = SubjectID{}        // empty subjectID

	// store the unassigned periods
	rPeriods := (c.Subjects[i].Req + c.Subjects[j].Req + up)

	up = c.CalRemCap() // calculate the no. of unassigned periods

	// check no of unassigned periods
	if up != rPeriods {
		t.Errorf("> Error: unassigned=%v", up)
	}
}

func TestClass_CalCap(t *testing.T) {
	// generate two random no
	i := utils.GenerateRandomInt(len(TSubjectL), 10)
	j := wrongIndex(len(TSubjectL), i) // get different index from i

	// create a new class
	var c Class
	c.Init(TClassIDL[i].Bytes())                 // assign id and capacity
	c.Subjects = make([]Subject, len(TSubjectL)) // make the subject slice
	copy(c.Subjects, TSubjectL)                  // copy the TSubjectL slice

	c.CalCap() // calculating the capacity

	// check capacity
	if c.Capacity != 0 {
		t.Errorf("> Error: capacity=%v, expecting 0", c.Capacity)
	}

	// make new no of change the subject requirement
	changeReq := c.Subjects[j].Req / 2
	uiS := c.Subjects[i]
	ujSReq := c.Subjects[j].Req - changeReq
	// make changes
	c.Subjects[i].ID = (SubjectID{}) // un assign a subject
	c.Subjects[j].Req = changeReq    // change the class requirement

	// store
	rPeriods := (uiS.Req + c.Capacity + ujSReq)

	c.CalCap()

	// check no of unassigned periods
	if c.Capacity != rPeriods {
		t.Errorf("> Error: capacity=%v rPeriods=%v\n", c.Capacity, rPeriods)
	}
}

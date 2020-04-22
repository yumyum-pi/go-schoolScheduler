package models

import (
	"fmt"
	"testing"

	"github.com/yumyum-pi/go-schoolScheduler/internal/utils"
)

// ClassIDs is a slice of ClassID for test
var tClassIDL = []ClassID{
	{Year: [4]byte{2, 0, 2, 0}, Standerd: [2]byte{0, 1}, Section: [2]byte{0, 1}, Group: [2]byte{0, 1}}, // 0
	{Year: [4]byte{2, 0, 2, 0}, Standerd: [2]byte{0, 2}, Section: [2]byte{0, 1}, Group: [2]byte{0, 1}}, // 1
	{Year: [4]byte{2, 0, 2, 0}, Standerd: [2]byte{0, 3}, Section: [2]byte{0, 1}, Group: [2]byte{0, 1}}, // 2
	{Year: [4]byte{2, 0, 2, 0}, Standerd: [2]byte{0, 4}, Section: [2]byte{0, 1}, Group: [2]byte{0, 1}}, // 3
	{Year: [4]byte{2, 0, 2, 0}, Standerd: [2]byte{0, 5}, Section: [2]byte{0, 1}, Group: [2]byte{0, 1}}, // 4
	{Year: [4]byte{2, 0, 2, 0}, Standerd: [2]byte{0, 6}, Section: [2]byte{0, 1}, Group: [2]byte{0, 1}}, // 5
	{Year: [4]byte{2, 0, 2, 0}, Standerd: [2]byte{0, 7}, Section: [2]byte{0, 1}, Group: [2]byte{0, 1}}, // 6
	{Year: [4]byte{2, 0, 2, 0}, Standerd: [2]byte{0, 8}, Section: [2]byte{0, 1}, Group: [2]byte{0, 1}}, // 7
	{Year: [4]byte{2, 0, 2, 0}, Standerd: [2]byte{0, 9}, Section: [2]byte{0, 1}, Group: [2]byte{0, 1}}, // 8
	{Year: [4]byte{2, 0, 2, 0}, Standerd: [2]byte{1, 0}, Section: [2]byte{0, 1}, Group: [2]byte{0, 1}}, // 9
}

// tClassIDBytes is a slice of test bytes of classIDs
var tClassIDBL [][10]byte = [][10]byte{
	[ClassIDBS]byte{2, 0, 2, 0, 0, 1, 0, 1, 0, 1},
	[ClassIDBS]byte{2, 0, 2, 0, 0, 2, 0, 1, 0, 1},
	[ClassIDBS]byte{2, 0, 2, 0, 0, 3, 0, 1, 0, 1},
	[ClassIDBS]byte{2, 0, 2, 0, 0, 4, 0, 1, 0, 1},
	[ClassIDBS]byte{2, 0, 2, 0, 0, 5, 0, 1, 0, 1},
	[ClassIDBS]byte{2, 0, 2, 0, 0, 6, 0, 1, 0, 1},
	[ClassIDBS]byte{2, 0, 2, 0, 0, 7, 0, 1, 0, 1},
	[ClassIDBS]byte{2, 0, 2, 0, 0, 8, 0, 1, 0, 1},
	[ClassIDBS]byte{2, 0, 2, 0, 0, 9, 0, 1, 0, 1},
	[ClassIDBS]byte{2, 0, 2, 0, 1, 0, 0, 1, 0, 1},
}

// Classes is a slice of class for test
var tClasses = []Class{
	{ID: tClassIDL[0], Subjects: tSubjectL[:1], Capacity: 42}, // 0-9: 6 - 48
	{ID: tClassIDL[1], Subjects: tSubjectL[:2], Capacity: 36}, // 0-4: 12 - 48
	{ID: tClassIDL[2], Subjects: tSubjectL[:3], Capacity: 30}, // 0-5: 18 - 48
	{ID: tClassIDL[3], Subjects: tSubjectL[:4], Capacity: 24}, // 0-4: 24 - 48
	{ID: tClassIDL[4], Subjects: tSubjectL[:5], Capacity: 18}, // 0-3: 30 - 48
	{ID: tClassIDL[5], Subjects: tSubjectL[:6], Capacity: 12}, // 0-6: 36 - 48
	{ID: tClassIDL[6], Subjects: tSubjectL[:7], Capacity: 9},  // 0-5: 39 - 48
	{ID: tClassIDL[7], Subjects: tSubjectL[:8], Capacity: 6},  // 0-2: 42 - 48
	{ID: tClassIDL[8], Subjects: tSubjectL[:9], Capacity: 3},  // 0-3: 45 - 48
	{ID: tClassIDL[9], Subjects: tSubjectL[:], Capacity: 0},   // 0-9: 48 - 48
}

// wrong index create wrong index from the current index
// l = length of the array
// i = current index
func wrongIndex(l, i int) (j int) {
	j = i + 1 // make wrong index

	// check of out of range index
	if j >= l {
		// loop the value to be begining
		j = j - l // subract the lenght to loop over
	}

	return
}

func tClassIDBytes(i int) (e error) {
	j := wrongIndex(len(tClassIDBL), i) // create an incorrect index

	cID := tClassIDL[i] // get classID from the list
	b := cID.Bytes()    // create byte

	cBytes := tClassIDBL[i] // get correct bytes
	iBytes := tClassIDBL[j] // get incorrect class bytes

	// bytes match check
	p := (b == cBytes) // check with correct index
	f := (b == iBytes) // check with incorrect index

	// correct index should not be false
	if !p {
		e = fmt.Errorf(`> Error: cID.Bytes=%v cBytes=%v, should match at=%v`, b, cBytes, i)
	}

	// incorrect index should not be true
	if f {
		e = fmt.Errorf(`> Error: cID.Bytes=%v iBytes=%v, should not match at=%v`, b, iBytes, i)
	}
	return
}

func TestClassID_Bytes(t *testing.T) {
	l := len(tClassIDL) // get classID list length

	for i := 0; i < l; i++ {
		e := tClassIDBytes(i) // run the test function
		if e != nil {
			t.Error(e)
		}
	}
}

func tClassIDInit(i int) (e error) {
	j := wrongIndex(len(tClassIDBL), i) // create an incorrect index

	var cID ClassID         // create a new classID
	cID.Init(tClassIDBL[i]) // assign value to the id

	cCID := tClassIDL[i] // correct classID
	iCID := tClassIDL[j] // incorrect classID

	// check if the classIDs match
	p := (cID == cCID) // this is the correct index
	f := (cID == iCID) // this is the incorrect index

	// check correct index
	// p should be true
	if !p {
		e = fmt.Errorf(`> Error: cID=%v cCID=%v, should match at i=%v`, cID, cCID, i)
	}

	// check incorrect
	// f should not match
	if f {
		e = fmt.Errorf(`> Error: cID=%v iCID=%v, should not match at i=%v`, cID, iCID, i)
	}
	return
}

func TestClassID_Init(t *testing.T) {
	l := len(tClassIDL) // get classID list length

	for i := 0; i < l; i++ {
		e := tClassIDInit(i) // run the test function
		if e != nil {
			t.Error(e)
		}
	}
}

func tClassInit(i int) (e error) {
	j := wrongIndex(len(tClassIDL), i)

	// get ids
	cCID := tClassIDL[i] // correct index
	iCID := tClassIDL[j] // incorrect index

	var c Class          // create a new class
	c.Init(cCID.Bytes()) // initiate class with correct id

	// check correct index
	// class id should match
	if c.ID != cCID {
		return fmt.Errorf("> Error: cID=%v cCID=%v, should match at i=%v", c.ID, cCID, i)
	}

	// class subjects list should be 0
	if len(c.Subjects) != 0 {
		e = fmt.Errorf("> Error: cID=%v c.Subjects=%v, subject list should be 0 at i=%v ", c.ID, c.Subjects, i)
	}

	// class capacity should be full
	if c.Capacity != MaxCap {
		e = fmt.Errorf("> Error: cID=%v c.Capacity=%v, capacity should be %v at i=%v, ", c.ID, c.Capacity, MaxCap, i)
	}

	// check incorrect index
	// class id should not match
	if c.ID == iCID {
		return fmt.Errorf("> Error: cID=%v iCID=%v, should not match j=%v at i=%v", c.ID, iCID, j, i)
	}
	return
}

func TestClass_Init(t *testing.T) {
	l := len(tClassIDL) // get length of the classID list

	for i := 0; i < l; i++ {
		e := tClassInit(i) // run the test function
		if e != nil {
			t.Error(e)
		}
	}
}

func TestClass_AddSubejct(t *testing.T) {
	cID := tClassIDBL[0]
	// create a new class
	var c Class
	c.Init(cID)

	// loop through all the subjects
	for _, s := range tSubjectL {
		err := c.AddSubject(s) // add subject to the class
		if err != nil {
			t.Error(err)
		}
	}

	// length should be equal
	if len(c.Subjects) != len(tSubjectL) {
		t.Errorf("> Error: c.Subject length is %v, length should be %v", len(c.Subjects), len(tSubjectL))
	}

	// capacity should be 0
	if c.Capacity != 0 {
		t.Errorf("> Error: c.Capacity=%v, should be 0", c.Capacity)
	}

	// loop through all the subjects
	for i, s := range tSubjectL {
		// try to add another subject
		err := c.AddSubject(s)
		// check if successful
		if err == nil {
			t.Errorf(`> Error: added to class with no capacity`)
			return
		}

		// added subject with existing id
		j := wrongIndex(len(tSubjectL), i) // create wrong index of subject
		iS := c.Subjects[j]                // get the incorrect subject
		c.Subjects[i] = (Subject{})        // unassign a subject
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
	is := utils.GenerateRandomInt(len(tSubjectL), 10)   // for subject index
	it := utils.GenerateRandomInt(len(tTeacherIDL), 10) // for teacher index

	var c Class // create a teacherID
	c.Init(tClassIDL[0].Bytes())

	// assign all subjects
	for _, s := range tSubjectL {
		c.AddSubject(s)
	}

	// store id
	sID := c.Subjects[is].ID
	tID := tTeacherIDL[it]

	// assign a new teacher to a class with no capacity
	// expect to get error
	if e := c.AssignTeacher(sID, tID); e == nil {
		// not got on error
		t.Errorf("> Error: assigend a new teacher to a class with no capacity")
	}
	// unassign the teacher at the subject index
	c.Subjects[is].TeacherID = TeacherID{}

	// assign a new teacher to a class with capacity
	// expects no error
	if e := c.AssignTeacher(sID, tID); e != nil {
		t.Error(e)
	}

	// check is the teacher is assigned
	if c.Subjects[is].TeacherID != tID {
		t.Errorf("> Error: TeacherID=%v,SubjectsID=%v tTeacherIDL[i]=%v", c.Subjects[is].TeacherID, c.Subjects[is].ID, tTeacherIDL[it])
	}

	// check for unknow subject id
	// Remove the element at index "is" from class.Subjects.
	c.Subjects[is] = c.Subjects[len(c.Subjects)-1] // Copy last element to index "is".
	c.Subjects = c.Subjects[:len(c.Subjects)-1]    // Truncate slice.

	// assign a new teacher to a subject that does not exist
	// expect to get error
	if e := c.AssignTeacher(sID, tID); e == nil {
		// not got on error
		t.Errorf("> Error: assigend a new teacher to a subject that does not exist")
	}
}

func TestClass_CalRemCap(t *testing.T) {
	// generate two random no
	i := utils.GenerateRandomInt(len(tSubjectL), 10)
	j := wrongIndex(len(tSubjectL), i) // get different index from i

	// create a new class with subjects
	var c Class
	c.Init(tClassIDL[i].Bytes())                 // assign id and capacity
	c.Subjects = make([]Subject, len(tSubjectL)) // make the subject slice
	copy(c.Subjects, tSubjectL)                  // copy the tSubjectL slice

	up := c.CalRemCap() // calculate the no. of unassigned periods

	// check no of unassigned periods
	if up != 0 {
		t.Errorf("> Error: unassigned periods=%v", up)
	}

	// unassign a subjects
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
	i := utils.GenerateRandomInt(len(tSubjectL), 10)
	j := wrongIndex(len(tSubjectL), i) // get different index from i

	// create a new class
	var c Class
	c.Init(tClassIDL[i].Bytes())                 // assign id and capacity
	c.Subjects = make([]Subject, len(tSubjectL)) // make the subject slice
	copy(c.Subjects, tSubjectL)                  // copy the tSubjectL slice

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
	c.Subjects[i].ID = (SubjectID{}) // unassign a subject
	c.Subjects[j].Req = changeReq    // change the class requirement
	//fmt.Println(i, j)
	// store
	rPeriods := (uiS.Req + c.Capacity + ujSReq)
	//fmt.Println(uiS.Req, c.Capacity, ujSReq, changeReq)
	c.CalCap()

	// check no of unassigned periods
	if c.Capacity != rPeriods {
		t.Errorf("> Error: capacity=%v rPeriods=%v\n", c.Capacity, rPeriods)
	}
}

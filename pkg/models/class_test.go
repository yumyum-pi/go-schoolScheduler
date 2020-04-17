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

// Classes is a slice of class for test
var tClasses = Classes{
	Class{ID: tClassIDL[0], Subjects: tSubjectL[:1], Capacity: 42}, // 0-9: 6 - 48
	Class{ID: tClassIDL[1], Subjects: tSubjectL[:2], Capacity: 36}, // 0-4: 12 - 48
	Class{ID: tClassIDL[2], Subjects: tSubjectL[:3], Capacity: 30}, // 0-5: 18 - 48
	Class{ID: tClassIDL[3], Subjects: tSubjectL[:4], Capacity: 24}, // 0-4: 24 - 48
	Class{ID: tClassIDL[4], Subjects: tSubjectL[:5], Capacity: 18}, // 0-3: 30 - 48
	Class{ID: tClassIDL[5], Subjects: tSubjectL[:6], Capacity: 12}, // 0-6: 36 - 48
	Class{ID: tClassIDL[6], Subjects: tSubjectL[:7], Capacity: 9},  // 0-5: 39 - 48
	Class{ID: tClassIDL[7], Subjects: tSubjectL[:8], Capacity: 6},  // 0-2: 42 - 48
	Class{ID: tClassIDL[8], Subjects: tSubjectL[:9], Capacity: 3},  // 0-3: 45 - 48
	Class{ID: tClassIDL[9], Subjects: tSubjectL[:], Capacity: 0},   // 0-9: 48 - 48
}

// tClassIDBytes is a slice of test bytes of classIDs
var tClassIDBytes [][10]byte = [][10]byte{
	[ClassIDBS]byte{2, 0, 2, 0, 0, 1, 0, 1, 0, 1},
	[ClassIDBS]byte{2, 0, 2, 0, 0, 2, 0, 1, 0, 1},
	[ClassIDBS]byte{2, 0, 2, 0, 0, 3, 0, 1, 0, 1},
	[ClassIDBS]byte{2, 0, 2, 0, 0, 4, 0, 1, 0, 1},
	[ClassIDBS]byte{2, 0, 2, 0, 0, 5, 0, 1, 0, 1}, // false
	[ClassIDBS]byte{2, 0, 2, 0, 0, 6, 0, 1, 0, 1},
	[ClassIDBS]byte{2, 0, 2, 0, 0, 7, 0, 1, 0, 1},
	[ClassIDBS]byte{2, 0, 2, 0, 0, 8, 0, 1, 0, 1},
	[ClassIDBS]byte{2, 0, 2, 0, 0, 9, 0, 1, 0, 1},
	[ClassIDBS]byte{2, 0, 2, 0, 1, 0, 0, 1, 0, 1}, // false
}

// wrong index create wrong index from the current index
// l = length of the array
// i = current index
func wrongIndex(l, i int) (j int) {
	j = i + 1 // make wrong index

	// check of out of range index
	if j >= l {
		j = j - l
	}

	return
}

// classIDTestBytes tests the byte method of the classID struct
func classIDTestBytes(i int) (e error) {
	j := wrongIndex(len(tClassIDBytes), i) // create wrong index

	cID := tClassIDL[i] // the classID from the list
	b := cID.Bytes()    // run the method to be tested

	// bytes match check
	p := (b == tClassIDBytes[i]) // check with correct index
	f := (b == tClassIDBytes[j]) // check with wrong index

	// correct index should be true
	if !p {
		e = fmt.Errorf(`> Error: cID=%v, bytes=%v, tClassIDBytes=%v. where n=%v and tClassIDBytesResult=%v`, cID, b, tClassIDBytes[i], i, true)
	}

	// wrong index should be false
	if f {
		e = fmt.Errorf(`> Error: cID=%v, bytes=%v, tClassIDBytes=%v. where n=%v and tClassIDBytesResult=%v`, cID, b, tClassIDBytes[i], i, true)
	}
	return
}

// TestClassID_Bytes_One test's the byte method of the classID struct and use all
func TestClassID_Bytes_All(t *testing.T) {
	l := len(tClassIDL)

	for i := 0; i < l; i++ {
		e := classIDTestBytes(i) // run the test function
		if e != nil {
			t.Error(e)
		}
	}
}

// classIDTestInit test the Init method of ClassID struct
func classIDTestInit(i int) (e error) {
	j := wrongIndex(len(tClassIDBytes), i) // create wrong index

	var cID ClassID            // create a new classID
	cID.Init(tClassIDBytes[i]) // assign value to the id

	// check if the classIDs match
	p := (cID == tClassIDL[i]) // this is the correct index
	f := (cID == tClassIDL[j]) // this is the incorrect index

	// p used the correct index, it should be true
	if !p {
		return fmt.Errorf("> Error: cID=%v bytes=%v tCID=%v where i=%v", cID, tClassIDBytes[i], tClassIDL[i], i)
	}

	// f used the incorrect index, it should not be true
	if f {
		return fmt.Errorf("> Error: cID=%v bytes=%v tCID=%v where i=%v", cID, tClassIDBytes[i], tClassIDL[i], i)
	}

	return
}

// TestClassID_Init_One test the Init method of ClassID struct and uses one element
func TestClassID_Init_All(t *testing.T) {
	// Get length of the tClassIDTUlist
	l := len(tClassIDL)

	for i := 0; i < l; i++ {
		e := classIDTestInit(i)
		if e != nil {
			t.Error(e)
		}
	}
}

func tClassInit(i int) (e error) {
	var nc Class         // create a new class
	tCID := tClassIDL[i] // create a classID

	// run the function to be tested
	nc.Init(tCID.Bytes()) // initiate class  witht the cID bytes

	// new class id and test class id should match
	if nc.ID != tCID {
		e = fmt.Errorf("> Error: IDs don't match newClassID=%v tClassId=%v", nc.ID, tCID)
	}

	// new class should not have subjects assigned
	if len(nc.Subjects) != 0 {
		e = fmt.Errorf("> Error: Subject is not 0, newClassID=%v tClassId=%v", nc.ID, nc.ID)
	}

	if nc.Capacity != MaxCap {
		e = fmt.Errorf("> Error: Subject is not 0, capacity=%v", nc.Capacity)
	}
	return
}

func TestClass_Init(t *testing.T) {
	l := len(tClassIDL)

	for i := 0; i < l; i++ {
		e := classIDTestInit(i)
		if e != nil {
			t.Error(e)
		}
	}
}

func TestClass_AddSubejct(t *testing.T) {
	cID := tClassIDBytes[0]
	// create a new class
	var c Class
	c.Init(cID)

	// loop through all the subjects
	for _, s := range tSubjectL {
		c.AddSubject(s) // add subject to the class
	}

	// check subject length
	if len(c.Subjects) != len(tSubjectL) {
		t.Errorf("> Error: c.Subject length=%v values.Subject length=%v", len(c.Subjects), len(tSubjectL))
	}

	// check capacity
	if c.Capacity != 0 {
		t.Errorf("> Error: c.Capacity is not 0. c.Capacity=%v", c.Capacity)
	}

	// try to add another subject
	diff := c.AddSubject(tSubjectL[0])
	// check if successful
	if diff >= 0 {
		t.Errorf("> Error: Added another subject when no capacity. c.Capacity=%v", c.Capacity)
	}
}

func TestClass_AssignTeacher(t *testing.T) {
	// generate a random numbers
	is := utils.GenerateRandomInt(len(tSubjectL), 10)   // for subject index
	it := utils.GenerateRandomInt(len(tTeacherIDL), 10) // for teacher index
	// create a teacherID
	// assign the teacher to the class
	var c Class
	c.Init(tClassIDL[0].Bytes())
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
	j := utils.GenerateRandomInt(len(tSubjectL), 10)

	// create a new class
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
	j := utils.GenerateRandomInt(len(tSubjectL), 10)

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

	// make changes
	c.Subjects[i].ID = (SubjectID{}) // unassign a subject
	c.Subjects[j].Req = changeReq    // change the class requirement

	// store
	rPeriods := (c.Subjects[i].Req + changeReq + c.Capacity)

	c.CalCap()
	// check no of unassigned periods
	if c.Capacity != rPeriods {
		t.Errorf("> Error: capacity=%v rPeriods=%v\n", c.Capacity, rPeriods)
	}
}

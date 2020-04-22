package models

import (
	"fmt"
	"testing"

	"github.com/yumyum-pi/go-schoolScheduler/internal/utils"
)

// SubjectIDs is the slice of SubjectID for test
var tSubjectIDL = []SubjectID{
	{Standerd: [2]byte{1, 0}, Type: [4]byte{1, 1, 3, 0}}, // 0
	{Standerd: [2]byte{2, 9}, Type: [4]byte{3, 2, 2, 4}}, // 1
	{Standerd: [2]byte{3, 8}, Type: [4]byte{5, 4, 3, 3}}, // 2
	{Standerd: [2]byte{4, 7}, Type: [4]byte{4, 6, 5, 4}}, // 3
	{Standerd: [2]byte{5, 6}, Type: [4]byte{5, 5, 6, 0}}, // 4
	{Standerd: [2]byte{6, 5}, Type: [4]byte{0, 6, 6, 7}}, // 5
	{Standerd: [2]byte{7, 4}, Type: [4]byte{8, 0, 7, 7}}, // 6
	{Standerd: [2]byte{8, 3}, Type: [4]byte{8, 9, 3, 8}}, // 7
	{Standerd: [2]byte{9, 2}, Type: [4]byte{8, 9, 2, 8}}, // 8
	{Standerd: [2]byte{0, 1}, Type: [4]byte{8, 9, 1, 8}}, // 9
}

// list of subjectID bytes slice
var tSubjectIDBL = [][SubjectIDBS]byte{
	{1, 0, 1, 1, 3, 0}, // 0
	{2, 9, 3, 2, 2, 4}, // 1
	{3, 8, 5, 4, 3, 3}, // 2
	{4, 7, 4, 6, 5, 4}, // 3
	{5, 6, 5, 5, 6, 0}, // 4
	{6, 5, 0, 6, 6, 7}, // 5
	{7, 4, 8, 0, 7, 7}, // 6
	{8, 3, 8, 9, 3, 8}, // 7
	{9, 2, 8, 9, 2, 8}, // 8
	{0, 1, 8, 9, 1, 8}, // 9
}

// Subjects is a slice of Subject for test
var tSubjectL = []Subject{
	{ID: tSubjectIDL[0], TeacherID: tTeacherIDL[0], Req: 6}, // 6
	{ID: tSubjectIDL[1], TeacherID: tTeacherIDL[1], Req: 6}, // 12
	{ID: tSubjectIDL[2], TeacherID: tTeacherIDL[2], Req: 6}, // 18
	{ID: tSubjectIDL[3], TeacherID: tTeacherIDL[3], Req: 6}, // 24
	{ID: tSubjectIDL[4], TeacherID: tTeacherIDL[4], Req: 6}, // 30
	{ID: tSubjectIDL[5], TeacherID: tTeacherIDL[5], Req: 6}, // 36
	{ID: tSubjectIDL[6], TeacherID: tTeacherIDL[6], Req: 3}, // 39
	{ID: tSubjectIDL[7], TeacherID: tTeacherIDL[7], Req: 3}, // 42
	{ID: tSubjectIDL[8], TeacherID: tTeacherIDL[8], Req: 3}, // 45
	{ID: tSubjectIDL[9], TeacherID: tTeacherIDL[9], Req: 3}, // 48
}

func tSubjectIDBytes(i int) (e error) {
	j := wrongIndex(len(tSubjectIDBL), i) // create an incorrect index

	sID := tSubjectIDL[i] // get subjectID from the list
	b := sID.Bytes()      // create byte

	cByte := tSubjectIDBL[i] // get the correct byte
	iByte := tSubjectIDBL[j] // get the incorret byte

	// bytes match check
	p := (b == cByte) // check with correct index
	f := (b == iByte) // check with incorrect index

	// correct index should not be false
	if !p {
		e = fmt.Errorf(`> Error: sID.Bytes=%v cByte=%v, should match at i=%v`, b, cByte, i)
	}

	// incorrect index should not be true
	if f {
		e = fmt.Errorf(`> Error: sID.Bytes=%v, iByte=%v, should not match at n=%v`, b, iByte, i)
	}
	return
}

// check byte function
func TestSubjectID_Bytes(t *testing.T) {
	l := len(tSubjectIDBL) // get subjectID list length

	for i := 0; i < l; i++ {
		e := tSubjectIDBytes(i) // run the test function
		if e != nil {
			t.Error(e)
		}
	}
}

func tSubjectIDInit(i int) (e error) {
	j := wrongIndex(len(tSubjectIDBL), i) // create an incorrect index

	var sID SubjectID         // create new subjectID
	sID.Init(tSubjectIDBL[i]) // initiate the value

	cSID := tSubjectIDL[i] // correct subjectID
	iSID := tSubjectIDL[j] // incorrect subjectID

	// check if the teacherIDs match
	p := (sID == cSID) // check with correct index
	f := (sID == iSID) // check with correct index

	// check corrent index
	// p should be true
	if !p {
		e = fmt.Errorf(`> Error: sID=%v cSID=%v, should match at i=%v`, sID, cSID, i)
	}

	// check incorrect
	// f should not match
	if f {
		e = fmt.Errorf(`> Error: sID=%v iSID=%v, should not match at i=%v`, sID, iSID, i)
	}

	return
}

// check init function
func TestSubjectID_Init(t *testing.T) {
	l := len(tSubjectIDBL) // get subjectID list length

	for i := 0; i < l; i++ {
		e := tSubjectIDInit(i) // run the test function
		if e != nil {
			t.Error(e)
		}
	}
}

func tSubjectIsAssigned(i int) (e error) {
	s := tSubjectL[i] // get subjectID from the list

	// #1 test
	// check if assigned
	if !s.IsAssigned() {
		e = fmt.Errorf(`> Error: sID=%v should be true at i=%v`, s.ID, i)
	}

	// #2 test -check false true - unassigned teacher
	s.TeacherID = (TeacherID{}) // unassign a teacherID

	// check false true
	if s.IsAssigned() {
		e = fmt.Errorf(`> Error: sID=%v teacher removed, should be false at i=%v`, s.ID, i)
	}

	// #2 test -check false true - unassigned teacher
	s.TeacherID = tTeacherIDL[i] // reassign the teacher ID
	s.ID = (SubjectID{})         // unassign a subjectID

	// check false true
	if s.IsAssigned() {
		e = fmt.Errorf(`> Error: sID=%v ID removed, should be false at i=%v`, s.ID, i)
	}
	return
}

func TestSubject_IsAssigned(t *testing.T) {
	// Get length of the names
	l := len(tSubjectL)

	// check for true
	for i := 0; i < l; i++ {
		e := tSubjectIsAssigned(i)
		if e != nil {
			t.Error(e)
		}
	}
}

func TestSubjectL_FindByID(t *testing.T) {
	// copy the subject slice
	var tSL SubjectL
	tSL = append(tSubjectL[:0:0], tSubjectL...)

	l := len(tSL) // get the length of the subject list

	// loop though the length of the subject list
	for i := 0; i < l; i++ {
		j := utils.GenerateRandomInt(len(tSubjectIDBL), 10) // generate a random index
		cSID := tSubjectIDBL[j]                             // get the correct subjectID
		iSID := [SubjectIDBS]byte{6, 6, 6, 6, 6, 6}         // generate uncorrect subjectID

		// find the index with ids
		p := tSL.FindByID(cSID) // correct subjectID
		f := tSL.FindByID(iSID) // incorrect subjectID

		// check the correct subjectID
		// p must be 1
		if p == -1 {
			t.Errorf(`> Error: cSID=%v was not found in the list`, cSID)
		}

		// check the incorrect subjectID
		if f != -1 {
			t.Errorf(`> Error: iSID=%v was found in the list`, iSID)
		}
	}

}

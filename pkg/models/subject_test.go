package models

import (
	"fmt"
	"testing"

	"github.com/yumyum-pi/go-schoolScheduler/internal/utils"
)

func tSubjectIDBytes(i int) (e error) {
	j := wrongIndex(len(TSubjectIDBL), i) // create an incorrect index

	sID := TSubjectIDL[i] // get subjectID from the list
	b := sID.Bytes()      // create byte

	cByte := TSubjectIDBL[i] // get the correct byte
	iByte := TSubjectIDBL[j] // get the incorret byte

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
	l := len(TSubjectIDBL) // get subjectID list length

	for i := 0; i < l; i++ {
		e := tSubjectIDBytes(i) // run the test function
		if e != nil {
			t.Error(e)
		}
	}
}

func tSubjectIDInit(i int) (e error) {
	j := wrongIndex(len(TSubjectIDBL), i) // create an incorrect index

	var sID SubjectID         // create new subjectID
	sID.Init(TSubjectIDBL[i]) // initiate the value

	cSID := TSubjectIDL[i] // correct subjectID
	iSID := TSubjectIDL[j] // incorrect subjectID

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
	l := len(TSubjectIDBL) // get subjectID list length

	for i := 0; i < l; i++ {
		e := tSubjectIDInit(i) // run the test function
		if e != nil {
			t.Error(e)
		}
	}
}

func tSubjectIsAssigned(i int) (e error) {
	s := TSubjectL[i] // get subjectID from the list

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
	s.TeacherID = TTeacherIDL[i] // reassign the teacher ID
	s.ID = (SubjectID{})         // unassign a subjectID

	// check false true
	if s.IsAssigned() {
		e = fmt.Errorf(`> Error: sID=%v ID removed, should be false at i=%v`, s.ID, i)
	}
	return
}

func TestSubject_IsAssigned(t *testing.T) {
	// Get length of the names
	l := len(TSubjectL)

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
	tSL = append(TSubjectL[:0:0], TSubjectL...)

	l := len(tSL) // get the length of the subject list

	// loop though the length of the subject list
	for i := 0; i < l; i++ {
		j := utils.GenerateRandomInt(len(TSubjectIDBL), 10) // generate a random index
		cSID := TSubjectIDBL[j]                             // get the correct subjectID
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

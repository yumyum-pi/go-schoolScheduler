package models

import (
	"fmt"
	"testing"

	"github.com/yumyum-pi/go-schoolScheduler/internal/utils"
)

// TeacherID

func tTeacherIDBytes(i int) (e error) {
	// create an incorrect id index
	j := wrongIndex(len(TTeacherIDL), i)

	tID := TTeacherIDL[i] // get teacherID from the list
	b := tID.Bytes()      // create bytes

	cByte := TTeacherIDBL[i] // get the correct byte
	iByte := TTeacherIDBL[j] // get the incorrect byte

	// bytes match check
	p := (b == cByte) // check with correct index
	f := (b == iByte) // check with incorrect index

	// correct index should not be false
	if !p {
		e = fmt.Errorf(`> Error: tID=%v, byte=%v cByte=%v match should be true`, tID, b, cByte)
	}

	// incorrect index should not true
	if f {
		e = fmt.Errorf(`> Error: tID=%v, byte=%v iByte=%v match should not be true`, tID, b, iByte)
	}
	return
}

func TestTeacherID_Bytes(t *testing.T) {
	l := len(TTeacherIDL) // get teacherID list length

	for i := 0; i < l; i++ {
		e := tTeacherIDBytes(i) // run the test function
		if e != nil {
			t.Error(e)
		}
	}
}

func tTeacherIDInit(i int) (e error) {
	j := wrongIndex(len(TTeacherIDL), i) // create an incorrect index

	var tID TeacherID         // create new teacherID
	tID.Init(TTeacherIDBL[i]) // initiate the value

	cTID := TTeacherIDL[i] // correct teacherID
	iTID := TTeacherIDL[j] // incorrect teacherID

	// check if the teacherIDs match
	p := (tID == cTID) // check with correct index
	f := (tID == iTID) // check with correct index

	// check correct index
	// p should be true
	if !p {
		e = fmt.Errorf(`> Error: tID=%v cTID=%v, should match at i=%v`, tID, cTID, i)
	}

	// check incorrect
	// f should not match
	if f {
		e = fmt.Errorf(`> Error: tID=%v iTID=%v, should not match at i=%v`, tID, iTID, i)
	}
	return
}

func TestTeacherID_Init(t *testing.T) {
	l := len(TTeacherIDL) // get teacherID list length

	for i := 0; i < l; i++ {
		e := tTeacherIDInit(i) // run the test function
		if e != nil {
			t.Error(e)
		}
	}
}

// ClassAssigned

func TClassAssigned(i int) (e error) {
	j := wrongIndex(len(TTeacherIDL), i)    // create an incorrect index
	req := utils.GenerateRandomInt(100, 10) // create a random number

	// get correct info
	cSID := TSubjectIDL[i] // get correct subjectID
	cCID := TClassIDL[i]   // get correct classID

	// get incorrect info
	iSID := TSubjectIDL[j] // get incorrect subjectID
	iCID := TClassIDL[j]   // get incorrect classID
	// create new classAssigned struct
	var ca ClassAssigned
	ca.Init(TSubjectIDL[i].Bytes(), TClassIDL[i].Bytes(), req)

	// correct index
	// classID should be equal
	if ca.ClassID != cCID {
		return fmt.Errorf(`> Error: ca.ClassID=%v cCID=%v should be equal at i=%v`, ca.ClassID, cCID, i)
	}
	// subjectID should be equal
	if ca.SubjectID != cSID {
		return fmt.Errorf(`> Error: ca.SubjectID=%v cSID=%v should be equal at i=%v`, ca.SubjectID, cSID, i)
	}
	// Requirement should be equal
	if ca.Assigned != req {
		return fmt.Errorf(`> Error: ca.Req=%v req=%v should be equal at i=%v`, ca.Assigned, req, i)
	}

	// incorrect index
	// classID should not be equal
	if ca.ClassID == iCID {
		return fmt.Errorf(`> Error: ca.ClassID=%v iCID=%v should not be equal at i=%v`, ca.ClassID, iCID, i)
	}
	// subjectID should not be equal
	if ca.SubjectID == iSID {
		return fmt.Errorf(`> Error: ca.SubjectID=%v iSID=%v should be equal at i=%v`, ca.SubjectID, iSID, i)
	}
	return
}

func TestClassAssigned_Init(t *testing.T) {
	l := len(TTeacherIDL) // get teacherID list length

	for i := 0; i < l; i++ {
		e := TClassAssigned(i) // run the test function
		if e != nil {
			t.Error(e)
		}
	}
}

// Teacher

func tTeacherInit(i int) (e error) {
	j := wrongIndex(len(TTeacherIDL), i) // create incorrect index

	// get ids
	cTID := TTeacherIDL[i] // correct index
	iTID := TTeacherIDL[j] // incorrect index

	var t Teacher        //create a new teacher
	t.Init(cTID.Bytes()) // initiate teacher with correct id

	// check correct index
	// teacher id should match
	if t.ID != cTID {
		return fmt.Errorf(`> Error: tID=%v  cTID=%v, should match at i=%v`, t.ID, cTID, i)
	}

	// subjectCT should be 0
	if len(t.SubjectCT) != 0 {
		return fmt.Errorf(`> Error: tID=%v  t.SubjectCT=%v, should be 0 at i=%v`, t.ID, t.SubjectCT, i)
	}

	// class assigned should be 0
	if len(t.AClassL) != 0 {
		return fmt.Errorf(`> Error: tID=%v  t.ClassesAssigned=%v, should be 0 at i=%v`, t.ID, t.AClassL, i)
	}

	// capacity should be teacherCap
	if t.Capacity != TeacherCap {
		return fmt.Errorf(`> Error: tID=%v  t.Capacity=%v, should be %v at i=%v`, t.ID, t.Capacity, TeacherCap, i)
	}

	// check incorrect index
	// teacher should not match
	if t.ID == iTID {
		return fmt.Errorf(`> Error: tID=%v  iTID=%v should match at i=%v`, t.ID, iTID, i)
	}

	// capacity should not be 0
	if t.Capacity == 0 {
		return fmt.Errorf(`> Error: tID=%v  t.Capacity=%v, should be 0 at i=%v`, t.ID, t.Capacity, i)
	}
	return
}

func TestTeacher_Init(t *testing.T) {
	l := len(TTeacherIDL) // get list's length

	for i := 0; i < l; i++ {
		// run the test function
		if e := tTeacherInit(i); e != nil {
			t.Error(e)
		}
	}
}

func tTeacherCanTeach(i int) error {
	// get ids
	cSID := TSubjectIDL[i] // correct
	iSID := cSID           // incorrect
	iSID.Standard = 99     // modify to be incorrect

	// test
	p := TTeacherL[i].CanTeach(cSID) // correct id
	f := TTeacherL[i].CanTeach(iSID) // incorrect id

	// check correct id
	// should return true
	if !p {
		return fmt.Errorf("> Error: p should return true at i=%v", i)
	}

	// check incorrect id
	// should return false
	if f {
		return fmt.Errorf("> Error: f should return false i=%v", i)
	}

	return nil
}

func TestTeacher_CanTeach(t *testing.T) {
	l := len(TTeacherL) // get list length

	for i := 0; i < l; i++ {
		// run the test function
		if e := tTeacherCanTeach(i); e != nil {
			t.Error(e)
		}
	}
}

func tTeacherAssignClass(i int) error {
	t := TTeacherL[i]                    // get teacher
	cID := TClassIDL[i]                  // get classID
	sID := TSubjectIDL[i]                // get subjectID
	tcc := t.Capacity                    // teacher current capacity
	r := utils.GenerateRandomInt(10, 10) // periods required by the new subject

	// assign correct class
	diff := t.AssignClass(cID, sID, r)

	// if the subject and class can not be assigned to the teacher but are assigned
	// through error
	if r > tcc && diff >= 0 {
		return fmt.Errorf("> Error: class and subject should not be assigned at i:=%v", i)
	}

	// if the subject and class can be assigned to the teacher but are not assigned
	// through error
	if r < tcc && diff < 0 {
		return fmt.Errorf("> Error: class and subject should be assigned at i:=%v", i)
	}

	return nil
}

func TestTeacher_AssignClass(t *testing.T) {
	l := len(TTeacherL) // get list length

	for i := 0; i < l; i++ {
		// run the test function
		if e := tTeacherAssignClass(i); e != nil {
			t.Error(e)
		}
	}
}

// Teachers

func tTeachersFindIndex(i int) (e error) {
	// get ids
	cTID := TTeacherIDL[i] // correct id
	iTID := TeacherID{}

	ts := Teachers{}
	// check of empty array
	em := ts.FindIndex(cTID)

	// should not return an index
	if em != -1 {
		return fmt.Errorf(`> Error: em=%v, empty list should return -1 at i=%v`, em, i)
	}

	// add teacher to the array
	ts = TTeacherL

	// find index
	p := ts.FindIndex(cTID) // use correct id
	f := ts.FindIndex(iTID) // use incorrect id

	// check correct index
	// should return an index
	if p == -1 {
		return fmt.Errorf(`> Error: p=%v, should return an index at i=%v`, p, i)
	}

	// should match the index
	if p != i {
		return fmt.Errorf(`> Error: p=%v, should match the index at i=%v`, p, i)
	}

	// check wrong index
	// should not return an index
	if f != -1 {
		return fmt.Errorf(`> Error: f=%v, should be -1 at i=%v`, f, i)
	}

	return
}

func TestTeachers_FindIndex(t *testing.T) {
	l := len(TTeacherL) // get list's length

	for i := 0; i < l; i++ {
		// run the test function
		if e := tTeachersFindIndex(i); e != nil {
			t.Error(e)
		}
	}
}

func tTeachersFindBySub(i int) error {
	l := len(TTeacherL) // get list's length
	// get subject id
	sID := TSubjectIDL[i]

	// get the teacher index
	var ts Teachers = TTeacherL
	tsR := ts.FindBySub(sID)

	// check if the no. of teacher returned is correct or not
	if len(tsR)+i != l {
		return fmt.Errorf("> Error: no. of teacher=%v, expecting=%v at i=%v", len(tsR), l-i, i)
	}

	//loop through each teacher and check if can teacher
	for _, index := range tsR {
		if !TTeacherL[index].CanTeach(sID) {
			return fmt.Errorf("> Error: teacher with given index does not teach the subject")
		}
	}

	return nil
}

func TestTeachers_FindBySub(t *testing.T) {
	l := len(TTeacherL) // get list's length

	for i := 0; i < l; i++ {
		// run the test function
		if e := tTeachersFindBySub(i); e != nil {
			t.Error(e)
		}
	}
}

func tTeacherFindBySubType(i int) error {
	//copy the teachers list
	var ts Teachers
	ts = append(ts[:0:0], TTeacherL...)

	// get the last subjectID in the from the teacher
	lSCT := len(ts[i].SubjectCT)
	sID := ts[i].SubjectCT[lSCT-1]

	// test the function
	indexes := ts.FindBySubType(sID.Type)

	// check
	//index length should be 3
	r := len(ts) - i
	if len(indexes) != r {
		return fmt.Errorf("> Error: sID.Type=%v no of indexes returned=%v, exprected=%v at i=%v", sID.Type, len(indexes), r, i)
	}

	// loop through the returned indexes
	// check if they can teach the subject
	for _, index := range indexes {
		if !ts[index].CanTeach(sID) {
			return fmt.Errorf("> Teacher: %v", ts[index])
		}
	}
	return nil
}

func TestTeachers_FindBySubType(t *testing.T) {
	l := len(TTeacherL) // get list's length

	for i := 0; i < l; i++ {
		// run the test function
		if e := tTeacherFindBySubType(i); e != nil {
			t.Error(e)
		}
	}

}

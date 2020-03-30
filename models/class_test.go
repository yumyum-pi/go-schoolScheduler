package models

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/yumyum-pi/go-schoolScheduler/utils"
)

type classIDTestUnit struct {
	Stn [2]byte
	Sec [2]byte
	Grp [2]byte
	Yr  [4]byte
}

var classIDTestGrp []classIDTestUnit = []classIDTestUnit{
	classIDTestUnit{[2]byte{'0', '8'}, [2]byte{'0', '1'}, [2]byte{'0', '1'}, [4]byte{2, 0, 2, 0}},
	classIDTestUnit{[2]byte{'0', '8'}, [2]byte{'0', '1'}, [2]byte{'0', '2'}, [4]byte{2, 0, 2, 0}},
	classIDTestUnit{[2]byte{'0', '8'}, [2]byte{'0', '2'}, [2]byte{'0', '1'}, [4]byte{2, 0, 2, 0}},
	classIDTestUnit{[2]byte{'0', '8'}, [2]byte{'0', '2'}, [2]byte{'0', '2'}, [4]byte{2, 0, 2, 0}},
	classIDTestUnit{[2]byte{'0', '8'}, [2]byte{'0', '2'}, [2]byte{'0', '1'}, [4]byte{2, 0, 2, 0}},
	classIDTestUnit{[2]byte{'0', '7'}, [2]byte{'0', '1'}, [2]byte{'0', '1'}, [4]byte{2, 0, 2, 0}},
	classIDTestUnit{[2]byte{'0', '7'}, [2]byte{'0', '1'}, [2]byte{'0', '2'}, [4]byte{2, 0, 2, 0}},
	classIDTestUnit{[2]byte{'0', '7'}, [2]byte{'0', '2'}, [2]byte{'0', '1'}, [4]byte{2, 0, 2, 0}},
	classIDTestUnit{[2]byte{'0', '7'}, [2]byte{'0', '2'}, [2]byte{'0', '2'}, [4]byte{2, 0, 2, 0}},
	classIDTestUnit{[2]byte{'0', '7'}, [2]byte{'0', '2'}, [2]byte{'0', '2'}, [4]byte{2, 0, 2, 0}},
}

var classIDs []ClassID = []ClassID{
	ClassID{[2]byte{'0', '8'}, [2]byte{'0', '1'}, [2]byte{'0', '1'}, [4]byte{2, 0, 2, 0}},
	ClassID{[2]byte{'0', '8'}, [2]byte{'0', '1'}, [2]byte{'0', '2'}, [4]byte{2, 0, 2, 0}},
	ClassID{[2]byte{'0', '8'}, [2]byte{'0', '2'}, [2]byte{'0', '1'}, [4]byte{2, 0, 2, 0}},
	ClassID{[2]byte{'0', '8'}, [2]byte{'0', '2'}, [2]byte{'0', '2'}, [4]byte{2, 0, 2, 0}},
	ClassID{[2]byte{'0', '8'}, [2]byte{'0', '2'}, [2]byte{'0', '2'}, [4]byte{2, 0, 2, 0}}, //
	ClassID{[2]byte{'0', '7'}, [2]byte{'0', '1'}, [2]byte{'0', '1'}, [4]byte{2, 0, 2, 0}},
	ClassID{[2]byte{'0', '7'}, [2]byte{'0', '1'}, [2]byte{'0', '2'}, [4]byte{2, 0, 2, 0}},
	ClassID{[2]byte{'0', '7'}, [2]byte{'0', '2'}, [2]byte{'0', '1'}, [4]byte{2, 0, 2, 0}},
	ClassID{[2]byte{'0', '7'}, [2]byte{'0', '2'}, [2]byte{'0', '2'}, [4]byte{2, 0, 2, 0}},
	ClassID{[2]byte{'0', '7'}, [2]byte{'0', '2'}, [2]byte{'0', '1'}, [4]byte{2, 0, 2, 0}}, //
}

var createTestResultBool []bool = []bool{
	true,
	true,
	true,
	true,
	false,
	true,
	true,
	true,
	true,
	false,
}

var byteTestResult [][10]byte = [][10]byte{
	[10]byte{2, 0, 2, 0, '0', '8', '0', '1', '0', '1'},
	[10]byte{2, 0, 2, 0, '0', '8', '0', '1', '0', '2'},
	[10]byte{2, 0, 2, 0, '0', '8', '0', '2', '0', '1'},
	[10]byte{2, 0, 2, 0, '0', '8', '0', '2', '0', '2'},
	[10]byte{'0', '8', '0', '2', '0', '5', '0', '1', '0', '1'}, //
	[10]byte{2, 0, 2, 0, '0', '7', '0', '1', '0', '1'},
	[10]byte{2, 0, 2, 0, '0', '7', '0', '1', '0', '2'},
	[10]byte{2, 0, 2, 0, '0', '7', '0', '2', '0', '1'},
	[10]byte{2, 0, 2, 0, '0', '7', '0', '2', '0', '2'},
	[10]byte{'0', '7', '0', '2', '0', '0', '2', '0', '2', '0'}, //
}

func TestEqualSign(t *testing.T) {
	// Get length of the classIDTestGrp & subreact 1 from l to avoid
	// generating random number which is out of bound of the array.
	l := len(classIDTestGrp) - 1

	n := rand.Intn(l) // generate a random no. between 0 and l
	m := rand.Intn(l) // generate a random no. between 0 and l

	// loop if n == m
	for m == n {
		m = rand.Intn(l)
	}

	// check if the ids are same or not
	if classIDs[n] == classIDs[m] {
		t.Errorf(`> Error: n="%v", m="%v". The classes should not be equal`, n, m)
	}
}

func classIDTestUnitF(n int) (e error) {
	var t bool
	var testID ClassID
	// Get the classID Unit for testing
	var unit classIDTestUnit = classIDTestGrp[n]
	// Create a new test classID
	testID = testID.Create(unit.Stn, unit.Sec, unit.Grp, unit.Yr)

	// Compare the testID with the classIDUnit classID
	t = testID == classIDs[n]

	// check the result to the list of results
	if t != createTestResultBool[n] {
		e = fmt.Errorf(`> Error: testID="%v", classID="%v". where n="%v" and createTestResultBool="%v"`, testID, classIDs[n], n, createTestResultBool[n])
	}
	return e
}

func TestClassID_Create_One(t *testing.T) {
	// Get length of the classIDTestGrp & subreact 1 from l to avoid
	// generating random number which is out of bound of the array.
	l := len(classIDTestGrp) - 1
	n := rand.Intn(l) // generate a random no. between 0 and l

	e := classIDTestUnitF(n)
	if e != nil {
		t.Error(e)
	}
}

func TestClassID_Create_All(t *testing.T) {
	// Get length of the classIDTestGrp
	l := len(classIDTestGrp)

	for i := 0; i < l; i++ {
		e := classIDTestUnitF(i)
		if e != nil {
			t.Error(e)
		}
	}
}

func classIDByteF(n int) (e error) {
	var t bool
	// Get the classID Unit for testing
	testID := classIDs[n]
	b := testID.Bytes()
	t = (b == byteTestResult[n])
	// check the result to the list of results
	if t != createTestResultBool[n] {
		e = fmt.Errorf(`> Error: testID="%v", bytes="%v", byteTestResult="%v". where n="%v" and createTestResultBool="%v"`, testID, b, byteTestResult[n], n, createTestResultBool[n])
	}
	return e
}

func TestClassID_BytestOne(t *testing.T) {
	// Get length of the classIDTestGrp & subreact 1 from l to avoid
	// generating random number which is out of bound of the array.
	l := len(classIDTestGrp) - 1
	n := utils.GenerateRandomInt(l, 10) // generate a random no. between 0 and l

	e := classIDByteF(n)
	if e != nil {
		t.Error(e)
	}
}

var subjects = []Subject{
	Subject{SubjectID{[2]byte{0, 5}, [4]byte{3, 2, 4, 0}}, 5, TeacherID{}},
	Subject{SubjectID{[2]byte{0, 5}, [4]byte{0, 3, 2, 4}}, 6, TeacherID{}},
	Subject{SubjectID{[2]byte{0, 5}, [4]byte{4, 0, 3, 2}}, 7, TeacherID{}},
	Subject{SubjectID{[2]byte{0, 5}, [4]byte{2, 4, 0, 3}}, 8, TeacherID{}},
}

var teacherIDs = []TeacherID{
	TeacherID{[4]byte{2, 0, 2, 2}, [4]byte{1, 0, 2, 2}},
	TeacherID{[4]byte{2, 0, 2, 2}, [4]byte{2, 0, 2, 2}},
	TeacherID{[4]byte{2, 0, 2, 2}, [4]byte{3, 0, 2, 2}},
	TeacherID{[4]byte{2, 0, 2, 2}, [4]byte{4, 0, 2, 2}},
}

func TestClass_AddSubject(t *testing.T) {
	// create a class
	var cls Class
	cls.Init(classIDs[0])
	for _, s := range subjects {
		cls.AddSubject(s)
	}

	// loop through the class subject list
	for i, s := range cls.Subjects {
		// check subjects match
		if s != subjects[i] {
			t.Errorf("> Error: s=\"%v\", subjects=\"%v\" where i=\"%v\"\n", s, subjects[i], i)
		}
	}
}

func TestClass_AssignTeacher(t *testing.T) {
	// generate a random no
	i := utils.GenerateRandomInt(len(teacherIDs), 10)
	// create a teacherID
	// assign the teacher to the class
	var cls Class
	cls.Init(classIDs[i])
	for _, s := range subjects {
		cls.AddSubject(s)
	}

	cls.AssignTeacher(cls.Subjects[i].ID, teacherIDs[i])

	// check is the teacher is assigned
	if cls.Subjects[i].TeacherID != teacherIDs[i] {
		t.Errorf("> Error: cls.Subjects[i].TeacherID=\"%v\",cls.Subjects[i].ID=\"%v\" teacherIDs[i]=\"%v\" where i=\"%v\"\n", cls.Subjects[i].TeacherID, cls.Subjects[i].ID, teacherIDs[i], i)
	}

	// check if the assigned function works
	if cls.Subjects[i].TeacherID == (TeacherID{}) {
		t.Errorf("> Error: cls.Subjects[i].TeacherID=\"%v\", TeacherID{}=\"%v\" where i=\"%v\"\n", cls.Subjects[i].TeacherID, TeacherID{}, i)
	}
	// check if the NfreePeriod reduces
	if cls.NFreePeriod != (MaxCap - cls.Subjects[i].ReqClasses) {
		t.Errorf("> Error: cls.NFreePeriod=\"%v\", (MaxCap - cls.Subjects[i].ReqClasses)=\"%v\" where i=\"%v\"\n", cls.NFreePeriod, (MaxCap - cls.Subjects[i].ReqClasses), i)
	}
}

func TestClass_CalRemCap(t *testing.T) {
	// generate a random no
	i := utils.GenerateRandomInt(len(teacherIDs), 10)
	// create a teacherID
	// assign the teacher to the class
	var periodAssigned int
	var cls Class
	cls.Init(classIDs[i])
	for _, s := range subjects {
		s.TeacherID = teacherIDs[i]
		periodAssigned += s.ReqClasses
		cls.AddSubject(s)
	}
	cls.CalRemCap()
	// check if the NfreePeriod reduces
	if cls.NFreePeriod != (MaxCap - periodAssigned) {
		t.Errorf("> Error: cls.NFreePeriod=\"%v\", periodAssigned=\"%v\" where i=\"%v\"\n", cls.NFreePeriod, periodAssigned, i)
	}
}

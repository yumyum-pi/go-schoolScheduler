package models

import (
	"fmt"
	"testing"
)

// classIDTU stores essingtail data for assigning a classID
type classIDTU struct {
	Yr  [YearBS]byte     // year
	Stn [StanderdBS]byte // standerd
	Sec [SectionBS]byte  // section
	Grp [GroupBS]byte    // group
}

// classIDTestGrp is a slice of classIDTU for multiple unit test
var tClassIDTUlist []classIDTU = []classIDTU{
	{[YearBS]byte{2, 0, 2, 0}, [StanderdBS]byte{0, 8}, [SectionBS]byte{0, 1}, [GroupBS]byte{0, 1}},
	{[YearBS]byte{2, 0, 2, 0}, [StanderdBS]byte{0, 8}, [SectionBS]byte{0, 1}, [GroupBS]byte{0, 2}},
	{[YearBS]byte{2, 0, 2, 0}, [StanderdBS]byte{0, 8}, [SectionBS]byte{0, 2}, [GroupBS]byte{0, 1}},
	{[YearBS]byte{2, 0, 2, 0}, [StanderdBS]byte{0, 8}, [SectionBS]byte{0, 2}, [GroupBS]byte{0, 2}},
	{[YearBS]byte{2, 0, 2, 0}, [StanderdBS]byte{0, 8}, [SectionBS]byte{0, 2}, [GroupBS]byte{0, 1}},
	{[YearBS]byte{2, 0, 2, 0}, [StanderdBS]byte{0, 7}, [SectionBS]byte{0, 1}, [GroupBS]byte{0, 1}},
	{[YearBS]byte{2, 0, 2, 0}, [StanderdBS]byte{0, 7}, [SectionBS]byte{0, 1}, [GroupBS]byte{0, 2}},
	{[YearBS]byte{2, 0, 2, 0}, [StanderdBS]byte{0, 7}, [SectionBS]byte{0, 2}, [GroupBS]byte{0, 1}},
	{[YearBS]byte{2, 0, 2, 0}, [StanderdBS]byte{0, 7}, [SectionBS]byte{0, 2}, [GroupBS]byte{0, 2}},
	{[YearBS]byte{2, 0, 2, 0}, [StanderdBS]byte{0, 7}, [SectionBS]byte{0, 2}, [GroupBS]byte{0, 2}},
}

// tClassIDList is a slice of test classIDs
var tClassIDList []ClassID = []ClassID{
	{[YearBS]byte{2, 0, 2, 0}, [StanderdBS]byte{0, 8}, [SectionBS]byte{0, 1}, [GroupBS]byte{0, 1}},
	{[YearBS]byte{2, 0, 2, 0}, [StanderdBS]byte{0, 8}, [SectionBS]byte{0, 1}, [GroupBS]byte{0, 2}},
	{[YearBS]byte{2, 0, 2, 0}, [StanderdBS]byte{0, 8}, [SectionBS]byte{0, 2}, [GroupBS]byte{0, 1}},
	{[YearBS]byte{2, 0, 2, 0}, [StanderdBS]byte{0, 8}, [SectionBS]byte{0, 2}, [GroupBS]byte{0, 2}},
	{[YearBS]byte{2, 0, 2, 0}, [StanderdBS]byte{0, 8}, [SectionBS]byte{0, 2}, [GroupBS]byte{0, 2}}, // false
	{[YearBS]byte{2, 0, 2, 0}, [StanderdBS]byte{0, 7}, [SectionBS]byte{0, 1}, [GroupBS]byte{0, 1}},
	{[YearBS]byte{2, 0, 2, 0}, [StanderdBS]byte{0, 7}, [SectionBS]byte{0, 1}, [GroupBS]byte{0, 2}},
	{[YearBS]byte{2, 0, 2, 0}, [StanderdBS]byte{0, 7}, [SectionBS]byte{0, 2}, [GroupBS]byte{0, 1}},
	{[YearBS]byte{2, 0, 2, 0}, [StanderdBS]byte{0, 7}, [SectionBS]byte{0, 2}, [GroupBS]byte{0, 2}},
	{[YearBS]byte{2, 0, 2, 0}, [StanderdBS]byte{0, 7}, [SectionBS]byte{0, 2}, [GroupBS]byte{0, 1}}, // false
}

// tClassIDBytes is a slice of test bytes of classIDs
var tClassIDBytes [][10]byte = [][10]byte{
	[ClassIDBS]byte{2, 0, 2, 0, 0, 8, 0, 1, 0, 1},
	[ClassIDBS]byte{2, 0, 2, 0, 0, 8, 0, 1, 0, 2},
	[ClassIDBS]byte{2, 0, 2, 0, 0, 8, 0, 2, 0, 1},
	[ClassIDBS]byte{2, 0, 2, 0, 0, 8, 0, 2, 0, 2},
	[ClassIDBS]byte{0, 8, 0, 2, 0, 5, 0, 1, 0, 1}, // false
	[ClassIDBS]byte{2, 0, 2, 0, 0, 7, 0, 1, 0, 1},
	[ClassIDBS]byte{2, 0, 2, 0, 0, 7, 0, 1, 0, 2},
	[ClassIDBS]byte{2, 0, 2, 0, 0, 7, 0, 2, 0, 1},
	[ClassIDBS]byte{2, 0, 2, 0, 0, 7, 0, 2, 0, 2},
	[ClassIDBS]byte{0, 7, 0, 2, 0, 0, 2, 0, 2, 0}, // false
}

// tClassIDBytesResult is a slice of bool for the test results
var tClassIDBytesResult []bool = []bool{
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

// classIDTestBytes tests the byte methord of the classID struct
func classIDTestBytes(n int) (e error) {
	cID := tClassIDList[n] // the classID from the list
	b := cID.Bytes()       // use the Byte methord

	// compare b and the tClassIDBytesResult
	t := (b == tClassIDBytes[n])

	// check the result matches the expectioned result
	if t != tClassIDBytesResult[n] {
		return fmt.Errorf(`> Error: cID=%v, bytes=%v, tClassIDBytes=%v. where n=%v and tClassIDBytesResult=%v`, cID, b, tClassIDBytes[n], n, tClassIDBytesResult[n])
	}
	return e
}

// TestClassID_Bytes_One test's the byte methord of the classID struct and use all
func TestClassID_Bytes_All(t *testing.T) {
	// Get length of the tClassIDTUlist
	l := len(tClassIDTUlist)

	for i := 0; i < l; i++ {
		e := classIDTestBytes(i)
		if e != nil {
			t.Error(e)
		}
	}
}

// classIDTestInit test the Init methord of ClassID struct
func classIDTestInit(i int) (e error) {
	var cID ClassID
	// uss the Init methord to assign value to the id
	cID.Init(tClassIDBytes[i])

	// check if the classIDs match
	t := (cID == tClassIDList[i])

	// check if the result matchs the exprected result
	if t != tClassIDBytesResult[i] {
		return fmt.Errorf("> Error: cID=%v bytes=%v tCID=%v where i=%v", cID, tClassIDBytes[i], tClassIDList[i], i)
	}
	return
}

// TestClassID_Init_One test the Init methord of ClassID struct and uses one element
func TestClassID_Init_All(t *testing.T) {
	// Get length of the tClassIDTUlist
	l := len(tClassIDTUlist)

	for i := 0; i < l; i++ {
		e := classIDTestInit(i)
		if e != nil {
			t.Error(e)
		}
	}
}

/*
func TestClass_Create(t *testing.T) {
	newClass := Class{}
	// Get length of the tClassIDTUlist
	l := len(tClassIDTUlist)
	cID := tClassIDList[l]

	tClass := Class{
		ID:          cID,
		StudentID:   []StudentID{},
		Subjects:    []Subject{},
		NFreePeriod: MaxCap,
	}

	newClass.Create(cID)
	// store error
	ss := ""
	// check if they match
	if newClass.ID != tClass.ID {
		ss += fmt.Sprintf(" newClassID=%v tClassId=%v", newClass.ID, tClass.ID)
	}
	if newClass.Subjects != ([]Subject{}) {
		ss += fmt.Sprintf(" newClassID=%v tClassId=%v", newClass.ID, tClass.ID)
	}
}
*/
/*
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
			t.Errorf("> Error: s=%v, subjects=%v where i=%v\n", s, subjects[i], i)
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
		t.Errorf("> Error: cls.Subjects[i].TeacherID=%v,cls.Subjects[i].ID=%v teacherIDs[i]=%v where i=%v\n", cls.Subjects[i].TeacherID, cls.Subjects[i].ID, teacherIDs[i], i)
	}

	// check if the assigned function works
	if cls.Subjects[i].TeacherID == (TeacherID{}) {
		t.Errorf("> Error: cls.Subjects[i].TeacherID=%v, TeacherID{}=%v where i=%v\n", cls.Subjects[i].TeacherID, TeacherID{}, i)
	}
	// check if the NfreePeriod reduces
	if cls.NFreePeriod != (MaxCap - cls.Subjects[i].ReqClasses) {
		t.Errorf("> Error: cls.NFreePeriod=%v, (MaxCap - cls.Subjects[i].ReqClasses)=%v where i=%v\n", cls.NFreePeriod, (MaxCap - cls.Subjects[i].ReqClasses), i)
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
		t.Errorf("> Error: cls.NFreePeriod=%v, periodAssigned=%v where i=%v\n", cls.NFreePeriod, periodAssigned, i)
	}
}
/*
*/

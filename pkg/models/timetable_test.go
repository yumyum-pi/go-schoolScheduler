package models

import (
	"fmt"
	"testing"
)

var TPeriodL = []Period{
	{TClassIDL[0], TSubjectIDL[0], TTeacherIDL[0]},
	{TClassIDL[1], TSubjectIDL[1], TTeacherIDL[1]},
	{TClassIDL[2], TSubjectIDL[2], TTeacherIDL[2]},
	{TClassIDL[3], TSubjectIDL[3], TTeacherIDL[3]},
	{TClassIDL[4], TSubjectIDL[4], TTeacherIDL[4]},
	{TClassIDL[5], TSubjectIDL[5], TTeacherIDL[5]},
	{TClassIDL[6], TSubjectIDL[6], TTeacherIDL[6]},
	{TClassIDL[7], TSubjectIDL[7], TTeacherIDL[7]},
	{TClassIDL[8], TSubjectIDL[8], TTeacherIDL[8]},
	{TClassIDL[9], TSubjectIDL[9], TTeacherIDL[9]},
}

var TPeriodBL = []PeriodB{
	{7, 228, 1, 0, 0, 7, 228, 0},
	{7, 228, 2, 1, 1, 7, 228, 1},
	{7, 228, 3, 2, 2, 7, 228, 2},
	{7, 228, 4, 3, 3, 7, 228, 3},
	{7, 228, 5, 4, 4, 7, 228, 4},
	{7, 228, 6, 5, 5, 7, 228, 5},
	{7, 228, 7, 6, 6, 7, 228, 6},
	{7, 228, 8, 7, 7, 7, 228, 7},
	{7, 228, 9, 8, 8, 7, 228, 8},
	{7, 228, 10, 9, 9, 7, 228, 9},
}

var TCPeriodL = []Period{
	{ClassID{}, SubjectID{}, TTeacherIDL[0]},
	{ClassID{}, SubjectID{}, TTeacherIDL[1]},
	{ClassID{}, SubjectID{}, TTeacherIDL[2]},
	{ClassID{}, SubjectID{}, TTeacherIDL[3]},
	{ClassID{}, SubjectID{}, TTeacherIDL[4]},
	{TClassIDL[5], SubjectID{}, TeacherID{}},
	{TClassIDL[6], SubjectID{}, TeacherID{}},
	{TClassIDL[7], SubjectID{}, TeacherID{}},
	{TClassIDL[8], SubjectID{}, TeacherID{}},
	{TClassIDL[9], SubjectID{}, TeacherID{}},
}

var TIPeriodL = []Period{
	{ClassID{}, SubjectID{}, TeacherID{}},
	{ClassID{}, SubjectID{}, TeacherID{}},
	{ClassID{}, SubjectID{}, TeacherID{}},
	{ClassID{}, SubjectID{}, TeacherID{}},
	{ClassID{}, SubjectID{}, TeacherID{}},
	{ClassID{}, SubjectID{}, TeacherID{}},
	{ClassID{}, SubjectID{}, TeacherID{}},
	{ClassID{}, SubjectID{}, TeacherID{}},
	{ClassID{}, SubjectID{}, TeacherID{}},
	{ClassID{}, SubjectID{}, TeacherID{}},
}

var TCPeriodBL = []PeriodB{
	{0, 0, 0, 0, 0, 7, 228, 0},  // 0
	{0, 0, 0, 0, 1, 7, 228, 1},  //
	{0, 0, 0, 0, 2, 7, 228, 2},  //
	{0, 0, 0, 0, 3, 7, 228, 3},  //
	{0, 0, 0, 0, 4, 7, 228, 4},  //
	{7, 228, 6, 5, 5, 0, 0, 0},  //
	{7, 228, 7, 6, 6, 0, 0, 0},  //
	{7, 228, 8, 7, 7, 0, 0, 0},  //
	{7, 228, 9, 8, 8, 0, 0, 0},  //
	{7, 228, 10, 9, 9, 0, 0, 0}, //
}

var TIPeriodBL = []PeriodB{
	{0, 0, 0, 0, 0, 0, 0, 0}, // 0
	{0, 0, 0, 0, 1, 0, 0, 0}, //
	{0, 0, 0, 0, 2, 0, 0, 0}, //
	{0, 0, 0, 0, 3, 0, 0, 0}, //
	{0, 0, 0, 0, 4, 0, 0, 0}, //
	{0, 0, 0, 0, 5, 0, 0, 0}, //
	{0, 0, 0, 0, 6, 0, 0, 0}, //
	{0, 0, 0, 0, 7, 0, 0, 0}, //
	{0, 0, 0, 0, 8, 0, 0, 0}, //
	{0, 0, 0, 0, 9, 0, 0, 0}, //
}

func tPeriodByte(i int) error {
	j := wrongIndex(len(TPeriodL), i) // create incorrect index

	cpb := TPeriodL[i].Bytes() // byte from correct index
	ipb := TPeriodL[j].Bytes() // byte from incorrect index

	// byte match check
	p := (cpb == TPeriodBL[i]) // check correct index
	f := (ipb == TPeriodBL[i]) // check incorrect index

	// correct index should not be false
	if !p {
		return fmt.Errorf(
			"> Error: cpd=%v pd=%v, should match at i=%v",
			cpb,
			TPeriodBL[i],
			i,
		)
	}

	// check incorrect
	// f should not match
	if f {
		return fmt.Errorf(
			"> Error: cpd=%v pd=%v, , should not match at i=%v",
			ipb,
			TPeriodBL[i],
			i,
		)
	}
	return nil
}

func TestPeriod_Bytes(t *testing.T) {
	l := len(TPeriodL) // get classID list length

	for i := 0; i < l; i++ {
		e := tPeriodByte(i) // run the test function
		if e != nil {
			t.Error(e)
		}
	}
}

func tPeriodInit(i int) error {
	j := wrongIndex(len(TPeriodL), i) // create incorrect index

	var cp, ip Period

	cp.Init(TPeriodBL[i]) // byte from correct index
	ip.Init(TPeriodBL[j]) // byte from incorrect index

	// byte match check
	p := (cp == TPeriodL[i]) // check correct index
	f := (ip == TPeriodL[i]) // check incorrect index

	// correct index should not be false
	if !p {
		return fmt.Errorf(
			"> Error: cp=%v pd=%v, should match at i=%v",
			cp,
			TPeriodL[i],
			i,
		)
	}

	// check incorrect
	// f should not match
	if f {
		return fmt.Errorf(
			"> Error: cpd=%v pd=%v, , should not match at i=%v",
			ip,
			TPeriodL[i],
			i,
		)
	}
	return nil
}

func TestPeriod_Init(t *testing.T) {
	l := len(TPeriodL) // get classID list length

	for i := 0; i < l; i++ {
		e := tPeriodInit(i) // run the test function
		if e != nil {
			t.Error(e)
		}
	}
}

func tPeriodIsAssigned(i int) error {
	cp := TCPeriodL[i] // assigned period
	ip := TIPeriodL[i] // unassigned periods

	// cp should be true
	if !cp.IsAssigned() {
		return fmt.Errorf(
			"> Error: cp=%v should be true",
			cp,
		)
	}

	// should be false
	if ip.IsAssigned() {
		return fmt.Errorf(
			"> Error: ip=%v should be false",
			ip,
		)
	}
	return nil
}

func TestPeriod_IsAssigned(t *testing.T) {
	l := len(TCPeriodL) // get classID list length

	for i := 0; i < l; i++ {
		e := tPeriodIsAssigned(i) // run the test function
		if e != nil {
			t.Error(e)
		}
	}
}

func tPeriodBIsAssigned(i int) error {
	cp := TCPeriodBL[i] // assigned period
	ip := TIPeriodBL[i] // unassigned periods

	// cp should be true
	if !cp.IsAssigned() {
		return fmt.Errorf(
			"> Error: cp=%v should be true",
			cp,
		)
	}

	// should be false
	if ip.IsAssigned() {
		return fmt.Errorf(
			"> Error: ip=%v should be false",
			ip,
		)
	}
	return nil
}

func TestPeriodB_IsAssigned(t *testing.T) {
	l := len(TCPeriodBL) // get classID list length

	for i := 0; i < l; i++ {
		e := tPeriodBIsAssigned(i) // run the test function
		if e != nil {
			t.Error(e)
		}
	}
}

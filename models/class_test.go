package models

import (
	"fmt"
	"math/rand"
	"testing"
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

var byteTestResult [][6]byte = [][6]byte{
	[6]byte{'0', '8', '0', '1', '0', '1'},
	[6]byte{'0', '8', '0', '1', '0', '2'},
	[6]byte{'0', '8', '0', '2', '0', '1'},
	[6]byte{'0', '8', '0', '2', '0', '2'},
	[6]byte{'0', '8', '0', '2', '0', '5'}, //
	[6]byte{'0', '7', '0', '1', '0', '1'},
	[6]byte{'0', '7', '0', '1', '0', '2'},
	[6]byte{'0', '7', '0', '2', '0', '1'},
	[6]byte{'0', '7', '0', '2', '0', '2'},
	[6]byte{'0', '7', '0', '2', '0', '0'}, //
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
	t = b == byteTestResult[n]
	// check the result to the list of results
	if t != createTestResultBool[n] {
		e = fmt.Errorf(`> Error: testID="%v", bytes="%v", byteTestResult="%v". where n="%v" and createTestResultBool="%v"`, testID, b, byteTestResult[n], n, createTestResultBool[n])
	}
	return e
}

func TestClassID_BytesOne(t *testing.T) {
	// Get length of the classIDTestGrp & subreact 1 from l to avoid
	// generating random number which is out of bound of the array.
	l := len(classIDTestGrp) - 1
	n := rand.Intn(l) // generate a random no. between 0 and l

	e := classIDByteF(n)
	if e != nil {
		t.Error(e)
	}
}

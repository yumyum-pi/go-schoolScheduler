package models

import (
	"fmt"
	"testing"
)

var byte6Unit [][6]byte = [][6]byte{
	[6]byte{5, 8, 3, 0, 7, 1},
	[6]byte{5, 8, 4, 'q', 9, 1},
	[6]byte{5, 3, 0, 7, 4, 1},
	[6]byte{4, 4, 6, 4, 9, 1},
	[6]byte{'2', 6, 's', 6, 6, 1},
	[6]byte{1, 8, 7, 5, 'a', 1},
	[6]byte{2, 3, 5, 'v', 3, 1},
}

func testOUSubjectIDInit(n int) (e error) {
	sID := new(SubjectID)
	sID.Init(byte6Unit[n])
	stn := [2]byte{
		byte6Unit[n][0],
		byte6Unit[n][1],
	}

	typ := [4]byte{
		byte6Unit[n][2],
		byte6Unit[n][3],
		byte6Unit[n][4],
		byte6Unit[n][5],
	}
	if sID.Standerd != stn {
		e = fmt.Errorf(`> Error: sID.Sranderd="%v",stn="%v" where n="%v"`, sID.Standerd, stn, n)
	} else if sID.Type != typ {
		e = fmt.Errorf(`> Error: sID.Type="%v",typ="%v" where n="%v"`, sID.Type, typ, n)
	} else if sID.Bytes() != byte6Unit[n] {
		e = fmt.Errorf(`> Error: sID.Bytes="%v",byte6Unit="%v" where n="%v"`, sID.Bytes(), byte6Unit[n], n)
	}

	return
}

// check init function
func TestSubjectID_Init(t *testing.T) {
	// Get length of the names
	l := len(byte6Unit)

	for i := 0; i < l; i++ {
		e := testOUSubjectIDInit(i)

		if e != nil {
			t.Error(e)
		}
	}
}

// check byte function

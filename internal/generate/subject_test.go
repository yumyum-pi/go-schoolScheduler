package generate

import (
	"fmt"
	"testing"

	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

type subDub map[models.SubjectIDB]int

func tGenerateSubject(ss []models.Subject) error {
	cap := 0 // calculate the capacity

	// no of subjects should be equal to lstcode
	// length of the given list
	if len(ss) != lSTCode {
		return fmt.Errorf("> len(s)=%v is not LSTCode=%v ", len(ss), lSTCode)
	}

	subj := make(subDub) // store the index of a subjectID
	// no subject should have 0 requirements
	for j, s := range ss {
		if s.Req == 0 {
			return fmt.Errorf("> Error: subject req=0 at index=%v. sID=%v", j, s.ID.Bytes())
		}
		// find the current subjectID
		k, ok := subj[s.ID.Bytes()]
		if ok {
			// subjectID found in the map
			// subjectID is reperating
			return fmt.Errorf("> Error: sID=%v is repearting at index=%v", s.ID.Bytes(), k)
		}
		subj[s.ID.Bytes()] = k
		cap += s.Req
	}

	// capacity should not be 0
	if cap != models.MaxCap {
		return fmt.Errorf("> c.Capacity=%v in not MaxCap=%v", cap, models.MaxCap)
	}
	return nil
}

func TestGenerateSubject(t *testing.T) {
	// loop
	for i := 0; i < 100; i++ {
		j := i % lSTCode
		ss := generateSubject(j)
		if e := tGenerateSubject(ss); e != nil {
			t.Error(e)
		}
	}

}

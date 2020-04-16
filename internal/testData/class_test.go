package testdata

import (
	"fmt"
	"testing"

	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

func testGenerateSubjects(s []models.Subject) (e error) {
	var sec models.Class // create a section
	sec.Subjects = s     // create all the subjects
	sec.CalCap()         // recalculate free periods

	// check the length of subject slice
	if len(sec.Subjects) == 0 {
		e = fmt.Errorf("> Error: not subjects assigned")
		return
	}

	// check no of free periods
	if sec.NFreePeriod != 0 {
		e = fmt.Errorf("> Error: Free periods=%v", sec.NFreePeriod)
		return
	}

	// check requirement for each subject
	for j, sub := range sec.Subjects {
		if sub.Req == 0 {
			e = fmt.Errorf("> Error: subject req=0 at index=%v. sID=%v", j, sub.ID.Bytes())
			return
		}
	}
	return
}

func TestGenerateSubject(t *testing.T) {
	i := 0                       // create an index of standered ID
	s := generateSubject(i)      // generate subjects
	e := testGenerateSubjects(s) // test subjects
	if e != nil {                // errror check
		t.Error(e)
	}
}

func TestGenerateSection(t *testing.T) {
	i := 0 // create an index of standered ID
	sec := generateSection(i)
	e := testGenerateClasses(sec)
	if e != nil {
		t.Error(e)
	}
}

func testGenerateClasses(cs models.Classes) (e error) {
	// check if the section are no 0
	if len(cs) == 0 {
		e = fmt.Errorf("> Error: no of classes is 0")
		return
	}

	// check each section's subjects
	for _, c := range cs {
		e = testGenerateSubjects(c.Subjects)
		if e != nil { // errror check
			return
		}
	}
	return
}
func TestGenerateClasses(t *testing.T) {
	var cs models.Classes
	generateClasses(&cs)

	e := testGenerateClasses(cs)
	if e != nil {
		t.Error(e)
	}
}

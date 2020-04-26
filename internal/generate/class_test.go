package generate

import (
	"fmt"
	"testing"

	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

func tGenerateClassL(cs []models.Class) error {
	// check if the section are no 0
	if len(cs) == 0 {
		return fmt.Errorf("> Error: no of classes is 0")
	}

	// check each section's subjects
	for _, c := range cs {
		e := tGenerateSubject(c.Subjects)
		if e != nil { // errror check
			return e
		}
	}
	return nil
}

func TestGenerateSection(t *testing.T) {
	i := 0 // create an index of standered ID
	sec := generateSection(i)
	e := tGenerateClassL(sec)
	if e != nil {
		t.Error(e)
	}
}

func TestGenerateClassL(t *testing.T) {
	var cs []models.Class
	generateClassL(&cs)

	e := tGenerateClassL(cs)
	if e != nil {
		t.Error(e)
	}
}

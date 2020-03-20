package models

import (
	"fmt"
	"testing"
)

// create a list of name
var names []Name = []Name{
	Name{"Vivek", "Singh", "Rawat"},
	Name{"Pooja", "", "Rawat"},
	Name{"Kapil ", "", "Gupta"},
	Name{"Vibhor", " ", "Mudgle"},
	Name{" Kabir", "", " Bhan"},
}

// create a toString of that list
var nameToString []string = []string{
	"Vivek Singh Rawat",
	"Pooja Rawat",
	"Kapil Gupta",
	"Vibhor Mudgle",
	"Kabir Bhan",
}

func testOneToString(i int) (e error) {
	names[i].Sanitizer()
	s := names[i].ToString()

	t := s == nameToString[i]

	if !t {
		e = fmt.Errorf(`> Error: name.ToString="%v", nameToString="%v" and t="%v". Where i="%v"`, s, nameToString[i], t, i)
	}
	return
}

func TestName_ToString(t *testing.T) {
	// Get length of the names
	l := len(names)

	for i := 0; i < l; i++ {
		e := testOneToString(i)

		if e != nil {
			t.Error(e)
		}
	}
}

package generate

import (
	"github.com/yumyum-pi/go-schoolScheduler/internal/utils"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

var nSec = utils.RangeInt{Min: 1, Max: 5} // range for generating random no. of sections

// template of classID with year and group info
var temCID models.ClassID

func initialization() {
	yr := make(utils.B256, 2)
	yr.Encode(2020)
	copy(temCID.Year[:], yr)
}

// generateSection return generate section if the given standard index i
func generateSection(i int) (sections []models.Class) {
	n := nSec.Random() // generate a random number of sections

	// adding standard bytes
	temCID.Standard = StanL[i]

	// loop for each section
	for noOfSec := 0; noOfSec < n; noOfSec++ {
		var sec models.Class // create a section
		// adding section bytes at the right position
		temCID.Section = byte(noOfSec + 1)

		sec.ID.Init(temCID.Bytes()) // creating a new ClassID

		sec.Subjects = generateSubject(i) // generate subject data
		sec.CalRemCap()                   // calculate the free periods

		sections = append(sections, sec) // append the the class to classes
	}
	return
}

// generateClassL return array of class
func generateClassL(c *[]models.Class) {
	initialization()
	// loop for each standard
	for i := range StanL {
		secs := generateSection(i)   // generate sections
		(*c) = append((*c), secs...) // add sections to the classes grp
	}
}

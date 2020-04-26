package generate

import (
	"github.com/yumyum-pi/go-schoolScheduler/internal/utils"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

var nSec = utils.RangeInt{Min: 1, Max: 5} // range for generating random no. of sections

// template of classID with year and group info
var temCID = [models.ClassIDBS]byte{2, 0, 2, 0, 0, 0, 0, 0, 0, 1}

// generateSection return generate section if the given standard index i
func generateSection(i int) (sections []models.Class) {
	n := nSec.Random() // generate a random number of sections

	// adding standard bytes
	temCID[models.YearBS] = StanL[i][0]
	temCID[models.YearBS+1] = StanL[i][1]

	// loop for each section
	for noOfSec := 0; noOfSec < n; noOfSec++ {
		var sec models.Class // create a section
		// adding section bytes at the right position
		temCID[models.YearBS+models.StandardBS+1] = byte(noOfSec + 1)

		sec.ID.Init(temCID) // creating a new ClassID

		sec.Subjects = generateSubject(i) // generate subject data
		sec.CalRemCap()                   // calculate the free periods

		sections = append(sections, sec) // append the the class to classes
	}
	return
}

// generateClassL return array of class
func generateClassL(c *[]models.Class) {
	// loop for each standard
	for i := range StanL {
		secs := generateSection(i)   // generate sections
		(*c) = append((*c), secs...) // add sections to the classes grp
	}
}

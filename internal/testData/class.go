package generate

import (
	"github.com/yumyum-pi/go-schoolScheduler/internal/utils"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

// STCodeL is a slice subject codes
var STCodeL = [][models.TypeBS]byte{
	{0, 0, 0, 1}, // english
	{0, 0, 0, 2}, // hindi
	{0, 0, 0, 3}, // maths
	{0, 0, 0, 4}, // science
	{0, 0, 0, 5}, // social science
	{0, 0, 0, 6}, // physical education
	{0, 0, 0, 7}, // art & craft
	{0, 0, 0, 8}, // computer science
	{0, 0, 0, 9}, // dance & music
}

var lSTCode = len(STCodeL)   // slice length of subject type code
var l2STCode = (lSTCode / 2) // half of slice lenghth of subject type code

// StanL is a list of standerds
var StanL = [][models.StanderdBS]byte{
	{0, 1},
	{0, 2},
	{0, 3},
	{0, 4},
	{0, 5},
	{0, 6},
	{0, 7},
	{0, 8},
	{0, 9},
	{1, 0},
	{1, 1},
	{1, 2},
}

var nSec = utils.RangeInt{Min: 1, Max: 5} // range for generating random no. of sections
var grp = [models.GroupBS]byte{0, 1}      // fixed group bytes
var yr = [models.YearBS]byte{2, 0, 0, 20} // fixed year bytes

// this is to return non main class index
var nonMainIndex = utils.RangeInt{Min: l2STCode, Max: lSTCode}

// genereateSubject return an array of subjects. i is the index of StanL, to get the standers byte info.
// TODO add test
func generateSubject(i int) (subjects []models.Subject) {
	rp := models.MaxCap // remain periods set to max capacity of the class

	// ranage for different type of classes
	nonMain := utils.RangeInt{Min: 4, Max: 6} // main class -eg: English, Maths, Science.
	main := utils.RangeInt{Min: 7, Max: 10}   // non main class-eg: Physical Education

	//loop for main subjects
	for j := 0; j < lSTCode; j++ {

		var s models.Subject // create a subject
		// create subjectID
		s.ID = models.SubjectID{
			Standerd: StanL[i],   // get the standerd of the index i
			Type:     STCodeL[j], // get the subject code of the index j
		}
		if j < l2STCode {
			s.Req = main.Random()          // assign random number of periods require
			rp -= s.Req                    // decrease the remaining period
			subjects = append(subjects, s) // add the subject to the list
			continue                       // iterate to the next index
		}

		nonMainP := nonMain.Random() // generate a random no. for assigning the remaining subjects

		// if remaining period is more than nonMainP
		if rp-nonMainP >= 0 {
			s.Req = nonMainP               // assign nonMainP to the subject require
			rp -= nonMainP                 // reduce the remaining period
			subjects = append(subjects, s) // add the subject to the list

			// break the loop if remaining periods is 0
			if rp == 0 {
				break
			}
		} else { // remaining periods is less than nonMainP
			s.Req = rp                     // assign the remainig periods to the subject
			rp = 0                         // resetting the remaining periods to 0
			subjects = append(subjects, s) // add the subject to the list
			break                          // break the loop
		}
	}
	return
}

// generateSection return generate section if the given standerd index i
func generateSection(i int) (sections []models.Class) {
	n := nSec.Random() // generate a random number of sections

	// loop for each section
	for noOfSec := 0; noOfSec < n; noOfSec++ {
		var sec models.Class // create a section

		secB := [2]byte{0, byte(noOfSec + 1)}  // create section byte
		sec.ID.Create(yr, StanL[i], secB, grp) // creating a new ClassID

		sec.Subjects = generateSubject(i) // generate subject data
		sec.CalRemCap()                   // calculate the free periods

		sections = append(sections, sec) // append the the class to classes
	}
	return
}

// generateClasses return array of class
func generateClasses(c *models.Classes) {
	// loop for each standerd
	for i := range StanL {
		secs := generateSection(i)   // generate sections
		(*c) = append((*c), secs...) // add sections to the classes grp
	}
}

package generate

import (
	"math/rand"

	"github.com/yumyum-pi/go-schoolScheduler/models"
)

type minMax struct {
	Min int
	Max int
}

func (m *minMax) GetRandomNo() int {
	return rand.Intn(m.Max-m.Min) + m.Min
}

// list of subject codes
var subjectTypeCode = [][4]byte{
	[4]byte{0, 0, 0, 1}, // english
	[4]byte{0, 0, 0, 2}, // hindi
	[4]byte{0, 0, 0, 3}, // maths
	[4]byte{0, 0, 0, 4}, // science
	[4]byte{0, 0, 0, 5}, // social science
	[4]byte{0, 0, 0, 6}, // physical education
	[4]byte{0, 0, 0, 7}, // art & craft
	[4]byte{0, 0, 0, 8}, // computer science
	[4]byte{0, 0, 0, 9}, // dance & music
}

// list of standerds
var standerds = [][2]byte{
	[2]byte{0, 5},
	[2]byte{0, 6},
	[2]byte{0, 7},
	[2]byte{0, 8},
}

var noOfSection = minMax{1, 4}
var group = [2]byte{0, 1}

// genereateSubject return an array of subjects
func generateSubject(i int) (subjects []models.Subject) {
	standerd := &standerds[i]     // get the standerd info from the array
	remainingNonMainClasses := 10 // number of non main classes

	var nonMainClassMinMax minMax = minMax{1, 3}

	//loop for each subject
	for j := 0; j < 9; j++ {
		// create a subject
		var subject models.Subject

		// create subjectID
		subject.ID = models.SubjectID{
			Standerd: *standerd,
			Type:     subjectTypeCode[j],
		}
		// main 5 subject
		if j < 5 {
			subject.ReqClasses = 6
			subjects = append(subjects, subject)
		} else {
			// generate a random no. for assigning the remaining subjects
			nonMainClass := nonMainClassMinMax.GetRandomNo()

			// check of remaining-non-main-classes more then non-main-class
			if remainingNonMainClasses-nonMainClass >= 0 {
				subject.ReqClasses = nonMainClass
				remainingNonMainClasses -= nonMainClass
				// check of remaining-non-main-classes more then non-main-class
			} else if remainingNonMainClasses > 0 {
				// assign all the remainig classes to the subject
				subject.ReqClasses = remainingNonMainClasses
				remainingNonMainClasses = 0
				// the no of non-main-classes in 0
			} else {
				subject.ReqClasses = 0
			}

			subjects = append(subjects, subject)
		}
	}
	return
}

func generateSection(i int) (classes []models.Class) {
	// generate a random number NoOfSection
	standerd := &standerds[i]
	n := noOfSection.GetRandomNo()

	// loop for each section
	for noOfSec := 0; noOfSec < n; noOfSec++ {
		var class models.Class               // create a new class
		sec := [2]byte{0, byte(noOfSec + 1)} // create section data
		var id models.ClassID                // generate ClassID
		id = id.Create(*standerd, sec, group)

		class.ID = id                       // assign the classID
		class.Subjects = generateSubject(i) // generate subject data

		classes = append(classes, class) // append the the class to classes
	}

	return
}

// generateClasses return array of class
func generateClasses() (classes []models.Class) {
	// loop for each standerd
	for i := range standerds {
		secs := generateSection(i)         // generate sections
		classes = append(classes, secs...) // add sections to the classes grp
	}
	return
}

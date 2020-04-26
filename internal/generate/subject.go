package generate

import (
	"github.com/yumyum-pi/go-schoolScheduler/internal/utils"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

// STCodeL is a slice subject codes
var STCodeL = []byte{
	1, // english
	2, // hindi
	3, // maths
	4, // science
	5, // social science
	6, // physical education
	7, // art & craft
	8, // computer science
	9, // dance & music
}

var lSTCode = len(STCodeL)       // slice length of subject type code
var l2STCode = (lSTCode / 2) + 1 // half of slice length of subject type code

// StanL is a list of standards
var StanL = []byte{
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
	10,
	11,
	12,
}

// genereateSubject return an array of subjects. i is the index of standard list(StanL) - to get the standers byte info.
func generateSubject(i int) (subjects []models.Subject) {
	rp := models.MaxCap // remain periods set to max capacity of the class

	// ranage for different type of classes
	nonMain := utils.RangeInt{Min: 1, Max: 3} // main class -eg: English, Maths, Science.
	main := utils.RangeInt{Min: 6, Max: 9}    // non main class -eg: Physical Education

	//create subjects and assign requirement for the main subjects
	for j := 0; j < lSTCode; j++ {
		var s models.Subject // create a subject
		// create subjectID
		s.ID = models.SubjectID{
			Standard: StanL[i],   // get the standard of the index i
			Type:     STCodeL[j], // get the subject code of the index j
		}
		subjects = append(subjects, s) // add the subject to the list

		// assign main periods
		if j < l2STCode {
			subjects[j].Req = main.Random() // assign random number of periods require
			rp -= subjects[j].Req           // decrease the remaining period
			continue                        // iterate to the next index
		}
	}

	nmi := l2STCode // index for the non main subjects
	// loop for assigning non main subject
	for rp > 0 {
		nonMainP := nonMain.Random() // generate a random no. for assigning the remaining subjects

		// remaining period is more than required periods
		// add the requirement to the subject
		if rp-nonMainP >= 0 {
			rp -= nonMainP                // reduce the remaining period
			subjects[nmi].Req += nonMainP // add the requirement to the subject

			// check nmi exceeds the index
			if nmi == lSTCode-1 {
				nmi = l2STCode // loop the index from l2STCode
				continue
			}
			nmi++
			continue
		}

		// remaining periods is less than nonMainP
		subjects[nmi].Req += rp // add the requirement to the subject
		rp = 0                  // resetting the remaining periods to 0
		break                   // break the loop
	}
	return
}

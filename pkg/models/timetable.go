package models

import "fmt"

// NDays is the number of days in a week
const NDays = 6

// NPeriods is the number of periods in a day
const NPeriods = 8

// MaxCap is the max number of period in a week
const MaxCap = NDays * NPeriods

// Period is a single cell in the timetable
type Period struct {
	ClassID   ClassID   // assigned class's ID
	SubjectID SubjectID // assigned subject's ID
	TeacherID TeacherID // assigned teacher's ID
}

// Bytes return the byte array of period data
func (p *Period) Bytes() (b PeriodB) {
	b[0] = (*p).ClassID.Year[0]   // 0 - Class year
	b[1] = (*p).ClassID.Year[1]   // 1
	b[2] = (*p).ClassID.Standard  // 2 - Class standard
	b[3] = (*p).ClassID.Section   // 3 - Class section
	b[4] = (*p).SubjectID.Type    // 4 - Subject type
	b[5] = (*p).TeacherID.Year[0] // 5 - Teacher year
	b[6] = (*p).TeacherID.Year[1] // 6
	b[7] = (*p).TeacherID.JoinNo  // 7 - Teacher join no
	return
}

// Init return the byte array of period data
func (p *Period) Init(b PeriodB) {
	(*p).ClassID.Year[0] = b[0]    // 0 - Class year
	(*p).ClassID.Year[1] = b[1]    // 1
	(*p).ClassID.Standard = b[2]   // 2 - Class standard
	(*p).ClassID.Section = b[3]    // 3 - Class section
	(*p).SubjectID.Standard = b[2] // 3 - Subject section
	(*p).SubjectID.Type = b[4]     // 4 - Subject type
	(*p).TeacherID.Year[0] = b[5]  // 5 - Teacher year
	(*p).TeacherID.Year[1] = b[6]  // 6
	(*p).TeacherID.JoinNo = b[7]   // 7 - Teacher join no
}

// PID2TTID return timetable ID from the class peroid ID
func PID2TTID(classIndex, pID int) int {
	return (classIndex*MaxCap + pID)
}

// IsAssigned return true if period is not assigned
func (p *Period) IsAssigned() bool {
	if p.ClassID != (ClassID{}) || p.TeacherID != (TeacherID{}) {
		return true
	}
	return false
}

// IsAssigned return true if periodB is not assigned
func (pb *PeriodB) IsAssigned() bool {
	var cID ClassIDB
	var tID TeacherIDB
	copy(cID[:], (*pb)[:ClassIDBS])
	copy(tID[:], (*pb)[ClassIDBS+1:])
	eCID := ClassIDB{}
	eTID := TeacherIDB{}

	if cID != eCID || tID != eTID {
		return true
	}
	return false
}

// TimeTables is struct to store timetable data
type TimeTables []Period

// Bytes converts timetable to slice of array
// TODO add test
func (tt *TimeTables) Bytes() []byte {
	l := len(*tt)
	byteL := l * PeriodBS
	b := make([]byte, byteL)
	var pb PeriodB
	for i, p := range *tt {
		pb = p.Bytes()
		b[i] = pb[0]
		b[i+1] = pb[1]
		b[i+2] = pb[2]
		b[i+3] = pb[3]
		b[i+4] = pb[4]
		b[i+5] = pb[5]
		b[i+6] = pb[6]
		b[i+7] = pb[7]
	}
	return b
}

// ClassL2TT return a timetable for the class list
func ClassL2TT(cs *[]Class) TimeTables {
	cl := len(*cs) // class length
	// create timetable with all the class periods combined
	tt := make(TimeTables, cl*MaxCap)

	// assign variable to be used
	lPID := 0     // to store last PeriodID of a subject
	p := Period{} // store Peroid
	pID := 0

	// loop through each class
	for ci, c := range *cs {
		pID = 0 // set the 1st periodID

		// loop through all the subject assigned
		for _, s := range c.Subjects {
			// calculate the class periods if of a subject
			lPID = pID + s.Req

			// create a new period
			p = Period{
				ClassID:   c.ID,
				SubjectID: s.ID,
				TeacherID: s.TeacherID,
			}

			// loop until the pID == lPID
			for pID < lPID {
				// generate geneID for the period
				ttID := PID2TTID(ci, pID)
				// assign the period
				tt[ttID] = p
				// increase the pID and continue the loop
				pID++
			}
		}
	}
	return tt
}

// Print write the value to the console
func (tt *TimeTables) Print() {
	for _, p := range *tt {
		fmt.Printf(
			"cID=%v\t\tsID=%v\t\ttID=%v \n",
			p.ClassID.Bytes(),
			p.SubjectID.Bytes(),
			p.TeacherID.Bytes(),
		)
	}
}

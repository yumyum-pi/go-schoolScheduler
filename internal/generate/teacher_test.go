package generate

import (
	"fmt"
	"testing"

	rl "github.com/yumyum-pi/go-schoolScheduler/internal/requestlist"
	"github.com/yumyum-pi/go-schoolScheduler/internal/utils"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

func TestGenerateTeacherID(t *testing.T) {
	const nIndex = 10         // number for iteration
	var tL []models.TeacherID // create a list teacherID

	for i := 0; i < nIndex; i++ {
		// loop and populate teacherID list
		for j := 0; j < nIndex; j++ {
			// add the generated teacherID to the list
			tL = append(tL, gTeacherID())
		}

		// check if the teacherID matches
		for j := 0; j < nIndex; j++ {
			// get the next index
			k := j + 1
			// k should be in bound of the index count
			if k < nIndex {
				// the teacherIDs should not match
				if tL[j] == tL[k] {
					t.Errorf(
						"> Error: tID1=%v && tID2=%v are equal at i=%v j=%v k=%v",
						tL[j],
						tL[k],
						i,
						j,
						k,
					)
				}
			}
		}

	}
}

func tGTeacher(i int) error {
	sID := models.TSubjectIDL[i]
	cID := models.TClassIDL[i]
	req := i
	t := gTeacher(sID, cID, req)
	eReq := models.TeacherCap - req

	// check capacity
	if t.Capacity != eReq {
		return fmt.Errorf(
			"> Error: t.Capacity=%v, exprected eReq=%v at i=%v",
			t.Capacity,
			eReq,
			i,
		)
	}

	// check subject can teacher list length
	if len(t.SubjectCT) == 0 {
		return fmt.Errorf("> Error: subjectCT should not be empty at i=%v", i)
	}
	// check subject to teacher
	if t.SubjectCT[0] != sID {
		return fmt.Errorf("> Error: subject not assigned")
	}
	return nil
}
func TestGTeacher(t *testing.T) {
	l := len(models.TSubjectIDL)

	for i := 0; i < l; i++ {
		if e := tGTeacher(i); e != nil {
			t.Error(e)
		}
	}
}

// Test the following
//	-max value
//	-diff value
// TODO unit testing
func tCDistributed(diff, max int, ts *models.Teachers) error {
	return nil
}

func TestCDistributed(t *testing.T) {
	var cs []models.Class // create a class for test
	generateClassL(&cs)   // add info to the class

	rls := make(rl.Subject) //create new subject request list
	rls.Init(&cs)           // assign the value

	for SIDT, crl := range rls {
		if len(crl) == 0 {
			// skip if class required list is empty
			continue
		}

		tL := cTeacherL(crl.TotalReq()) // create the creates list

		autoAssignM1(SIDT, &crl, &tL) // assign subjects to the teachers

		diff, max := cDistributed(&tL)

		// test if assigned properly
		if e := tCDistributed(diff, max, &tL); e != nil {
			t.Error(e)
		}

	}
}

// Test the following
//	-assigned class is not empty
//	-correct subject id- standard, type
//	-teacher capacity is correct
//	-all classes are assigned
func tAutoAssignM1(SIDT byte, rlc *rl.Class, tL *models.Teachers) error {
	crl := make(rl.Class) // store classID and req assigned to teacher

	for i, T := range *tL {
		// assigned class should not be empty
		if len(T.AClassL) == 0 {
			return fmt.Errorf(
				"> Error: tID=%v, assigned class should not be empty at i=%v",
				T.ID.Bytes(),
				i,
			)
		}

		cap := models.TeacherCap // calculate the teachers capacity

		// loop though each assigned class
		for _, ac := range T.AClassL {
			cID := ac.ClassID.Bytes() // get the classID
			req, ok := crl[cID]       // get the no of periods required
			if !ok {                  // check if cID does not exist
				req = 0 // assign initial value
			}

			req += ac.Assigned // add the assigned no. of periods
			crl[cID] = req     // re-map the req to cID
			cap -= ac.Assigned // subreact the assigned no. of periods

			// check if subjectID.Type is correct
			if ac.SubjectID.Type != SIDT {
				return fmt.Errorf(
					"> Error: tID=%v, sID=%v is incorrect SIDT=%vat i=%v",
					T.ID.Bytes(),
					ac.SubjectID.Bytes(),
					SIDT,
					i,
				)
			}
			// check if subjectID.Standard is correct
			if ac.SubjectID.Standard != ac.ClassID.Standard {
				return fmt.Errorf(
					"> Error: tID=%v, sID=%v is incorrect stan=%vat i=%v",
					T.ID.Bytes(),
					ac.SubjectID.Bytes(),
					ac.ClassID.Standard,
					i,
				)
			}
		}

		// check teacher capacity
		if T.Capacity != cap {
			return fmt.Errorf(
				"> Error: tID=%v t.Cap=%v exprected=%v at i=%v",
				T.ID.Bytes(),
				T.Capacity,
				cap,
				i,
			)
		}
	}

	// check the assigned classes
	for cID, req := range *rlc {
		// get the required periods from the classes assigned to teacher
		r, ok := crl[cID]
		if !ok { // the cID should exist
			return fmt.Errorf("> Error: cID=%v not found", cID)
		}

		if r != req { // the required periods should match
			return fmt.Errorf("> Error: cID=%v req=%v not equal to r=%v", cID, r, req)
		}
	}

	return nil
}
func TestAutoAssignM1(t *testing.T) {
	var cs []models.Class // create a class for test
	generateClassL(&cs)   // add info to the class

	rls := make(rl.Subject) //create new subject request list
	rls.Init(&cs)           // assign the value

	for SIDT, crl := range rls {
		if len(crl) == 0 {
			// skip if class required list is empty
			continue
		}

		tL := cTeacherL(crl.TotalReq()) // create the creates list

		autoAssignM1(SIDT, &crl, &tL) // assign subjects to the teachers
		// test if assigned properly
		if e := tAutoAssignM1(SIDT, &crl, &tL); e != nil {
			t.Error(e)
		}

	}
}

// TODO unit testing
func TestASubjectCT(t *testing.T) {

}

// Test the following:
//	-no. of teachers
//	-has teacherID
//	-capacity of each teacher
func tCTeacherL(tr int) error {
	ts := cTeacherL(tr)
	nT := (tr / models.TeacherCap) + 1 // calculate the no max no of teacher

	// length check
	if len(ts) != nT {
		return fmt.Errorf(
			"> Error: len(ts)=%v nT=%v, expected no. of teachers",
			len(ts),
			nT,
		)
	}
	for _, t := range ts {
		// teacherID should not be empty
		if t.ID != (models.TeacherID{}) {
			return fmt.Errorf(
				"> Error: t.ID=%v should not be empty",
				t.ID,
			)
		}
		// capacity should be full
		if t.Capacity != models.TeacherCap {
			return fmt.Errorf(
				"> Error: t.Cap=%v expected %v",
				t.Capacity,
				models.TeacherCap,
			)
		}
	}
	return nil
}
func TestCTeacherL(t *testing.T) {
	var tRange utils.RangeInt
	tRange.Min, tRange.Max = 3, 10

	for i := 0; i < 100; i++ {
		// generate a number
		tr := tRange.Random()
		tr += tr * models.TeacherCap

		if e := tCTeacherL(tr); e != nil {
			t.Error(e)
		}
	}
}

// TODO check test
func TestGTeacherL(t *testing.T) {
	var cs []models.Class // create a class for test
	generateClassL(&cs)   // add info to the class

	// get subject request list
	tL := gTeacherLM2(&cs)
	crl := make(rl.Class)
	// check subjectCT
	for i, T := range tL {
		if len(T.SubjectCT) == 0 {
			t.Fatalf(
				"> Error: tID=%v, subjectCT should not be 0 at i=%v",
				T.ID.Bytes(),
				i,
			)
		}

		if len(T.AClassL) == 0 {
			t.Fatalf(
				"> Error: tID=%v, AClassL should not be 0 at i=%v",
				T.ID.Bytes(),
				i,
			)
		}
		cap := 0
		for _, ac := range T.AClassL {
			// check if ac in sct
			if !T.CanTeach(ac.SubjectID) {
				t.Fatalf(
					"> Error: tID=%v, subjectCT=%v should be found at i=%v",
					T.ID.Bytes(),
					ac.SubjectID.Bytes(),
					i,
				)
			}
			// check if the class exit int the crl list
			c, ok := crl[ac.ClassID.Bytes()]
			if !ok {
				c = 48
			}
			c -= ac.Assigned
			crl[ac.ClassID.Bytes()] = c
			cap += ac.Assigned
		}

		cap = models.TeacherCap - cap

		if T.Capacity != cap {
			t.Fatalf(
				"> Error: tID=%v t.Cap=%v exprected=%v at i=%v",
				T.ID.Bytes(),
				T.Capacity,
				cap,
				i,
			)
		}

	}

	for _, c := range cs {
		if c.Capacity != crl[c.ID.Bytes()] {
			t.Errorf(
				"> Error: cID=%v c.Cap=%v crl=%v should match",
				c.ID.Bytes(),
				c.Capacity,
				crl[c.ID.Bytes()],
			)
		}
	}
}

package generate

import (
	c "crypto/rand"

	rl "github.com/yumyum-pi/go-schoolScheduler/internal/requestlist"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

// Generate random teacherID
func gTeacherID() (tID models.TeacherID) {
	tID.Year = [models.YearBS]byte{2, 0, 2, 0} // fixed data
	// create a new join id
	join := make([]byte, 4)
	c.Read(join) // assign new values

	// copy join data
	copy(tID.JoinNo[:], join[:])
	return
}

// createTeacher returns a new randomly generated teacher
func gTeacher(sID models.SubjectID, cID models.ClassID, req int) (t models.Teacher) {
	t.ID = gTeacherID()                   // generate random teacherID
	t.Capacity = models.TeacherCap        // capacity
	t.SubjectCT = []models.SubjectID{sID} // Add to subject could teach list
	t.AssignClass((cID), sID, req)        // Assign class

	return t
}

// generate teacher from the given request list
func gTeacherLM1(rls *rl.Subject) (ts models.Teachers) {
	// Loop through all the Teacher Request List
	for SIDT, cc := range *rls {
		for CID, req := range cc {
			var cID models.ClassID
			var sID models.SubjectID
			cID.Init(CID)
			sID.Type = SIDT
			sID.Standard = cID.Standard
			// check if teacher list is empty
			if len(ts) == 0 {
				// create a new and teacher to an empty teachers list
				t := gTeacher(sID, cID, req) // Create a new Teacher
				ts = append(ts, t)           // Add the teacher to the teacher list
				continue                     // skip to the next iteration
			}

			var ifAssigned bool // to check if the subject is assigned to a teacher

			TypeMatchL := ts.FindBySubType(SIDT) // find teacher using subjectID

			// check if type match list is empty
			if len(TypeMatchL) == 0 {
				// The list is empty
				t := gTeacher(sID, cID, req) // Create a new Teacher
				ts = append(ts, t)           // Add the teacher to the teacher list
				ifAssigned = true
				continue // skip to the next iteration
			}
			// The list in !empty
			// loop through the list
			for _, iT := range TypeMatchL {
				// check if the teacher has the capacity to teach the class
				diff := ts[iT].AssignClass(cID, sID, req)

				if diff > 0 {
					// Add the Subject to the teacher
					// check if teacher has the subjectID
					if !ts[iT].CanTeach(sID) {
						// add subject to can teach list
						ts[iT].SubjectCT = append(ts[iT].SubjectCT, sID)
					}
					ifAssigned = true
					break // exit the loop
				}
			}

			// check if subject not assigned
			if !ifAssigned {
				// create a new teacher
				t := gTeacher(sID, cID, req) // Create a new Teacher
				ts = append(ts, t)           // Add the teacher to the teacher list
			}
		}

	}
	return
}

// check if the teachers list needs is distribution by comparing
// the sum of all the teacher's capacity to the capacity of the last
// teacher
func cDistributed(ts *models.Teachers) (int, int) {
	l := len(*ts) // length of the slice
	// calculate the average
	max := 0 // max no. of periods assigned to a teacher
	a := 0
	for _, t := range *ts {
		for _, ca := range t.AClassL {
			if max < ca.Assigned {
				max = ca.Assigned
			}
		}
		a += t.Capacity
	}
	a /= l // calculate the average

	// calculate the difference between last teacher's capacity and the average capacity
	diff := (*ts)[l-1].Capacity - a
	return diff, max
}

// TODO add test
func aaRedistributed(diff, max int, ts *models.Teachers) error {
	l := len(*ts) // length of the slice
	// loop till the diff is bigger then 10
	for diff > max {
		// loop through all teacher exprect the last
		for i := 0; i < l-1; i++ {
			// get the teacher
			t2 := (*ts)[i]
			t2cll := len(t2.AClassL)
			t2cl := t2.AClassL[t2cll-1] //
			// assign the last class and subject for the teacher to the last teacher
			d := (*ts)[l-1].AssignClass(t2cl.ClassID, t2cl.SubjectID, t2cl.Assigned)
			if d >= 0 {
				t2.AClassL = t2.AClassL[:t2cll-1] // remove the class for the teacher list
				t2.Capacity += t2cl.Assigned      // free the unassigned periods
				(*ts)[i] = t2                     // update the teachers list

				// decrease the diff
				diff -= t2cl.Assigned
				// check if lower then 10
				if diff <= max {
					break
				}
			}

		}
	}

	return nil
}

// autoAssignM1 assigned the given subjects to the given teachers
func autoAssignM1(SIDT [models.TypeBS]byte, cc *rl.Class, ts *models.Teachers) {
	var nCID models.ClassID
	var nSID models.SubjectID
	nSID.Type = SIDT
	// loop through the subjectIDs

	for cID, req := range *cc {
		ifAssigned := false
		nCID.Init(cID)
		nSID.Standard = nCID.Standard
		for j := range *ts {
			// check if teacher has capacity
			diff := (*ts)[j].AssignClass(nCID, nSID, req)
			if diff < 0 {
				// not assigned
				continue
			}
			// assigned the teacher
			ifAssigned = true
			break // exit the current loop
		}
		if !ifAssigned {
			// add another teacher to the teachers list
			var t models.Teacher // create a new teacher
			t.ID = gTeacherID()  // generate random teacherID
			t.Capacity = models.TeacherCap
			t.AssignClass(nCID, nSID, req)
			(*ts) = append((*ts), t)
		}

	}
}

// add subject that the teacher could teach
func aSubjectCT(ts *models.Teachers) {
	// assign subjectCT
	for i, t := range *ts {
		for _, ac := range t.AClassL {
			// check if ac in sct
			if !t.CanTeach(ac.SubjectID) {
				t.SubjectCT = append(t.SubjectCT, ac.SubjectID)
			}
		}
		(*ts)[i] = t
	}
}

// genreate the minimum no of teacher required from the subject code
//  tr- total required periods of the subject
func cTeacherL(tr int) models.Teachers {
	nt := (tr / models.TeacherCap) + 1 // calculate the no max no of teacher
	tL := make(models.Teachers, nt)    // make teacher list for the subject type

	// loop through the no of teacher required for the subject type
	// create new teachers and add to the teachers list
	for i := 0; i < nt; i++ {
		var t models.Teacher           // create a new teacher
		t.ID = gTeacherID()            // generate random teacherID
		t.Capacity = models.TeacherCap // capacity
		tL[i] = t                      // add to the list
	}

	return tL
}

// generate teacher from the given request list
func gTeacherLM2(cs *[]models.Class) (ts models.Teachers) {
	rls := make(rl.Subject) //create new subject request list
	rls.Init(cs)            // assign the value

	// Loop through all the Teacher Request List
	for SIDT, cc := range rls {
		// check if cc is not nil
		if len(cc) == 0 {
			// skip this sID
			continue
		}
		tL := cTeacherL(cc.TotalReq()) // create the creates list

		autoAssignM1(SIDT, &cc, &tL) // assign the subject sequentially

		// check if the classes are distributed evenly amongst the teachers
		if diff, max := cDistributed(&tL); diff > max {
			// distribute the classes evenly amongst the teachers
			aaRedistributed(diff, max, &tL)
		}

		aSubjectCT(&tL) // assign subjectCT to each teacher

		ts = append(ts, tL...)
	}
	// a map to shore index of the given classID in the cs classes slice
	cIDIndex := make(rl.Class)
	// populate the the cIDIndex
	for i, c := range *cs {
		cIDIndex[c.ID.Bytes()] = i
	}
	// assign the teacher to the classes
	for _, t := range ts {
		// loop through assigned classes
		for _, c := range t.AClassL {
			cIndex := cIDIndex[c.ClassID.Bytes()]
			sIndex := (*cs)[cIndex].Subjects.FindByID(c.SubjectID.Bytes())
			// add the teacher to the subject
			(*cs)[cIndex].Subjects[sIndex].TeacherID = t.ID
		}

	}
	return
}

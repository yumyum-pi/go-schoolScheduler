package autoassign

import "github.com/yumyum-pi/go-schoolScheduler/pkg/models"

// Teachers assigns teacher to the class
// TODO Write test for this function
func Teachers(cs *[]models.Class, t *models.Teachers) (emptySubs []models.SubjectID) {
	// loop through classes
	for _, c := range *cs {
		// - loop through subjects
		for _, s := range c.Subjects {
			// -- for each subject select a teacher
			tMatchIndex := (*t).FindBySub(s.ID)

			// throw error an move on to the next subject if no match found
			if len(tMatchIndex) > 0 {
				// TODO through error
				// to emptySubs
				emptySubs = append(emptySubs, s.ID)
				continue // move on to the next subjectID
			}

			assigned := false // to check if assigned

			// loop through matches
			for _, i := range tMatchIndex {
				tID := tMatchIndex[i] // get the teacher id

				diff := (*t)[tID].AssignClass(c.ID, s.ID, s.Req) // try to assign the teacher to the class

				// check capacity
				// if diff is negative means that the requested period was greater than the capacity
				if diff < 0 {
					continue // skip to next match
				}

				// assign the teacher to the class
				// error is not check
				c.AssignTeacher(s.ID, (*t)[tID].ID)
				assigned = true

				// match found
				// exit the loop
				break
			}

			// add subject id to empty list if teacher is not assigned
			if !assigned {
				emptySubs = append(emptySubs, s.ID)
			}
		}
	}
	return
}

// TeachersM2 assigns teacher to the class
func TeachersM2(cs *[]models.Class, t *models.Teachers) (emptySubs []models.SubjectID) {
	// create a srl
	return
}

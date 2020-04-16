package testdata

import (
	c "crypto/rand"

	rl "github.com/yumyum-pi/go-schoolScheduler/internal/requestlist"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

// Generate random teacherID
func generateTeacherID() (id models.TeacherID) {
	id.Year = [4]byte{2, 0, 2, 0}
	key := make([]byte, 4)

	c.Read(key)

	id.JoinNo[0], id.JoinNo[1], id.JoinNo[2], id.JoinNo[3] = key[0], key[1], key[2], key[3]
	return
}

// createTeacher return a new randomly generated teacher
func createTeacher(subjectID models.SubjectID, req int) (t models.Teacher) {
	var class models.ClassID // create blank classID

	t.ID = generateTeacherID()                  // generate random teacherID
	t.Capacity = models.TeacherCap              // capacity
	t.SubjectCT = []models.SubjectID{subjectID} // Add to subject could teach list
	t.AssignClass(class, subjectID, req)        // Assign class

	return t
}

// generateTeacher generate teacher from the given request list
// TODO add test
func generateTeacherList(trl *rl.TeacherRL) (teacherList models.Teachers) {

	// Loop through all the Teacher Request List
	for _, t := range *trl {
		subjectID, req := t.SubjectID, t.Req

		// check if teacherList is empty
		if len(teacherList) == 0 {
			teacher := createTeacher(subjectID, req)                // Create a new Teacher
			teacher.AssignClass((models.ClassID{}), subjectID, req) // assigning the subject to the teacher
			teacherList = append(teacherList, teacher)              // Add the teacher to the teacherList
			//fmt.Printf("> !List:\tname\"%v\"\tCreated a new teacher when list is empty.\n", teacher.Name)
		} else {
			var ifAssigned bool // to check if the subject is assigned to a teacher

			// check if found any matches
			teacherTypeMatch := teacherList.FindBySubType(&subjectID) // find teacher using subjectID
			ttLen := len(teacherTypeMatch)

			// check if ttLen empty
			if ttLen == 0 {
				// The list is empty
				teacher := createTeacher(subjectID, req)                // Create a new teacher
				teacher.AssignClass((models.ClassID{}), subjectID, req) // assigning the subject to the teacher
				teacherList = append(teacherList, teacher)              // Add the teacher to the teacherList
				ifAssigned = true
				//fmt.Printf("> !Match:\tname\"%v\"\tCreated a new teacher when list is empty.\n", teacher.Name)
			} else {
				// The list in !empty
				// loop through the list
				for _, iT := range teacherTypeMatch {
					//	- check if the teacher has the capacity to teach the class
					diff := teacherList[iT].AssignClass(models.ClassID{}, subjectID, req)

					if diff > 0 {
						// Add the Subject to the teacher
						ifAssigned = true
						teacherList[iT].SubjectCT = append(teacherList[iT].SubjectCT, subjectID)
						//fmt.Printf("> Assigned:\tname\"%v\"\tSubect assigned to the teacher. \n", teacherList[iT].Name)
						break // exit the loop
					}
				}
			}

			// check if subject assigned
			if !ifAssigned {
				// create a new teacher
				teacher := createTeacher(subjectID, req)                // Create a new Teacher
				teacher.AssignClass((models.ClassID{}), subjectID, req) // assigning the subject to the teacher
				teacherList = append(teacherList, teacher)              // Add the teacher to the teacherList
				//fmt.Printf("> !ifAssigned:\tname\"%v\"\tCreated a new teacher.\n", teacher.Name)
			}

		}
	}

	return
}
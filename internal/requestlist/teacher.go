package requestlist

import "github.com/yumyum-pi/go-schoolScheduler/pkg/models"

// TeacherRLE is struct of requirement of a Teacher
// This struct contains the subjectID,
// list of class ID that require this subject
// Total No. of requirement
type TeacherRLE struct {
	SubjectID models.SubjectID
	ClassID   models.ClassID
	Req       int
}

// Init assign the proper value of the struct
func (t *TeacherRLE) Init(sID [models.SubjectIDBS]byte, cID [models.ClassIDBS]byte, req int) {
	(*t).SubjectID.Init(sID)
	(*t).ClassID.Init(cID)
	(*t).Req = req

}

// TeacherRL is a slice of Teacher Requried List Element
type TeacherRL []TeacherRLE

// Add Teacher Request List Element to the list
func (trl *TeacherRL) Add(t TeacherRLE) {
	*trl = append(*trl, t)
}

// Create the list from the given Subject
func (trl *TeacherRL) Create(rls *Subject) {
	// loop through subject request list
	for sID, cs := range *rls {
		// loop through each classes
		for cID, req := range cs {
			var t TeacherRLE      // create a new teacher request list element
			t.Init(sID, cID, req) // assign values to teacher request list element
			trl.Add(t)            // add teacher request list element to teacher request list
		}
	}
}

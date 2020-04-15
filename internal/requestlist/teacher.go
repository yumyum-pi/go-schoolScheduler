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
func (t *TeacherRLE) Init(s *SubjectRLE, i int) {
	(*t).SubjectID = (*s).SubjectID
	(*t).ClassID = (*s).Classes[i].ID
	(*t).Req = (*s).Classes[i].Req

}

// TeacherRL is a slice of Teacher Requried List Element
type TeacherRL []TeacherRLE

// Add Teacher Request List Element to the list
func (trl *TeacherRL) Add(t TeacherRLE) {
	*trl = append(*trl, t)
}

// Create the list from the given SubjectRL
func (trl *TeacherRL) Create(srl *SubjectRL) {
	// loop srl
	for _, s := range *srl {
		// loop through each classes
		for i := range s.Classes {
			var t TeacherRLE
			t.Init(&s, i)
			trl.Add(t)
		}
	}
}

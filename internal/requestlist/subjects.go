package requestlist

import (
	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

// Subject is an slice of subject request elements
type Subject map[[models.SubjectIDBS]byte]Class

// Init populate subject request list from the given classes
func (rls *Subject) Init(cs *models.Classes) {
	// loop through classes
	for _, c := range *cs {
		// loop through subjects
		for _, sub := range c.Subjects {
			// get the class request list
			cc, ok := (*rls)[sub.ID.Bytes()]
			// check if the class request list exist of the given subjectID
			if !ok {
				cc = make(Class)
			}
			cc[c.ID.Bytes()] = sub.Req  // assign subject request to the classID
			(*rls)[sub.ID.Bytes()] = cc // ressign the class request list to subjectID
		}
	}
	return
}

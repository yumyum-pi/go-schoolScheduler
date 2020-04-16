package requestlist

import (
	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

// Subject is an slice of subject request elements
type Subject map[[models.SubjectIDBS]byte]Class

// Init return added the class to subject request list
// TODO add test
func (rls *Subject) Init(cs *models.Classes) {
	// loop through classes
	for _, c := range *cs {
		// loop through subjects
		for _, sub := range c.Subjects {
			cc, ok := (*rls)[sub.ID.Bytes()] // the class request list
			// check if cc is not nil
			if !ok {
				cc = make(Class)
			}
			cc[c.ID.Bytes()] = sub.Req  // assign subject request to the classID
			(*rls)[sub.ID.Bytes()] = cc // ressign the class request list to subjectID
		}
	}
	return
}

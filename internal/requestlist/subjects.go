package requestlist

import (
	"fmt"

	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

// Subject is an slice of subject request elements
type Subject map[byte]Class

// Init populate subject request list from the given classes
func (rls *Subject) Init(cs *[]models.Class) {
	// loop through classes
	for _, c := range *cs {
		// loop through subjects
		for _, sub := range c.Subjects {
			// get the class request list
			crl, ok := (*rls)[sub.ID.Type]
			// check if the class request list exist of the given subjectID
			if !ok {
				crl = make(Class) // make a new class requrest list
			}
			crl[c.ID.Bytes()] = sub.Req // assign subject request to the classID
			(*rls)[sub.ID.Type] = crl   // ressign the class request list to subjectID
		}
	}
	return
}

// Print write subject request list value to the console
func (rls *Subject) Print() {
	for sID, crl := range *rls {
		fmt.Printf("> sID=%v\n", sID)
		for cID, req := range crl {
			fmt.Printf("cID=%v\treq=%v\n", cID, req)
		}
	}
}

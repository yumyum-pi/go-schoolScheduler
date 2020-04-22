package requestlist

import (
	"testing"

	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

func TestSubject_Init(t *testing.T) {
	srl := make(Subject)      // create a new subject request list
	srl.Init(&models.TClassL) // run the func

	// loop through all the subjects in the map
	for sID, crl := range srl {
		// check if subjectID is present
		cCRL, ok := TSRL[sID]
		if !ok {
			t.Errorf("> Error: sID=%v not found", sID)
			break
		}
		// check if the number of elements are same
		if len(cCRL) != len(crl) {
			t.Errorf("> Error: crl=%v != cCRL=%v at sID=%v  ", len(cCRL), len(crl), sID)
			break
		}

		// loop through and check if all the classes
		for cID, req := range cCRL {
			// check classID exist
			tReq, ok := crl[cID]
			if !ok {
				t.Errorf("> Error: cID=%v not found at sID=%v", cID, sID)
				break
			}

			// check if request are same
			if tReq != req {
				t.Errorf("> Error: cID=%v req=%v != tReq=%v at sID=%v", cID, tReq, req, sID)
				break
			}
		}

	}
}

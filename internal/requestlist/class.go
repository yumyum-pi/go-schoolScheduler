package requestlist

import "github.com/yumyum-pi/go-schoolScheduler/pkg/models"

// Class holds data that contain class id and no. of period required for SubjectRequestE
// unique identifier for the class
// total no. of period required by class of the subject
type Class map[[models.ClassIDBS]byte]int

// TotalReq return the total no. of request made all the classes
func (c *Class) TotalReq() (tReq int) {
	// loop through the map
	for _, r := range *c {
		tReq += r // add the request value to total value
	}
	return
}

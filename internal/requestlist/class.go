package requestlist

import "github.com/yumyum-pi/go-schoolScheduler/pkg/models"

// Class holds data that contain class id and period required by a subject
type Class map[models.ClassIDB]int

// TotalReq return the total no. of request made all the classes
func (c *Class) TotalReq() (tReq int) {
	// loop through the map
	for _, r := range *c {
		tReq += r // add the request value to total value
	}
	return
}

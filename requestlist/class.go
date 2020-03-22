package requestlist

import "github.com/yumyum-pi/go-schoolScheduler/models"

// ClassRLE holds data that contain class id and no. of period required
type ClassRLE struct {
	ID  models.ClassID // unique identifier for the class
	Req int            // total no. of period required by class of the subject
}

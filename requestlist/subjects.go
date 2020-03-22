package requestlist

import (
	"github.com/yumyum-pi/go-schoolScheduler/models"
)

// SubjectRLE is struct of requirement of a subject
// This struct contains the subjectID,
// list of class ID's that require this subject
// Total No. of requirement
type SubjectRLE struct {
	SubjectID models.SubjectID // for identification of the class
	Classes   []ClassRLE       // list of classes that require this subject
	TotalReq  int              // total no. of period needed
}

// Init assign the proper value of the Struct
func (s *SubjectRLE) Init(subject models.Subject, classID models.ClassID) {
	(*s).SubjectID = subject.ID
	(*s).Classes = []ClassRLE{ClassRLE{classID, subject.ReqClasses}}
	(*s).TotalReq = subject.ReqClasses

}

// SubjectRL is an slice of subject request elements
type SubjectRL []SubjectRLE

// Find retutn a index number of the Subject Request List Element
// with the given SubjectID and return -1 if not found
func (srl *SubjectRL) Find(id *models.SubjectID) int {
	// check if Subject Request List is !empty
	if len(*srl) == 0 {
		// List is empty
		// element not found
		return -1
	}
	//loop through the Subject Request List
	for i, s := range *srl {
		// - check if id == subject.ID
		if *id == s.SubjectID {
			return i
		}
	}
	//fmt.Printf("> Error: id=\"%v\", Element not found.\n", id)
	return -1 // element not found
}

// Add the Subeject Request List Element to the Subeject Request List
func (srl *SubjectRL) Add(s *SubjectRLE) {
	// check if the subject exists in the list
	if len(*srl) == 0 {
		// Add the elemen to the list
		(*s).TotalReq = (*s).Classes[0].Req // set total required to no. of classes in the list
		*srl = append(*srl, *s)             // add the whole element to the list
	} else {
		// get the index
		i := srl.Find(&s.SubjectID)
		// if ID matches
		if i != -1 {
			// Add the elemen to the list
			(*srl)[i].Classes = append((*srl)[i].Classes, s.Classes[0]) // add the class to the classes list
			(*srl)[i].TotalReq += (*s).TotalReq                         // add the length of classes to total required
		} else {
			// Add the elemen to the list
			(*s).TotalReq = (*s).Classes[0].Req // set total required to no. of classes in the list
			*srl = append(*srl, *s)             // add the whole element to the list
		}

	}
}

// Create the list form the given list class
func (srl *SubjectRL) Create(classes *[]models.Class) {
	// loop through each classs
	for _, class := range *classes {
		// loop throught each subject
		for _, subject := range class.Subjects {
			// create a subjectRLE
			var newReq SubjectRLE
			newReq.Init(subject, class.ID)

			// add the subjectRLE to the subjectRL
			(*srl).Add(&newReq)
		}
	}
}

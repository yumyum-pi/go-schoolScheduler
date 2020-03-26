package requestlist

import (
	"math/rand"
	"testing"

	"github.com/yumyum-pi/go-schoolScheduler/models"
	"github.com/yumyum-pi/go-schoolScheduler/utils"
)

/*
var ClassID = models.ClassID{
	Standerd: [2]byte{0, 0},
	Section:  [2]byte{0, 0},
	Group:    [2]byte{0, 0},
	Year:     [4]byte{0, 0, 0, 0},
}*/

var classRLE = []ClassRLE{
	ClassRLE{models.ClassID{}, 6},
	ClassRLE{models.ClassID{}, 6},
	ClassRLE{models.ClassID{}, 6},
	ClassRLE{models.ClassID{}, 6},
	ClassRLE{models.ClassID{}, 6},
	ClassRLE{models.ClassID{}, 6},
	ClassRLE{models.ClassID{}, 6},
	ClassRLE{models.ClassID{}, 6},
	ClassRLE{models.ClassID{}, 6},
}

// create a list of list of Subject Request List with just IDs
var list SubjectRL = []SubjectRLE{
	SubjectRLE{models.SubjectID{Standerd: [2]byte{9, 0}, Type: [4]byte{9, 9, 0, 0}}, classRLE[:9], 54},
	SubjectRLE{models.SubjectID{Standerd: [2]byte{1, 8}, Type: [4]byte{1, 1, 3, 0}}, classRLE[:8], 48},
	SubjectRLE{models.SubjectID{Standerd: [2]byte{2, 7}, Type: [4]byte{3, 2, 2, 4}}, classRLE[:7], 42},
	SubjectRLE{models.SubjectID{Standerd: [2]byte{3, 6}, Type: [4]byte{5, 4, 3, 3}}, classRLE[:6], 36},
	SubjectRLE{models.SubjectID{Standerd: [2]byte{4, 5}, Type: [4]byte{4, 6, 5, 4}}, classRLE[:5], 30},
	SubjectRLE{models.SubjectID{Standerd: [2]byte{5, 4}, Type: [4]byte{5, 5, 6, 0}}, classRLE[:4], 24},
	SubjectRLE{models.SubjectID{Standerd: [2]byte{6, 3}, Type: [4]byte{0, 6, 6, 7}}, classRLE[:3], 18},
	SubjectRLE{models.SubjectID{Standerd: [2]byte{7, 2}, Type: [4]byte{8, 0, 7, 7}}, classRLE[:2], 12},
	SubjectRLE{models.SubjectID{Standerd: [2]byte{8, 1}, Type: [4]byte{8, 9, 0, 8}}, classRLE[:1], 6},
}

// Funtion to test if Find works
func TestSubject_FindPass(t *testing.T) {
	// - loop through the list
	for i, e := range list {
		// -- check if the Find methord can find the index of the same element
		index := list.Find(&e.SubjectID)
		if index != i {
			t.Errorf(`> Error: index="%v", i="%v".SubjectID="%v"`, index, i, e.SubjectID)
		}
	}
}

func generateRandomID() *models.SubjectID {
	id := new(models.SubjectID)
	stn, typ := make([]byte, 2), make([]byte, 4)
	rand.Read(stn)
	rand.Read(typ)

	copy(id.Standerd[:], stn[:2])
	copy(id.Type[:], typ[:4])

	return id
}

// Funtion to test if Find works
func TestSubject_FindFail(t *testing.T) {
	// - loop through the list
	for _, e := range list {
		// create a new StudentID
		id := generateRandomID()

		for *id == e.SubjectID {
			id = generateRandomID()
		}

		// -- check if the Find methord can find the index of the same element
		index := list.Find(id)
		if index != -1 {
			t.Errorf(`> Error: index="%v", id="%v".SubjectID="%v"`, index, id, e.SubjectID)
		}
	}
}

var m []models.ClassID = make([]models.ClassID, 5)

var s1 SubjectRL
var s2 SubjectRL = list[:]

func TestSubject_Add_Empty(t *testing.T) {
	n := utils.GenerateRandomInt(9, 10)
	id := s2[n].SubjectID
	s1.Add(&(s2)[n])
	if s1[0].TotalReq != (s2)[n].Classes[0].Req {
		t.Errorf(`> Error: s1[0].TotalReq="%v", (s2)[n].Classes[0].Req="%v"`, s1[0].TotalReq, (s2)[n].Classes[0].Req)
	} else {
		i := s1.Find(&id)

		if i == -1 {
			t.Errorf(`> Error: ID="%v"Element not Found`, id)
		}
	}
}

func TestSubject_Add_Full(t *testing.T) {
	n := utils.GenerateRandomInt(9, 10)
	// saving old data
	classlength, totalRequest := len(s2[n].Classes), s2[n].TotalReq
	s2.Add(&(s2[n]))
	// check if the no. of classes has increased
	if len(s2[n].Classes) != (classlength + 1) {
		t.Errorf(`> Error:n="%v" len(s2[n].Classes)="%v" classlength="%v"`, n, len(s2[n].Classes), classlength)
	}
	if s2[n].TotalReq != 2*totalRequest {
		t.Errorf(`> Error:n="%v", s2[n].TotalReq="%v", totalRequest="%v"`, n, s2[n].TotalReq, totalRequest)
	}
}

func TestSubject_Add_New(t *testing.T) {
	n := utils.GenerateRandomInt(9, 10)
	// Create a new subjectID
	id := models.SubjectID{Standerd: [2]byte{8, 1}, Type: [4]byte{4, 4, 4, 4}}

	// Get old data from list
	e := s2[n]
	e.SubjectID = id // change to subjectID to make it a new data

	// record old data
	length := len(s2)
	// Add new data to existing list
	s2.Add(&e)

	// check if lenth in increased
	if len(s2) != (length + 1) {
		t.Errorf(`> Error: len(s2)="%v", oldLength="%v"`, len(s2), length)
		// check if list returns the proper element id
	} else {
		i := s2.Find(&id)

		if i == -1 {
			t.Errorf(`> Error: ID="%v"Element not Found`, id)
		}
	}
}

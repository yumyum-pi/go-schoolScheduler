package generator

import "github.com/yumyum-pi/go-schoolScheduler/models"

type classID [10]byte

func (c *classID) Init(id *models.ClassID) {
	*c = id.Bytes()
}

type classTT map[classID]models.TimeTable

func (ct *classTT) Init(cls *models.Classes) {
	for _, c := range *cls {
		(*ct)[c.ID.Bytes()] = models.TimeTable{}
	}
}

// redue this struct
type classSRE struct {
	SID    models.SubjectID
	TID    models.TeacherID
	CID    models.ClassID
	SIndex int
}

type classSRs []classSRE

func (csr *classSRs) Init(class *models.Class) {
	for i, s := range class.Subjects {
		(*csr) = append((*csr), classSRE{s.ID, s.TeacherID, class.ID, i})
	}
}

func (csr *classSRs) DeleteElm(i int) {
	l := len(*csr) - 1
	(*csr)[i] = (*csr)[l] // Copy last element to index i.
	(*csr) = (*csr)[:l]
}

package generator

import "github.com/yumyum-pi/go-schoolScheduler/models"

type teacherID [8]byte

func (t *teacherID) Init(id *models.TeacherID) {
	*t = id.Bytes()
}

type teacherTT map[teacherID]models.TimeTable

func (tt *teacherTT) Init(ts *models.Teachers) {
	for _, ts := range *ts {
		(*tt)[ts.ID.Bytes()] = models.TimeTable{}
	}
}

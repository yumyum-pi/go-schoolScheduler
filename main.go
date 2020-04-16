package main

import (
	"fmt"

	td "github.com/yumyum-pi/go-schoolScheduler/internal/testData"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

func printTeacherData(ts *models.Teachers) {
	for _, t := range *ts {
		//fmt.Printf("TeacherID=%v\t\tFree=%v, s=%v\n", t.ID.Bytes(), t.Capacity, t.SubjectCT)
		fmt.Printf("Free=%v\t\tTeacherID=%v\n", t.Capacity, t.ID.Bytes())
	}
}

func printClassData(cs *models.Classes) {
	for _, c := range *cs {
		c.CalCap()
		ss := ""
		for _, s := range c.Subjects {
			ss += fmt.Sprintf("(sID%v r%v)", s.ID.Bytes(), s.Req)
		}

		fmt.Printf("cID%v fP%v %v\n", c.ID.Bytes(), c.NFreePeriod, ss)
	}
}

var c models.Classes
var t models.Teachers

func main() {
	// generate classes and teachers
	td.Create(&c, &t)
	// start creating timetable
	//generator.Init(c, t)

	// get a timetable and print it
}

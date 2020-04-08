package main

import (
	"fmt"

	"github.com/yumyum-pi/go-schoolScheduler/data"
	"github.com/yumyum-pi/go-schoolScheduler/data/generate"
	"github.com/yumyum-pi/go-schoolScheduler/generator"
	"github.com/yumyum-pi/go-schoolScheduler/models"
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
			ss += fmt.Sprintf("(sID%v r%v)", s.ID.Bytes(), s.ReqClasses)
		}

		fmt.Printf("cID%v fP%v %v\n", c.ID.Bytes(), c.NFreePeriod, ss)
	}
}

var c models.Classes
var t models.Teachers

// Init initiate the process
func Init() {
	fmt.Println("> Initialization process: Started")
	c, t = data.Get()

	// fmt.Println(c)
	fmt.Println("> Initialization process: Finished ")
}

func main() {
	// generate classes and teachers
	c, t = generate.Init()

	// assigned teacher to class
	e := c.AssignTeachers(&t)
	// print the error
	if len(e) != 0 {
		fmt.Println("Error in assignTeacherAndClass()", e)
	}
	// start creating timetable
	generator.Init(c, t)

	// get a timetable and print it
}

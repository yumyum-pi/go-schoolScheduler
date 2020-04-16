package testdata

import (
	"fmt"

	rl "github.com/yumyum-pi/go-schoolScheduler/internal/requestlist"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

// printTeachers prints the teacher data one by one on the console
func printTeachers(ts *models.Teachers) {
	for _, t := range *ts {
		var subIDs models.SubjectIDs = t.SubjectCT
		fmt.Printf("ID=%v\tNOS=%v\tCapicity=%v\tSubjects=%v\n", t.ID.Bytes(), len(t.SubjectCT), t.Capacity, subIDs.Types())
	}
	fmt.Printf("\nNo. of Teacher=%v\n", len(*ts))
}

// printClasses prints the class data one by one on the console
func printClasses(cs *models.Classes) {
	for _, c := range *cs {
		c.CalCap()
		fmt.Println("No of free Periods=", c.Capacity)
		for _, sub := range c.Subjects {
			fmt.Printf("cID=%v\tsID=%v\tReq=%v\n", c.ID.Bytes(), sub.ID.Bytes(), sub.Req)
		}
	}
}

// Create class and teacher data
func Create(cs *models.Classes, ts *models.Teachers) {
	generateClasses(cs) // generate class and assign to the cs pointer

	srl := make(rl.Subject) // new Subject Requrest List
	srl.Init(cs)            // create the list using classes pointer

	var trl rl.TeacherRL            // create Teacher Requrest List
	trl.Create(&srl)                // populate teacher request list
	*ts = generateTeacherList(&trl) // generate teachers

	ac := []models.ClassAssigned{} // empty slice of class assigned struct

	// loop through all teacher and reset to default data
	for i := range *ts {
		(*ts)[i].ClassesAssigned = ac
		(*ts)[i].Capacity = models.TeacherCap
	}

	// assigned teacher to class
	e := (*cs).AssignTeachers(ts)
	// print the error
	if len(e) != 0 {
		fmt.Println("Error in assignTeacherAndClass()", e)
	}

	fmt.Println("> Finished: Creating class and teacher data")
	printClasses(cs)
	// printTeachers(ts)
}

package generate

import (
	"fmt"

	"github.com/yumyum-pi/go-schoolScheduler/models"
	"github.com/yumyum-pi/go-schoolScheduler/requestlist"
	"github.com/yumyum-pi/go-schoolScheduler/utils"
)

func printTeacher(ts *models.Teachers) {
	fmt.Printf("Name\t\tID\t\t\tNOS\tCapicity\tSubjects\n")
	for _, t := range *ts {
		var subIDs models.SubjectIDs = t.SubjectCT
		fmt.Printf("%v %v %v\t%v%v\t%v\t%v\t%v\n", t.Name.First, t.Name.Middle, t.Name.Last, t.ID.Year, t.ID.JoinNo, len(t.SubjectCT), t.Capacity, subIDs.Types())
	}
	fmt.Printf("\nNo. of Teacher=%v\n", len(*ts))
}

// Init Data
func Init() {
	classes := generateClasses()
	// create Subject Requrest List
	var srl requestlist.SubjectRL
	srl.Create(&classes)

	var trl requestlist.TeacherRL        // create Teacher Requrest List
	trl.Create(&srl)                     // populate teacher request list
	teacher := generateTeacherList(&trl) // generate teachers

	//printTeacher(&teacher)

	ac := []models.ClassAssigned{} // empty slice of class assigned struct
	// loop through all teacher and reset to default data
	for i := range teacher {
		teacher[i].ClassesAssigned = ac
		teacher[i].Capacity = models.MaxCap
	}
	// fmt.Println(teacher[12].Capacity)
	// write to files
	utils.WriteFile(utils.ResourceFilePath("classes"), classes)
	utils.WriteFile(utils.ResourceFilePath("teachers"), teacher)
}

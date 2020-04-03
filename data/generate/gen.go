package generate

import (
	"fmt"

	"github.com/yumyum-pi/go-schoolScheduler/models"
	"github.com/yumyum-pi/go-schoolScheduler/requestlist"
	"github.com/yumyum-pi/go-schoolScheduler/utils"
)

// printTeachers prints the teacher data one by one on the console
func printTeachers(ts *models.Teachers) {
	for _, t := range *ts {
		var subIDs models.SubjectIDs = t.SubjectCT
		fmt.Printf("Name=%v %v %v\tID=%v\tNOS=%v\tCapicity=%v\tSubjects=%v\n", t.Name.First, t.Name.Middle, t.Name.Last, t.ID.Bytes(), len(t.SubjectCT), t.Capacity, subIDs.Types())
	}
	fmt.Printf("\nNo. of Teacher=%v\n", len(*ts))
}

// printClasses prints the class data one by one on the console
func printClasses(classes *models.Classes) {
	for _, class := range *classes {
		class.CalCap()
		fmt.Println("No of free Periods=", class.NFreePeriod)
		for _, sub := range class.Subjects {
			fmt.Printf("cID=%v\tsID=%v\tReq=%v\n", class.ID.Bytes(), sub.ID.Bytes(), sub.ReqClasses)
		}
	}
}

// Init Data
func Init() {
	var classes models.Classes
	classes = generateClasses()

	// check class capacity and assign extra classes if requried
	// loop through all the classes
	for _, class := range classes {
		class.CalCap()              // calcuate the capacity
		n := class.NFreePeriod      // get the no. of free periods
		subL := len(class.Subjects) // get no. of subject

		// loop over the no. of free periods
		for i := 0; i < n; i++ {
			// loop the index no from 0 to subL
			looper := (i / subL)
			subIndex := i - (looper * subL)

			// increase the no. of required classes
			class.Subjects[subIndex].ReqClasses++
		}
	}

	// create Subject Requrest List
	var srl requestlist.SubjectRL
	srl.Create(&classes)

	var trl requestlist.TeacherRL        // create Teacher Requrest List
	trl.Create(&srl)                     // populate teacher request list
	teacher := generateTeacherList(&trl) // generate teachers

	ac := []models.ClassAssigned{} // empty slice of class assigned struct
	// loop through all teacher and reset to default data
	for i := range teacher {
		teacher[i].ClassesAssigned = ac
		teacher[i].Capacity = models.MaxCap
	}

	// write to files
	utils.WriteFile(utils.ResourceFilePath("classes"), classes)
	utils.WriteFile(utils.ResourceFilePath("teachers"), teacher)
}

package generator

import (
	"fmt"

	"github.com/yumyum-pi/go-schoolScheduler/models"
	"github.com/yumyum-pi/go-schoolScheduler/utils"
)

var c models.Classes
var t models.Teachers

var ctt classTT = make(classTT)
var ttt teacherTT = make(teacherTT)

func assignPeriod(cIndex, sIndex, pID int) (e error) {
	tID := c[cIndex].Subjects[sIndex].TeacherID // get teacher ID
	cID := c[cIndex].ID                         // get classID
	sID := c[cIndex].Subjects[sIndex].ID        // get subjectID

	tt := ttt[tID.Bytes()] // get the timetable with the given ID
	ct := ctt[cID.Bytes()] // get the timetable with the given ID

	// create empty varible to check
	emSID := models.SubjectID{}
	emCID := models.ClassID{}

	checkSID := tt[pID].SubjectID == emSID
	checkCID := tt[pID].ClassID == emCID

	// check if assigned
	if checkSID && checkCID {
		// assign subjectID and teacherID to period in class timeTable
		tt[pID].SubjectID = sID
		tt[pID].ClassID = cID

		// assign subjectID and teacherID to period in class timeTable
		ct[pID].SubjectID = sID
		ct[pID].TeacherID = tID

		ttt[tID.Bytes()] = tt // reassign the timetable
		ctt[cID.Bytes()] = ct // reassign the timetable
		return
	}

	return fmt.Errorf("> Error: tt[pID].SubjectID=%v checkSID=%v checkCID=%v", tt[pID].SubjectID, checkSID, checkSID)

}

func class(cIndex int, pID int) {

	// get current class crps
	var crl classSRs
	crl.Init(&c[cIndex]) // init the values
	assigned := false
	for len(crl) > 0 {
		// create a random number between 0 and crl length
		randInt := utils.GenerateRandomInt(len(crl), 10)

		// get random sub from classs
		sIndex := crl[randInt].SIndex

		//sIndex := getRandomSub(&c[cIndex])
		// check if the teacher has another assigned period
		e := assignPeriod(cIndex, sIndex, pID)
		if e == nil {
			assigned = true
			break
		}
		crl.DeleteElm(randInt)
	}

	if !assigned {
		fmt.Printf("Class not assigned: classID=%v\tperiodID=%v\n", c[cIndex].ID.Bytes(), pID)
	}
}

// Start will start the generating processing
func Start(classes []models.Class, teachers models.Teachers) {
	c = classes
	t = teachers

	ctt.Init(&c)
	ttt.Init(&t)

	for pID := 0; pID < models.MaxCap; pID++ {
		cLen := len(c)
		for cIndex := cLen - 1; cIndex >= 0; cIndex-- {
			class(cIndex, pID)
		}
	}
}

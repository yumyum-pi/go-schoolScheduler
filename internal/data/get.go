package data

import (
	"github.com/yumyum-pi/go-schoolScheduler/internal/utils"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

//var classes []models.Class

/*
Getreturn the following data:
 - Array of classes
 - Array of teacher
This func is reading from local file
*/
func Get() (classes []models.Class, teachers []models.Teacher) {
	// TODO get data from database

	// get class data
	getClass(&classes)
	getTeacher(&teachers)

	// return the data
	return
}

func getClass(c *[]models.Class) {
	utils.ReadFile(utils.ResourceFilePath("classes"), &c)
}

func getTeacher(t *[]models.Teacher) {
	utils.ReadFile(utils.ResourceFilePath("teachers"), &t)
}

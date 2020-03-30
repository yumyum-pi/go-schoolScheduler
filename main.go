package main

import (
	"fmt"

	"github.com/yumyum-pi/go-schoolScheduler/data"
	"github.com/yumyum-pi/go-schoolScheduler/generator"

	// "github.com/yumyum-pi/go-schoolScheduler/data/generate"
	"github.com/yumyum-pi/go-schoolScheduler/models"
)

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
	Init()
	//generate.Init()
	// Generate requrest list
	// assign teacher to the classes
	e := c.AssignTeachers(&t)
	if len(e) != 0 {
		fmt.Println("Error in assignTeacherAndClass()", e)
	}
	generator.Start(c, t)
}

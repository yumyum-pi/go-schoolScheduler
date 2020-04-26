package generate

import (
	"fmt"

	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

// Create class and teacher data
func Create(cs *[]models.Class, ts *models.Teachers) {
	generateClassL(cs) // generate class and assign to the cs pointer

	*ts = gTeacherLM2(cs) // generate teachers

	fmt.Println("> Finished: Creating class and teacher data")
	models.PrintClassL(*cs)
	ts.Print()
}

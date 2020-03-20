package generate

import "github.com/yumyum-pi/go-schoolScheduler/utils"

// Init Data
func Init() {
	classes := generateClasses()
	// create requrestList
	// generate teachers
	// write to files
	utils.WriteFile(utils.ResourceFilePath("classes"), classes)
}

package file

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/utils"
	"google.golang.org/protobuf/proto"
)

// Read file from thr disk
func Read(fileName string) *models.TimeTable {
	fmt.Println(fileName)
	in, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	tt := &models.TimeTable{}
	if err := proto.Unmarshal(in, tt); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	return tt
}

// ReadRand reads a random file from the directory
func ReadRand(dir string) *models.TimeTable {
	files, err := ioutil.ReadDir(dir)
	var filePath []string
	if err == nil {
		// loop through all the file
		for _, file := range files {
			if !file.IsDir() && filepath.Ext(file.Name()) == ".tt" {
				p := filepath.Join(dir, file.Name())
				filePath = append(filePath, p)
			}
		}
	}
	fl := len(filePath)
	// get random file
	if fl == 0 {
		return nil
	}
	i := utils.GenerateRandomInt(fl, 10)
	return Read(filePath[i])
}

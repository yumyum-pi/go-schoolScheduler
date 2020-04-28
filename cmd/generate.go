package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/generator"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/utils"
	"google.golang.org/protobuf/proto"
)

var input, output string

var genCMD = &cobra.Command{
	Use:   "gen",
	Short: "Generating timetable",
	Long:  "Generates timetable",
	Run: func(cmd *cobra.Command, args []string) {
		l := len(args)

		// check argument not empty
		if l == 0 {
			fmt.Printf("> Error: Please provide one path.\n")
			os.Exit(1)
		}
		// check if two path
		if l > 1 {
			fmt.Printf("> Error: Please provide only one path.\n")
			os.Exit(1)
		}
		// check if file or path
		info, e := os.Stat(args[0])
		if e != nil {
			fmt.Println(e)
			os.Exit(1)
		}

		var tt *models.TimeTable
		// if directory select random file
		if info.IsDir() {
			tt = ReadRand(args[0])
		} else {
			tt = Read(args[0])
		}
		// start the generating process
		generator.Start(tt)
	},
}

func init() {
	rootCmd.AddCommand(genCMD)
}

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

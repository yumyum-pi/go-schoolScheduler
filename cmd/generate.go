package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/file"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/generator"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
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
			tt = file.ReadRand(args[0])
		} else {
			tt = file.Read(args[0])
		}

		s0, geneSize, err := (*tt).Decode()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// start the generating process
		generator.Start(s0, geneSize)
	},
}

func init() {
	rootCmd.AddCommand(genCMD)
}

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var checkV bool
var version string
var logDir string // directory for log files

var rootCmd = &cobra.Command{
	Use:   "go-tt",
	Short: "Generates timetables",
	Long:  `go-tt:Go-Timetable, is a application that creates timetable for schools`,
	Run: func(cmd *cobra.Command, args []string) {
		if checkV {
			fmt.Printf("go-tt:%v\n", version)
		}
	},
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(readVersion)

	// for printing version
	rootCmd.PersistentFlags().BoolVarP(
		&checkV,   // variable
		"version", // name
		"v",       // shothand
		false,     // default
		"version of the application",
	)

	// for specifing log directory
	rootCmd.PersistentFlags().StringVarP(
		&logDir,  // variable
		"log",    // name
		"d",      // shothand
		"./logs", // default
		"directory for log files",
	)
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func readVersion() {
	v, err := ioutil.ReadFile("./version")
	if err != nil {
		panic(err)
	} else {
		version = string(v)
	}
}

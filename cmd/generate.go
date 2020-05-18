package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/file"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/generator"
	l "github.com/yumyum-pi/go-schoolScheduler/pkg/log"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
)

var genCMD = &cobra.Command{
	Use:   "gen",
	Short: "Generating timetable",
	Long:  "Generates timetable",
	Run: func(cmd *cobra.Command, args []string) {
		argL := len(args)

		// check argument not empty
		if argL == 0 {
			fmt.Printf("> Error: Please provide one path.\n")
			os.Exit(1)
		}
		// check if two path
		if argL > 1 {
			fmt.Printf("> Error: Please provide only one path.\n")
			os.Exit(1)
		}
		// check if file or path
		info, e := os.Stat(args[0])
		if e != nil {
			fmt.Println(e)
			os.Exit(1)
		}

		req := &models.GRequest{}
		// if directory select random file
		if info.IsDir() {
			req = file.ReadRand(args[0])
		} else {
			req = file.Read(args[0])
		}
		s0, geneSize, e := models.Decode(&req.Pkgs, req.GetGSize())
		if e != nil {
			l.Fatal(req.ClientID, req.ServerID, len(*s0), geneSize, 48, e.Error())
		}
		// start the generating process
		s1, nErr, e := generator.Start(s0, geneSize, int(req.NNType))
		res := models.GResponse{}
		res.NError = int32(nErr)
		res.Pkgs = *models.Encode(s1)
		if e != nil {
			l.Error(req.ClientID, req.ServerID, len(*s0), geneSize, 48, e.Error())
		}
		l.Info(req.ClientID, req.ServerID, len(*s0), geneSize, 48, "")
	},
}

func init() {
	rootCmd.AddCommand(genCMD)
}

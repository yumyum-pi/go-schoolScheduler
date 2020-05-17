package cmd

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/generator"
	l "github.com/yumyum-pi/go-schoolScheduler/pkg/log"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
	"google.golang.org/grpc"
)

// server variables
var port int32
var ip string

// server is used to implement customer.CustomerServer.
type server struct {
	savedCustomers *models.SequencePkgs
}

// Genereate is the gRPC function to resend response data back to the client
func (s *server) GenerateTT(ctx context.Context, seq *models.SequencePkgs) (*models.SequencePkgs, error) {
	s0, geneSize, e := (*seq).Decode()
	if e != nil {
		return nil, e
	}
	// start the generating process
	s1, e := generator.Start(s0, geneSize)
	seq.Encode(s1)
	if e != nil {
		l.ErrorLogger("0.0.0.0", "23rewG", 48, 51, e.Error())
		return seq, nil
	}
	return seq, nil
}

var serverCMD = &cobra.Command{
	Use:   "server",
	Short: "Start gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		l.Init(logDir)
		// address for the server
		a := fmt.Sprintf("%v:%v", ip, port)
		lis, err := net.Listen("tcp", a)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		// Creates a new gRPC server
		s := grpc.NewServer()
		models.RegisterSequenceServer(s, &server{})

		s.Serve(lis)
	},
}

func init() {
	// get port from the cli command
	serverCMD.PersistentFlags().Int32VarP(
		&port,  // variable
		"port", // name
		"p",    // shothand
		5501,   // default
		"port no for the server",
	)
	// get ip address from the cli command
	serverCMD.PersistentFlags().StringVarP(
		&ip,       // variable
		"ip",      // name
		"i",       // shothand
		"0.0.0.0", // default
		"Ip of the server",
	)
	rootCmd.AddCommand(serverCMD)
}

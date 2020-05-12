package cmd

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/generator"
	"github.com/yumyum-pi/go-schoolScheduler/pkg/models"
	"google.golang.org/grpc"
)

// server is used to implement customer.CustomerServer.
type server struct {
	savedCustomers *models.SequencePkgs
}

func (s *server) GenerateTT(ctx context.Context, seq *models.SequencePkgs) (*models.SequencePkgs, error) {
	s0, geneSize, e := (*seq).Decode()
	if e != nil {
		return nil, e
	}
	// start the generating process
	s1, e := generator.Start(s0, geneSize)
	seq.Encode(s1)
	if e != nil {
		return seq, e
	}
	return seq, nil
}

var port int32
var ip string
var serverCMD = &cobra.Command{
	Use:   "server",
	Short: "Start gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
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

	serverCMD.PersistentFlags().Int32VarP(
		&port,  // variable
		"port", // name
		"p",    // shothand
		5501,   // default
		"port no for the server",
	)

	serverCMD.PersistentFlags().StringVarP(
		&ip,       // variable
		"ip",      // name
		"i",       // shothand
		"0.0.0.0", // default
		"Ip of the server",
	)
	rootCmd.AddCommand(serverCMD)
}

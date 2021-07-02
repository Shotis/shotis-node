package cmd

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/rbrick/shotis-node/network"
	"github.com/rbrick/shotis-node/storage"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// rpcCmd represents the rpc command
var rpcCmd = &cobra.Command{
	Use:   "rpc",
	Short: "Runs an gRPC server on the Shotis network",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rpc called")

		//TODO: Implement a more robust server with go routines & worker pools and actually read from a config file
		//       etc. etc.
		server := grpc.NewServer()

		service, _ := storage.NewGoogleCloudStorage(context.Background(), "key", storage.Bucket("bucket"))

		network.RegisterShotisServiceServer(server, &network.ServerImpl{
			GCPService: service,
		})

		l, err := net.Listen("tcp", ":1337")

		if err != nil {
			log.Fatalln(err)
		}

		server.Serve(l)
	},
}

func init() {
	rootCmd.AddCommand(rpcCmd)
}

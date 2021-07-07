package cmd

import (
	"context"
	"crypto/tls"
	"log"
	"net"

	"github.com/shotis/shotis-node/network"
	"github.com/shotis/shotis-node/storage"
	"github.com/shotis/shotis-node/tasks"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type GRPCWorker struct {
	Backing       *network.ServerImpl
	AwaitingTasks *tasks.Queue
}

// rpcCmd represents the rpc command
var rpcCmd = &cobra.Command{
	Use:   "rpc",
	Short: "Runs an gRPC server on the Shotis network",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("starting gRPC server on", Config.Server.RPC.Host)

		cert, err := tls.LoadX509KeyPair(Config.Server.TLS.CertPath, Config.Server.TLS.KeyPath)

		if err != nil {
			log.Fatalln(err)
		}

		server := grpc.NewServer(grpc.Creds(credentials.NewServerTLSFromCert(&cert)))
		service, err := storage.NewGoogleCloudStorage(context.Background(), Config.Storage.AuthKey, storage.Bucket(Config.Storage.Bucket))

		if err != nil {
			log.Println(err)
		} else {
			log.Println("created Google Cloud Storage server")
		}

		worker := &GRPCWorker{
			Backing: &network.ServerImpl{
				GCPService: service,
			},
			AwaitingTasks: tasks.NewQueue(100),
		}

		network.RegisterShotisServiceServer(server, worker.Backing)

		l, err := net.Listen("tcp", Config.Server.RPC.Host)

		if err != nil {
			log.Fatalln(err)
		}

		server.Serve(l)
	},
}

func init() {
	rootCmd.AddCommand(rpcCmd)
}

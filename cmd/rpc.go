package cmd

import (
	"context"
	"log"
	"sync"

	"github.com/shotis/shotis-node/worker"
	"github.com/spf13/cobra"
)

// rpcCmd represents the rpc command
var rpcCmd = &cobra.Command{
	Use:   "rpc",
	Short: "Runs an gRPC server on the Shotis network",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("starting gRPC worker on", Config.Server.RPC.Host)
		worker, err := worker.NewGRPCWorker(context.Background(), Config)

		if err != nil {
			log.Fatalln(err)
		}

		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			if Config.Server.TLS.Enabled {
				worker.StartTLS(Config.Server.RPC.Host, Config.Server.TLS.CertPath, Config.Server.TLS.KeyPath)
			} else {
				worker.Start(Config.Server.Host)
			}
		}()

		log.Println("successfully started gRPC worker.")
		wg.Wait()

	},
}

func init() {
	rootCmd.AddCommand(rpcCmd)
}

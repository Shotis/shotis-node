package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/shotis/shotis-node/network"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// healthreportCmd represents the healthreport command
var healthreportCmd = &cobra.Command{
	Use:   "healthreport",
	Short: "Get a health report of a specific on the network",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("healthreport called")

		conn, err := grpc.Dial(":1337", grpc.WithInsecure())

		if err != nil {
			log.Fatalln(err)
		}

		shotisClient := network.NewShotisServiceClient(conn)

		report, err := shotisClient.Health(context.Background(), &network.HealthReportRequest{})

		if err != nil {
			log.Fatalln(err)
		}

		jsreport, _ := json.MarshalIndent(report, "", " ")

		fmt.Println(string(jsreport))

		conn.Close()
	},
}

func init() {
	rootCmd.AddCommand(healthreportCmd)
}

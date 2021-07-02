package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/rbrick/shotis-node/network"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Runs an HTTP API server for Shotis",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("http called")

		f, _ := os.Open(args[0])

		conn, err := grpc.Dial(":1337", grpc.WithInsecure())

		if err != nil {
			log.Fatalln(err)
		}

		shotisClient := network.NewShotisServiceClient(conn)

		b, _ := ioutil.ReadAll(f)

		report, err := shotisClient.UploadImage(context.Background(), &network.UploadImageMessage{
			FileName: args[0],
			MimeType: "whatever",
			Data:     b,
		})

		if err != nil {
			log.Fatalln(err)
		}

		jsreport, _ := json.MarshalIndent(report, "", " ")

		fmt.Println(string(jsreport))

		conn.Close()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}

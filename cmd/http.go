package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/shotis/shotis-node/network"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Runs an HTTP API server for Shotis",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		f, _ := os.Open(args[0])

		conn, err := grpc.Dial(":1337", grpc.WithInsecure())

		if err != nil {
			log.Fatalln(err)
		}

		shotisClient := network.NewShotisServiceClient(conn)

		stream, err := shotisClient.UploadImage(cmd.Context())

		if err != nil {
			log.Fatalln(err)
		}

		br := bufio.NewReader(f)

		buffer := make([]byte, 512)

		header := &network.FileHeader{
			FileName: f.Name(),
			FileType: "whatever",
		}

		totalRead := 0

		for read, err := br.Read(buffer); read != -1 && err != io.EOF; {
			stream.Send(&network.UploadImageMessage{
				Header: header,
				Data:   buffer[:read],
			})

			totalRead += read
			read, err = br.Read(buffer)
		}

		response, err := stream.CloseAndRecv()

		if err != nil {
			log.Fatalln(err)
		}

		j, _ := json.MarshalIndent(response, "", " ")

		fmt.Println(string(j))
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}

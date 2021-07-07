package cmd

import (
	"context"
	"log"

	"github.com/shotis/shotis-node/web"
	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Runs an HTTP API server for Shotis",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ws, err := web.Init(Config)

		if err != nil {
			log.Fatalln(err)
		}

		ws.Start(context.Background())
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}

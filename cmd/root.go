package cmd

import (
	"log"

	"github.com/shotis/shotis-node/config"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

var Config *config.NodeConfig

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "shotis-node",
	Short: "Command a node on the shotis network",
	Long: `shotis nodes are broken down into two possible configurations. 
	
	One is a gRPC server node, which starts running a gRPC server. These server nodes handle direct communications
	to Google Cloud Storage and are used as essentially workers for handling uploading and fetching of images.
	
	The second is an API server node, which starts an HTTP server and handles all the more forward facing API endpoints.
	This is in direct communication with the gRPC server node(s)

	The commands are 

	shotis-node http - Runs the HTTP API server node.
	shotis-node rpc  - Runs the gRPC server node. 
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.json", "Path to the configuration file")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	}

	conf, err := config.ReadConfig(cfgFile)

	if err != nil {
		log.Fatalln(err)
	}

	Config = conf
}

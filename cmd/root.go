package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Bitlatte/beam/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "beam",
	Short: "Beam commands across the cosmos",
	Long:  `Execute shell scripts across a multitude of servers with a single command`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := utils.GetSSHConfig()
		if err != nil {
			log.Fatal(err)
		}

		// Loop through hosts
		for host := range config.Hosts {
			vs, err := utils.CreateSession(config, &config.Hosts[host])
			if err != nil {
				log.Fatal(err)
			}

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			timeout, _ := time.ParseDuration("6s")

			respChan := vs.Run(ctx, "ping -c 4 192.168.55.10", timeout)

			resp := <-respChan
			if err := resp.Err(); err != nil {
				fmt.Printf("[%s] %s", config.Hosts[host].Address, err)
			}

			stream := resp.GetStream()
			defer stream.Close()

			for stream.ScanStdout() {
				txt := stream.TextStdout()
				fmt.Printf("[%s] %s", config.Hosts[host].Address, txt)
			}
		}

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	cwd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", (cwd + "/config.json"), "config file to use")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		cwd, err := os.Getwd()
		cobra.CheckErr(err)

		viper.AddConfigPath(cwd)
		viper.SetConfigType("json")
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using Config File:", viper.ConfigFileUsed())
	}
}

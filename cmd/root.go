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
	"github.com/yahoo/vssh"
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

			vs := vssh.New().Start().OnDemand()

			clientConfig, err := vssh.GetConfigPEM(config.Auth.User, config.Auth.Key)
			if err != nil {
				log.Fatal(err)
			}

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			timeout, _ := time.ParseDuration("6s")
			client := fmt.Sprintf("%s:%s", config.Hosts[host].Address, fmt.Sprint(config.Hosts[host].Port))
			vs.AddClient(client, clientConfig, vssh.SetMaxSessions(2))
			vs.Wait()

			respChan := vs.Run(ctx, "ping", timeout)

			resp := <-respChan
			if err := resp.Err(); err != nil {
				log.Fatal(err)
			}

			stream := resp.GetStream()
			defer stream.Close()

			for stream.ScanStdout() {
				txt := stream.TextStdout()
				fmt.Println(txt)
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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", (cwd + "/beamconf.json"), "config file to use")
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

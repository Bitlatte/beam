package cmd

import (
	"context"
	"errors"
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
	file    string
)

var rootCmd = &cobra.Command{
	Use:   "beam [flags] <file>",
	Short: "Beam commands across the cosmos",
	Long:  `Execute shell scripts across a multitude of servers with a single command`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("missing file path")
		}
		_, err := utils.ReadFile(args[0])
		if err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		file = args[0]
		vs := vssh.New().Start().OnDemand()
		// Get Beam Config from config.json
		config, err := utils.GetBeamConfig()
		if err != nil {
			panic(err)
		}
		// Loop through All The Hosts within the Config File
		for host := range config.Hosts {
			// Address:Port
			address := fmt.Sprintf("%s:%s", config.Hosts[host].Address, fmt.Sprint(config.Hosts[host].Port))
			//
			auth := utils.GetAuth(config, &config.Hosts[host])

			hostConfig, err := utils.GetClientConfig(auth)
			if err != nil {
				panic(err)
			}

			vs.AddClient(address, hostConfig, vssh.SetMaxSessions(4))
			vs.Wait()

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			timeout, _ := time.ParseDuration("6s")

			// Copy File To Remote Host
			err = utils.RemoteCopy(address, auth["user"], file, hostConfig)
			if err != nil {
				panic(err)
			}

			cmds := []string{
				"./script.sh",
				"rm -rf ./script.sh",
			}

			for cmd := range cmds {
				respChan := vs.Run(ctx, cmds[cmd], timeout)

				for resp := range respChan {
					if err := resp.Err(); err != nil {
						panic(err)
					}
					stream := resp.GetStream()
					defer stream.Close()

					for stream.ScanStdout() {
						txt := stream.TextStdout()
						fmt.Println(resp.ID(), txt)
					}

					if err := stream.Err(); err != nil {
						panic(err)
					}
				}
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

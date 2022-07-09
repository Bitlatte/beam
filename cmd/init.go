/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	user string
	key  string
)

func writeDefaultConfig(keypath string, user string) error {
	config := fmt.Sprintf(`{
	"auth": {
		"keyfile": "%s",
		"user": "%s"
	},
	"hosts": {}
}`, keypath, user)
	err := os.WriteFile("beamconf.json", []byte(config), 0755)

	if err != nil {
		return err
	}

	return nil
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Configure Beam",
	Long: `
Setup config for Beam. If run without flags
a general config file will be generated for you.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := writeDefaultConfig(key, user)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	home, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	initCmd.Flags().StringVarP(&user, "user", "u", "root", "User to login with")
	initCmd.Flags().StringVarP(&key, "key", "k", (home + "/.ssh/id_rsa"), "Key file to use")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

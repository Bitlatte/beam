/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Configure Beam",
	Long: `Setup the default config for Beam.
Until this is done Beam will default to the following:

	SSH Key for Auth
	Port 22 for SSH Connections`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().IntP("port", "p", 22, "Set the default SSH Connection Port")
	initCmd.Flags().Bool("keys", true, "Use SSH Keys over Password Auth")
	initCmd.Flags().String("keyfile", (os.Getenv("HOME") + "/.ssh/id_rsa"), "SSH Key path to use for Auth")
	initCmd.Flags().String("password", "", "Password to use if using Password Auth")
	initCmd.Flags().String("user", "root", "Default User to use when Logging in to Hosts")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

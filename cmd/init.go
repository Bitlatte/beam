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
		"key": "%s",
		"user": "%s"
	},
	"hosts": []
}`, keypath, user)
	err := os.WriteFile("config.json", []byte(config), 0755)

	if err != nil {
		return err
	}

	return nil
}

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
}

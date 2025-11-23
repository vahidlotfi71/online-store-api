package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vahidlotfi71/online-store-api/Utils/Constants"
)

var showVersion bool

var rootCmd = &cobra.Command{
	Use:   "shop",
	Short: "Online Store API CLI",
	Long:  `CLI tools for managing the Online Store API (serve, migrate, seed, routes, etc.).`,
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			cmd.Printf("Application version %s\n", Constants.VERSION)
			return
		}
		cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "show application version")
}

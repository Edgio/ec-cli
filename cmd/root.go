package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	cobra.OnInitialize()
	rootCmd := &cobra.Command{
		Use:   "EC CLI",
		Short: "EC CLI",
		Long: `A CLI library for Go that empowers users to onboard edgecast
		products and services.`,
	}
	err := rootCmd.Execute()

	handleError(err)
}

func handleError(err error) {
	if err != nil {
		os.Exit(1)
	}
}

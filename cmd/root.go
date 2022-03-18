package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "EC CLI",
		Short: "EC CLI",
		Long: `A CLI library for Go that empowers users to onboard edgecast
		products and services.`,
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

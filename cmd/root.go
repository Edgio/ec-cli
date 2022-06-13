package cmd

import (
	"log"

	"github.com/edgecast/ec-cli/cmd/waf"
	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "EC CLI",
		Short: "EC CLI",
		Long: `A CLI library for Go that empowers users to onboard edgecast
		products and services.`,
	}

	rootCmd.AddCommand(waf.Root())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

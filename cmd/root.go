// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "EC CLI",
		Short: "EC CLI",
		Long: `A CLI library for Go that empowers users to onboard edgecast
		products and services.`,
	}

	addSdkCommands(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

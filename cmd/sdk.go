// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.
package cmd

// This file is auto-generated. Modifications to this file may be overwritten.

import (
	"github.com/edgecast/ec-cli/cmd/waf"
	"github.com/spf13/cobra"
)

func addSdkCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(waf.Root())
}

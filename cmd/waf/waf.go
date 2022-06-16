// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.
package waf

// This file is auto-generated. Modifications to this file may be overwritten.

import (
	"log"

	"github.com/edgecast/ec-cli/cmd/waf/rules/access"
	"github.com/edgecast/ec-cli/cmd/waf/rules/bot"
	"github.com/edgecast/ec-cli/cmd/waf/rules/custom"
	"github.com/edgecast/ec-cli/cmd/waf/rules/managed"
	"github.com/edgecast/ec-cli/cmd/waf/rules/rate"
	"github.com/edgecast/ec-cli/cmd/waf/scopes"
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	root := &cobra.Command{
		Use:   "waf",
		Short: "waf",
		Long:  "waf",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	root.AddCommand(custom.Root())
	root.AddCommand(managed.Root())
	root.AddCommand(rate.Root())
	root.AddCommand(access.Root())
	root.AddCommand(bot.Root())
	root.AddCommand(scopes.Root())

	return root
}

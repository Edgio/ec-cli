// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0 
// license. See LICENSE file in project root for terms.
package waf

import (
    "log"

    "github.com/spf13/cobra"
    
    "github.com/edgecast/ec-cli/cmd/waf/rules/access"
    "github.com/edgecast/ec-cli/cmd/waf/rules/bot"
    "github.com/edgecast/ec-cli/cmd/waf/rules/custom"
    "github.com/edgecast/ec-cli/cmd/waf/rules/managed"
    "github.com/edgecast/ec-cli/cmd/waf/rules/rate"
    "github.com/edgecast/ec-cli/cmd/waf/scopes"
)

func Root() *cobra.Command {
    root := &cobra.Command{
        Use:   "waf",
        Short: "waf",
        Long: "waf",
        Run: func(cmd *cobra.Command, args []string) {
            err := cmd.Help()
            if err != nil {
                log.Fatal(err)
            }
        },
    }
    
	root.AddCommand(access.Root())
	root.AddCommand(bot.Root())
	root.AddCommand(custom.Root())
	root.AddCommand(managed.Root())
	root.AddCommand(rate.Root())
	root.AddCommand(scopes.Root())

	return root
}

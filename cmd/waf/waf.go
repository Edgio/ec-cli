package waf

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/edgecast/ec-cli/cmd/waf/access_rules"
	"github.com/edgecast/ec-cli/cmd/waf/bot_rule_sets"
	"github.com/edgecast/ec-cli/cmd/waf/custom_rule_sets"
	"github.com/edgecast/ec-cli/cmd/waf/managed_rules"
	"github.com/edgecast/ec-cli/cmd/waf/rate_rules"
	"github.com/edgecast/ec-cli/cmd/waf/scopes"
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

	root.AddCommand(access_rules.Root())
	root.AddCommand(bot_rule_sets.Root())
	root.AddCommand(custom_rule_sets.Root())
	root.AddCommand(managed_rules.Root())
	root.AddCommand(rate_rules.Root())
	root.AddCommand(scopes.Root())

	return root
}

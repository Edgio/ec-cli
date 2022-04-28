package rtld

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/edgecast/ec-cli/cmd/rtld/lookups"
	"github.com/edgecast/ec-cli/cmd/rtld/profiles_cdn"
	"github.com/edgecast/ec-cli/cmd/rtld/profiles_rl"
	"github.com/edgecast/ec-cli/cmd/rtld/profiles_waf"
	"github.com/edgecast/ec-cli/cmd/rtld/settings_internal"
)

func Root() *cobra.Command {
	root := &cobra.Command{
		Use:   "rtld",
		Short: "rtld",
		Long:  "rtld",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	root.AddCommand(lookups.Root())
	root.AddCommand(profiles_cdn.Root())
	root.AddCommand(profiles_rl.Root())
	root.AddCommand(profiles_waf.Root())
	root.AddCommand(settings_internal.Root())

	return root
}

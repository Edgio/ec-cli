package settings_internal

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/EdgeCast/ec-sdk-go/edgecast/rtld/settings_internal"
	"github.com/spf13/cobra"
)

func parseInputParams[T any](c *cobra.Command, target *T) error {
	data := c.Flag("input-params-json").Value.String()

	if err := json.Unmarshal([]byte(data), target); err != nil {
		return fmt.Errorf("error parsing input params: %w", err)
	}

	return nil
}

func displayResponse(obj interface{}) {
	bytes, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", string(bytes))
}

// Root returns the settings_internal root command
func Root() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "settings_internal",
		Short: "settings_internal",
		Long:  "settings_internal",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	cmdSettingsGetRlSettings := &cobra.Command{
		Use:   "SettingsGetRlSettings",
		Short: "SettingsGetRlSettings",
		Long:  "SettingsGetRlSettings",
		Run: func(cmd *cobra.Command, args []string) {
			var input settings_internal.SettingsGetRlSettingsParams
			err := parseInputParams(cmd, &input)
			if err != nil {
				log.Fatal(err)
			}
			displayResponse(input)
		},
	}
	cmdSettingsGetRlSettings.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
	rootCmd.AddCommand(cmdSettingsGetRlSettings)

	cmdSettingsGetSettingsByPlatform := &cobra.Command{
		Use:   "SettingsGetSettingsByPlatform",
		Short: "SettingsGetSettingsByPlatform",
		Long:  "SettingsGetSettingsByPlatform",
		Run: func(cmd *cobra.Command, args []string) {
			var input settings_internal.SettingsGetSettingsByPlatformParams
			err := parseInputParams(cmd, &input)
			if err != nil {
				log.Fatal(err)
			}
			displayResponse(input)
		},
	}
	cmdSettingsGetSettingsByPlatform.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
	rootCmd.AddCommand(cmdSettingsGetSettingsByPlatform)

	cmdSettingsGetWafSettings := &cobra.Command{
		Use:   "SettingsGetWafSettings",
		Short: "SettingsGetWafSettings",
		Long:  "SettingsGetWafSettings",
		Run: func(cmd *cobra.Command, args []string) {
			var input settings_internal.SettingsGetWafSettingsParams
			err := parseInputParams(cmd, &input)
			if err != nil {
				log.Fatal(err)
			}
			displayResponse(input)
		},
	}
	cmdSettingsGetWafSettings.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
	rootCmd.AddCommand(cmdSettingsGetWafSettings)

	return rootCmd
}

package profiles_cdn

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/EdgeCast/ec-sdk-go/edgecast/rtld/profiles_cdn"
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

// Root returns the profiles_cdn root command
func Root() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "profiles_cdn",
		Short: "profiles_cdn",
		Long:  "profiles_cdn",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	cmdProfilesAddCustomerSetting := &cobra.Command{
		Use:   "ProfilesAddCustomerSetting",
		Short: "ProfilesAddCustomerSetting",
		Long:  "ProfilesAddCustomerSetting",
		Run: func(cmd *cobra.Command, args []string) {
			var input profiles_cdn.ProfilesAddCustomerSettingParams
			err := parseInputParams(cmd, &input)
			if err != nil {
				log.Fatal(err)
			}
			displayResponse(input)
		},
	}
	cmdProfilesAddCustomerSetting.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
	rootCmd.AddCommand(cmdProfilesAddCustomerSetting)

	cmdProfilesDeleteCustomerSettingsByID := &cobra.Command{
		Use:   "ProfilesDeleteCustomerSettingsByID",
		Short: "ProfilesDeleteCustomerSettingsByID",
		Long:  "ProfilesDeleteCustomerSettingsByID",
		Run: func(cmd *cobra.Command, args []string) {
			var input profiles_cdn.ProfilesDeleteCustomerSettingsByIDParams
			err := parseInputParams(cmd, &input)
			if err != nil {
				log.Fatal(err)
			}
			displayResponse(input)
		},
	}
	cmdProfilesDeleteCustomerSettingsByID.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
	rootCmd.AddCommand(cmdProfilesDeleteCustomerSettingsByID)

	cmdProfilesGetCustomerSettings := &cobra.Command{
		Use:   "ProfilesGetCustomerSettings",
		Short: "ProfilesGetCustomerSettings",
		Long:  "ProfilesGetCustomerSettings",
		Run: func(cmd *cobra.Command, args []string) {
			var input profiles_cdn.ProfilesGetCustomerSettingsParams
			err := parseInputParams(cmd, &input)
			if err != nil {
				log.Fatal(err)
			}
			displayResponse(input)
		},
	}
	cmdProfilesGetCustomerSettings.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
	rootCmd.AddCommand(cmdProfilesGetCustomerSettings)

	cmdProfilesGetCustomerSettingsByID := &cobra.Command{
		Use:   "ProfilesGetCustomerSettingsByID",
		Short: "ProfilesGetCustomerSettingsByID",
		Long:  "ProfilesGetCustomerSettingsByID",
		Run: func(cmd *cobra.Command, args []string) {
			var input profiles_cdn.ProfilesGetCustomerSettingsByIDParams
			err := parseInputParams(cmd, &input)
			if err != nil {
				log.Fatal(err)
			}
			displayResponse(input)
		},
	}
	cmdProfilesGetCustomerSettingsByID.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
	rootCmd.AddCommand(cmdProfilesGetCustomerSettingsByID)

	cmdProfilesUpdateCustomerSetting := &cobra.Command{
		Use:   "ProfilesUpdateCustomerSetting",
		Short: "ProfilesUpdateCustomerSetting",
		Long:  "ProfilesUpdateCustomerSetting",
		Run: func(cmd *cobra.Command, args []string) {
			var input profiles_cdn.ProfilesUpdateCustomerSettingParams
			err := parseInputParams(cmd, &input)
			if err != nil {
				log.Fatal(err)
			}
			displayResponse(input)
		},
	}
	cmdProfilesUpdateCustomerSetting.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
	rootCmd.AddCommand(cmdProfilesUpdateCustomerSetting)

	return rootCmd
}

package profiles_rl

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/EdgeCast/ec-sdk-go/edgecast/rtld/profiles_rl"
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

// Root returns the profiles_rl root command
func Root() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "profiles_rl",
		Short: "profiles_rl",
		Long:  "profiles_rl",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	cmdProfilesRateLimitingAddCustomerSetting := &cobra.Command{
		Use:   "ProfilesRateLimitingAddCustomerSetting",
		Short: "ProfilesRateLimitingAddCustomerSetting",
		Long:  "ProfilesRateLimitingAddCustomerSetting",
		Run: func(cmd *cobra.Command, args []string) {
			var input profiles_rl.ProfilesRateLimitingAddCustomerSettingParams
			err := parseInputParams(cmd, &input)
			if err != nil {
				log.Fatal(err)
			}
			displayResponse(input)
		},
	}
	cmdProfilesRateLimitingAddCustomerSetting.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
	rootCmd.AddCommand(cmdProfilesRateLimitingAddCustomerSetting)

	cmdProfilesRateLimitingGetCustomerSettings := &cobra.Command{
		Use:   "ProfilesRateLimitingGetCustomerSettings",
		Short: "ProfilesRateLimitingGetCustomerSettings",
		Long:  "ProfilesRateLimitingGetCustomerSettings",
		Run: func(cmd *cobra.Command, args []string) {
			var input profiles_rl.ProfilesRateLimitingGetCustomerSettingsParams
			err := parseInputParams(cmd, &input)
			if err != nil {
				log.Fatal(err)
			}
			displayResponse(input)
		},
	}
	cmdProfilesRateLimitingGetCustomerSettings.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
	rootCmd.AddCommand(cmdProfilesRateLimitingGetCustomerSettings)

	cmdProfilesRlDeleteCustomerSettingsByID := &cobra.Command{
		Use:   "ProfilesRlDeleteCustomerSettingsByID",
		Short: "ProfilesRlDeleteCustomerSettingsByID",
		Long:  "ProfilesRlDeleteCustomerSettingsByID",
		Run: func(cmd *cobra.Command, args []string) {
			var input profiles_rl.ProfilesRlDeleteCustomerSettingsByIDParams
			err := parseInputParams(cmd, &input)
			if err != nil {
				log.Fatal(err)
			}
			displayResponse(input)
		},
	}
	cmdProfilesRlDeleteCustomerSettingsByID.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
	rootCmd.AddCommand(cmdProfilesRlDeleteCustomerSettingsByID)

	cmdProfilesRlGetCustomerSettingsByID := &cobra.Command{
		Use:   "ProfilesRlGetCustomerSettingsByID",
		Short: "ProfilesRlGetCustomerSettingsByID",
		Long:  "ProfilesRlGetCustomerSettingsByID",
		Run: func(cmd *cobra.Command, args []string) {
			var input profiles_rl.ProfilesRlGetCustomerSettingsByIDParams
			err := parseInputParams(cmd, &input)
			if err != nil {
				log.Fatal(err)
			}
			displayResponse(input)
		},
	}
	cmdProfilesRlGetCustomerSettingsByID.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
	rootCmd.AddCommand(cmdProfilesRlGetCustomerSettingsByID)

	cmdProfilesRlUpdateCustomerSetting := &cobra.Command{
		Use:   "ProfilesRlUpdateCustomerSetting",
		Short: "ProfilesRlUpdateCustomerSetting",
		Long:  "ProfilesRlUpdateCustomerSetting",
		Run: func(cmd *cobra.Command, args []string) {
			var input profiles_rl.ProfilesRlUpdateCustomerSettingParams
			err := parseInputParams(cmd, &input)
			if err != nil {
				log.Fatal(err)
			}
			displayResponse(input)
		},
	}
	cmdProfilesRlUpdateCustomerSetting.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
	rootCmd.AddCommand(cmdProfilesRlUpdateCustomerSetting)

	return rootCmd
}

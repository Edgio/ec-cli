package profiles_waf

import (
    "encoding/json"
    "fmt"
    "log"

    "github.com/EdgeCast/ec-sdk-go/edgecast/rtld/profiles_waf"
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


// Root returns the profiles_waf root command
func Root() *cobra.Command {
    rootCmd := &cobra.Command{
        Use:   "profiles_waf",
        Short: "profiles_waf",
        Long: "profiles_waf",
        Run: func(cmd *cobra.Command, args []string) {
            err := cmd.Help()
            if err != nil {
                log.Fatal(err)
            }
        },
    }
    
    
    cmdProfilesWafAddCustomerSetting := &cobra.Command{
        Use:   "ProfilesWafAddCustomerSetting",
        Short: "ProfilesWafAddCustomerSetting",
        Long:  "ProfilesWafAddCustomerSetting",
        Run: func(cmd *cobra.Command, args []string) {
            var input profiles_waf.ProfilesWafAddCustomerSettingParams
            err := parseInputParams(cmd, &input)
            if err != nil {
                log.Fatal(err)
            }
            displayResponse(input)
        },
    }
    cmdProfilesWafAddCustomerSetting.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
    rootCmd.AddCommand(cmdProfilesWafAddCustomerSetting)
    
    cmdProfilesWafDeleteCustomerSettingsByID := &cobra.Command{
        Use:   "ProfilesWafDeleteCustomerSettingsByID",
        Short: "ProfilesWafDeleteCustomerSettingsByID",
        Long:  "ProfilesWafDeleteCustomerSettingsByID",
        Run: func(cmd *cobra.Command, args []string) {
            var input profiles_waf.ProfilesWafDeleteCustomerSettingsByIDParams
            err := parseInputParams(cmd, &input)
            if err != nil {
                log.Fatal(err)
            }
            displayResponse(input)
        },
    }
    cmdProfilesWafDeleteCustomerSettingsByID.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
    rootCmd.AddCommand(cmdProfilesWafDeleteCustomerSettingsByID)
    
    cmdProfilesWafGetCustomerSettings := &cobra.Command{
        Use:   "ProfilesWafGetCustomerSettings",
        Short: "ProfilesWafGetCustomerSettings",
        Long:  "ProfilesWafGetCustomerSettings",
        Run: func(cmd *cobra.Command, args []string) {
            var input profiles_waf.ProfilesWafGetCustomerSettingsParams
            err := parseInputParams(cmd, &input)
            if err != nil {
                log.Fatal(err)
            }
            displayResponse(input)
        },
    }
    cmdProfilesWafGetCustomerSettings.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
    rootCmd.AddCommand(cmdProfilesWafGetCustomerSettings)
    
    cmdProfilesWafGetCustomerSettingsByID := &cobra.Command{
        Use:   "ProfilesWafGetCustomerSettingsByID",
        Short: "ProfilesWafGetCustomerSettingsByID",
        Long:  "ProfilesWafGetCustomerSettingsByID",
        Run: func(cmd *cobra.Command, args []string) {
            var input profiles_waf.ProfilesWafGetCustomerSettingsByIDParams
            err := parseInputParams(cmd, &input)
            if err != nil {
                log.Fatal(err)
            }
            displayResponse(input)
        },
    }
    cmdProfilesWafGetCustomerSettingsByID.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
    rootCmd.AddCommand(cmdProfilesWafGetCustomerSettingsByID)
    
    cmdProfilesWafUpdateCustomerSetting := &cobra.Command{
        Use:   "ProfilesWafUpdateCustomerSetting",
        Short: "ProfilesWafUpdateCustomerSetting",
        Long:  "ProfilesWafUpdateCustomerSetting",
        Run: func(cmd *cobra.Command, args []string) {
            var input profiles_waf.ProfilesWafUpdateCustomerSettingParams
            err := parseInputParams(cmd, &input)
            if err != nil {
                log.Fatal(err)
            }
            displayResponse(input)
        },
    }
    cmdProfilesWafUpdateCustomerSetting.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
    rootCmd.AddCommand(cmdProfilesWafUpdateCustomerSetting)
    
    return rootCmd
}

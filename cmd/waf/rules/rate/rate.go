// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0 
// license. See LICENSE file in project root for terms.
package rate

import (
    "github.com/EdgeCast/ec-sdk-go/edgecast"
    "github.com/EdgeCast/ec-sdk-go/edgecast/waf"
    "github.com/EdgeCast/ec-sdk-go/edgecast/waf/rules/rate"
    "github.com/edgecast/ec-cli/cmd/internal"
    "github.com/spf13/cobra"
)

func createClient() *waf.WafService {
	idsCredentials := edgecast.IDSCredentials{
		ClientID:     "",
		ClientSecret: "",
		Scope:        "",
	}

	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.IDSCredentials = idsCredentials

	return internal.CheckResult(
		waf.New(sdkConfig),
	)
}

// Root returns the rate root command
func Root() *cobra.Command {
    rootCmd := &cobra.Command{
        Use:   "rate",
        Short: "rate",
        Long: "rate",
        Run: func(cmd *cobra.Command, args []string) {
			internal.Check(
				cmd.Help(),
			)
		},
    }
    
    cmdAddRateRule := &cobra.Command{
        Use:   "AddRateRule",
        Short: "AddRateRule",
        Long:  "AddRateRule",
        Run: func(cmd *cobra.Command, args []string) {
            wafService := createClient()
            var params rate.AddRateRuleParams
            internal.ParseInputParams(cmd, &params)
            result := internal.CheckResult(
				wafService.Rate.AddRateRule(params),
			)
			internal.DisplayResponse(result)
            
            },
    }
    
    cmdAddRateRule.Flags().StringP(
        "input-params-json", 
        "i", 
        "{}", 
        "input data formatted in JSON")
    rootCmd.AddCommand(cmdAddRateRule)
    
    cmdDeleteRateRule := &cobra.Command{
        Use:   "DeleteRateRule",
        Short: "DeleteRateRule",
        Long:  "DeleteRateRule",
        Run: func(cmd *cobra.Command, args []string) {
            wafService := createClient()
            var params rate.DeleteRateRuleParams
            internal.ParseInputParams(cmd, &params)
            internal.Check(
				wafService.Rate.DeleteRateRule(params),
			)
            
            },
    }
    
    cmdDeleteRateRule.Flags().StringP(
        "input-params-json", 
        "i", 
        "{}", 
        "input data formatted in JSON")
    rootCmd.AddCommand(cmdDeleteRateRule)
    
    cmdGetAllRateRules := &cobra.Command{
        Use:   "GetAllRateRules",
        Short: "GetAllRateRules",
        Long:  "GetAllRateRules",
        Run: func(cmd *cobra.Command, args []string) {
            wafService := createClient()
            var params rate.GetAllRateRulesParams
            internal.ParseInputParams(cmd, &params)
            result := internal.CheckResult(
				wafService.Rate.GetAllRateRules(params),
			)
			internal.DisplayResponse(result)
            
            },
    }
    
    cmdGetAllRateRules.Flags().StringP(
        "input-params-json", 
        "i", 
        "{}", 
        "input data formatted in JSON")
    rootCmd.AddCommand(cmdGetAllRateRules)
    
    cmdGetRateRule := &cobra.Command{
        Use:   "GetRateRule",
        Short: "GetRateRule",
        Long:  "GetRateRule",
        Run: func(cmd *cobra.Command, args []string) {
            wafService := createClient()
            var params rate.GetRateRuleParams
            internal.ParseInputParams(cmd, &params)
            result := internal.CheckResult(
				wafService.Rate.GetRateRule(params),
			)
			internal.DisplayResponse(result)
            
            },
    }
    
    cmdGetRateRule.Flags().StringP(
        "input-params-json", 
        "i", 
        "{}", 
        "input data formatted in JSON")
    rootCmd.AddCommand(cmdGetRateRule)
    
    cmdUpdateRateRule := &cobra.Command{
        Use:   "UpdateRateRule",
        Short: "UpdateRateRule",
        Long:  "UpdateRateRule",
        Run: func(cmd *cobra.Command, args []string) {
            wafService := createClient()
            var params rate.UpdateRateRuleParams
            internal.ParseInputParams(cmd, &params)
            internal.Check(
				wafService.Rate.UpdateRateRule(params),
			)
            
            },
    }
    
    cmdUpdateRateRule.Flags().StringP(
        "input-params-json", 
        "i", 
        "{}", 
        "input data formatted in JSON")
    rootCmd.AddCommand(cmdUpdateRateRule)
    
    return rootCmd
}

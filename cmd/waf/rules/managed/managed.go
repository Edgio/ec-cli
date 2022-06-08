// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0 
// license. See LICENSE file in project root for terms.
package managed

import (
    "github.com/EdgeCast/ec-sdk-go/edgecast"
    "github.com/EdgeCast/ec-sdk-go/edgecast/waf"
    "github.com/EdgeCast/ec-sdk-go/edgecast/waf/rules/managed"
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

// Root returns the managed root command
func Root() *cobra.Command {
    rootCmd := &cobra.Command{
        Use:   "managed",
        Short: "managed",
        Long: "managed",
        Run: func(cmd *cobra.Command, args []string) {
			internal.Check(
				cmd.Help(),
			)
		},
    }
    
    cmdAddManagedRule := &cobra.Command{
        Use:   "AddManagedRule",
        Short: "AddManagedRule",
        Long:  "AddManagedRule",
        Run: func(cmd *cobra.Command, args []string) {
            wafService := createClient()
            var params managed.AddManagedRuleParams
            internal.ParseInputParams(cmd, &params)
            result := internal.CheckResult(
				wafService.Managed.AddManagedRule(params),
			)
			internal.DisplayResponse(result)
            
            },
    }
    
    cmdAddManagedRule.Flags().StringP(
        "input-params-json", 
        "i", 
        "{}", 
        "input data formatted in JSON")
    rootCmd.AddCommand(cmdAddManagedRule)
    
    cmdDeleteManagedRule := &cobra.Command{
        Use:   "DeleteManagedRule",
        Short: "DeleteManagedRule",
        Long:  "DeleteManagedRule",
        Run: func(cmd *cobra.Command, args []string) {
            wafService := createClient()
            var params managed.DeleteManagedRuleParams
            internal.ParseInputParams(cmd, &params)
            internal.Check(
				wafService.Managed.DeleteManagedRule(params),
			)
            
            },
    }
    
    cmdDeleteManagedRule.Flags().StringP(
        "input-params-json", 
        "i", 
        "{}", 
        "input data formatted in JSON")
    rootCmd.AddCommand(cmdDeleteManagedRule)
    
    cmdGetAllManagedRules := &cobra.Command{
        Use:   "GetAllManagedRules",
        Short: "GetAllManagedRules",
        Long:  "GetAllManagedRules",
        Run: func(cmd *cobra.Command, args []string) {
            wafService := createClient()
            var params managed.GetAllManagedRulesParams
            internal.ParseInputParams(cmd, &params)
            result := internal.CheckResult(
				wafService.Managed.GetAllManagedRules(params),
			)
			internal.DisplayResponse(result)
            
            },
    }
    
    cmdGetAllManagedRules.Flags().StringP(
        "input-params-json", 
        "i", 
        "{}", 
        "input data formatted in JSON")
    rootCmd.AddCommand(cmdGetAllManagedRules)
    
    cmdGetManagedRule := &cobra.Command{
        Use:   "GetManagedRule",
        Short: "GetManagedRule",
        Long:  "GetManagedRule",
        Run: func(cmd *cobra.Command, args []string) {
            wafService := createClient()
            var params managed.GetManagedRuleParams
            internal.ParseInputParams(cmd, &params)
            result := internal.CheckResult(
				wafService.Managed.GetManagedRule(params),
			)
			internal.DisplayResponse(result)
            
            },
    }
    
    cmdGetManagedRule.Flags().StringP(
        "input-params-json", 
        "i", 
        "{}", 
        "input data formatted in JSON")
    rootCmd.AddCommand(cmdGetManagedRule)
    
    cmdUpdateManagedRule := &cobra.Command{
        Use:   "UpdateManagedRule",
        Short: "UpdateManagedRule",
        Long:  "UpdateManagedRule",
        Run: func(cmd *cobra.Command, args []string) {
            wafService := createClient()
            var params managed.UpdateManagedRuleParams
            internal.ParseInputParams(cmd, &params)
            internal.Check(
				wafService.Managed.UpdateManagedRule(params),
			)
            
            },
    }
    
    cmdUpdateManagedRule.Flags().StringP(
        "input-params-json", 
        "i", 
        "{}", 
        "input data formatted in JSON")
    rootCmd.AddCommand(cmdUpdateManagedRule)
    
    return rootCmd
}

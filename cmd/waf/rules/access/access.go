// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0 
// license. See LICENSE file in project root for terms.
package access

import (
    "github.com/EdgeCast/ec-sdk-go/edgecast"
    "github.com/EdgeCast/ec-sdk-go/edgecast/waf"
    "github.com/EdgeCast/ec-sdk-go/edgecast/waf/rules/access"
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

// Root returns the access root command
func Root() *cobra.Command {
    rootCmd := &cobra.Command{
        Use:   "access",
        Short: "access",
        Long: "access",
        Run: func(cmd *cobra.Command, args []string) {
			internal.Check(
				cmd.Help(),
			)
		},
    }
    
    cmdAddAccessRule := &cobra.Command{
        Use:   "AddAccessRule",
        Short: "AddAccessRule",
        Long:  "AddAccessRule",
        Run: func(cmd *cobra.Command, args []string) {
            wafService := createClient()
            var params access.AddAccessRuleParams
            internal.ParseInputParams(cmd, &params)
            result := internal.CheckResult(
				wafService.Access.AddAccessRule(params),
			)
			internal.DisplayResponse(result)
            
            },
    }
    
    cmdAddAccessRule.Flags().StringP(
        "input-params-json", 
        "i", 
        "{}", 
        "input data formatted in JSON")
    rootCmd.AddCommand(cmdAddAccessRule)
    
    cmdDeleteAccessRule := &cobra.Command{
        Use:   "DeleteAccessRule",
        Short: "DeleteAccessRule",
        Long:  "DeleteAccessRule",
        Run: func(cmd *cobra.Command, args []string) {
            wafService := createClient()
            var params access.DeleteAccessRuleParams
            internal.ParseInputParams(cmd, &params)
            internal.Check(
				wafService.Access.DeleteAccessRule(params),
			)
            
            },
    }
    
    cmdDeleteAccessRule.Flags().StringP(
        "input-params-json", 
        "i", 
        "{}", 
        "input data formatted in JSON")
    rootCmd.AddCommand(cmdDeleteAccessRule)
    
    cmdGetAccessRule := &cobra.Command{
        Use:   "GetAccessRule",
        Short: "GetAccessRule",
        Long:  "GetAccessRule",
        Run: func(cmd *cobra.Command, args []string) {
            wafService := createClient()
            var params access.GetAccessRuleParams
            internal.ParseInputParams(cmd, &params)
            result := internal.CheckResult(
				wafService.Access.GetAccessRule(params),
			)
			internal.DisplayResponse(result)
            
            },
    }
    
    cmdGetAccessRule.Flags().StringP(
        "input-params-json", 
        "i", 
        "{}", 
        "input data formatted in JSON")
    rootCmd.AddCommand(cmdGetAccessRule)
    
    cmdGetAllAccessRules := &cobra.Command{
        Use:   "GetAllAccessRules",
        Short: "GetAllAccessRules",
        Long:  "GetAllAccessRules",
        Run: func(cmd *cobra.Command, args []string) {
            wafService := createClient()
            var params access.GetAllAccessRulesParams
            internal.ParseInputParams(cmd, &params)
            result := internal.CheckResult(
				wafService.Access.GetAllAccessRules(params),
			)
			internal.DisplayResponse(result)
            
            },
    }
    
    cmdGetAllAccessRules.Flags().StringP(
        "input-params-json", 
        "i", 
        "{}", 
        "input data formatted in JSON")
    rootCmd.AddCommand(cmdGetAllAccessRules)
    
    cmdUpdateAccessRule := &cobra.Command{
        Use:   "UpdateAccessRule",
        Short: "UpdateAccessRule",
        Long:  "UpdateAccessRule",
        Run: func(cmd *cobra.Command, args []string) {
            wafService := createClient()
            var params access.UpdateAccessRuleParams
            internal.ParseInputParams(cmd, &params)
            internal.Check(
				wafService.Access.UpdateAccessRule(params),
			)
            
            },
    }
    
    cmdUpdateAccessRule.Flags().StringP(
        "input-params-json", 
        "i", 
        "{}", 
        "input data formatted in JSON")
    rootCmd.AddCommand(cmdUpdateAccessRule)
    
    return rootCmd
}

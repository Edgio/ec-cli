// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0 
// license. See LICENSE file in project root for terms.
package custom

// This file is auto-generated. Modifications to this file may be overwritten.

import (
    "github.com/EdgeCast/ec-sdk-go/edgecast"
    "github.com/EdgeCast/ec-sdk-go/edgecast/waf"
    "github.com/EdgeCast/ec-sdk-go/edgecast/waf/rules/custom"
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

// Root returns the custom root command.
func Root() *cobra.Command {
    rootCmd := createRootCmd()
    
    registerAddCustomRuleSet(rootCmd)
    registerDeleteCustomRuleSet(rootCmd)
    registerGetAllCustomRuleSets(rootCmd)
    registerGetCustomRuleSet(rootCmd)
    registerUpdateCustomRuleSet(rootCmd)
    
    return rootCmd
}

func createRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "access",
		Short: "access",
		Long:  "access",
		Run: func(cmd *cobra.Command, args []string) {
			internal.Check(
				cmd.Help(),
			)
		},
	}
}

func registerAddCustomRuleSet(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
        Use:   "AddCustomRuleSet",
        Short: "AddCustomRuleSet",
        Long:  "AddCustomRuleSet",
        Run: func(cmd *cobra.Command, args []string) {
            wafService := createClient()
            var params custom.AddCustomRuleSetParams
            internal.ParseInputParams(cmd, &params)
            result := internal.CheckResult(
				wafService.Custom.AddCustomRuleSet(params),
			)
			internal.DisplayResponse(result)
            
            },
    }
    
    cmd.Flags().StringP(
        "input-params-json", 
        "i", 
        "{}", 
        "input data formatted in JSON")

    rootCmd.AddCommand(cmd)
}
func registerDeleteCustomRuleSet(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
        Use:   "DeleteCustomRuleSet",
        Short: "DeleteCustomRuleSet",
        Long:  "DeleteCustomRuleSet",
        Run: func(cmd *cobra.Command, args []string) {
            wafService := createClient()
            var params custom.DeleteCustomRuleSetParams
            internal.ParseInputParams(cmd, &params)
            internal.Check(
				wafService.Custom.DeleteCustomRuleSet(params),
			)
            
            },
    }
    
    cmd.Flags().StringP(
        "input-params-json", 
        "i", 
        "{}", 
        "input data formatted in JSON")

    rootCmd.AddCommand(cmd)
}
func registerGetAllCustomRuleSets(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
        Use:   "GetAllCustomRuleSets",
        Short: "GetAllCustomRuleSets",
        Long:  "GetAllCustomRuleSets",
        Run: func(cmd *cobra.Command, args []string) {
            wafService := createClient()
            var params custom.GetAllCustomRuleSetsParams
            internal.ParseInputParams(cmd, &params)
            result := internal.CheckResult(
				wafService.Custom.GetAllCustomRuleSets(params),
			)
			internal.DisplayResponse(result)
            
            },
    }
    
    cmd.Flags().StringP(
        "input-params-json", 
        "i", 
        "{}", 
        "input data formatted in JSON")

    rootCmd.AddCommand(cmd)
}
func registerGetCustomRuleSet(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
        Use:   "GetCustomRuleSet",
        Short: "GetCustomRuleSet",
        Long:  "GetCustomRuleSet",
        Run: func(cmd *cobra.Command, args []string) {
            wafService := createClient()
            var params custom.GetCustomRuleSetParams
            internal.ParseInputParams(cmd, &params)
            result := internal.CheckResult(
				wafService.Custom.GetCustomRuleSet(params),
			)
			internal.DisplayResponse(result)
            
            },
    }
    
    cmd.Flags().StringP(
        "input-params-json", 
        "i", 
        "{}", 
        "input data formatted in JSON")

    rootCmd.AddCommand(cmd)
}
func registerUpdateCustomRuleSet(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
        Use:   "UpdateCustomRuleSet",
        Short: "UpdateCustomRuleSet",
        Long:  "UpdateCustomRuleSet",
        Run: func(cmd *cobra.Command, args []string) {
            wafService := createClient()
            var params custom.UpdateCustomRuleSetParams
            internal.ParseInputParams(cmd, &params)
            internal.Check(
				wafService.Custom.UpdateCustomRuleSet(params),
			)
            
            },
    }
    
    cmd.Flags().StringP(
        "input-params-json", 
        "i", 
        "{}", 
        "input data formatted in JSON")

    rootCmd.AddCommand(cmd)
}
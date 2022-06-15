// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0 
// license. See LICENSE file in project root for terms.
package bot

// This file is auto-generated. Modifications to this file may be overwritten.

import (
    "github.com/EdgeCast/ec-sdk-go/edgecast"
    "github.com/EdgeCast/ec-sdk-go/edgecast/waf"
    "github.com/EdgeCast/ec-sdk-go/edgecast/waf/rules/bot"
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

// Root returns the bot root command.
func Root() *cobra.Command {
    rootCmd := createRootCmd()
    
    registerAddBotRuleSet(rootCmd)
    registerDeleteBotRuleSet(rootCmd)
    registerGetAllBotRuleSets(rootCmd)
    registerGetBotRuleSet(rootCmd)
    registerUpdateBotRuleSet(rootCmd)
    
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

func registerAddBotRuleSet(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
        Use:   "AddBotRuleSet",
        Short: "AddBotRuleSet",
        Long:  "AddBotRuleSet",
        Run: func(cmd *cobra.Command, args []string) {
            wafService := createClient()
            var params bot.AddBotRuleSetParams
            internal.ParseInputParams(cmd, &params)
            result := internal.CheckResult(
				wafService.Bot.AddBotRuleSet(params),
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
func registerDeleteBotRuleSet(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
        Use:   "DeleteBotRuleSet",
        Short: "DeleteBotRuleSet",
        Long:  "DeleteBotRuleSet",
        Run: func(cmd *cobra.Command, args []string) {
            wafService := createClient()
            var params bot.DeleteBotRuleSetParams
            internal.ParseInputParams(cmd, &params)
            internal.Check(
				wafService.Bot.DeleteBotRuleSet(params),
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
func registerGetAllBotRuleSets(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
        Use:   "GetAllBotRuleSets",
        Short: "GetAllBotRuleSets",
        Long:  "GetAllBotRuleSets",
        Run: func(cmd *cobra.Command, args []string) {
            wafService := createClient()
            var params bot.GetAllBotRuleSetsParams
            internal.ParseInputParams(cmd, &params)
            result := internal.CheckResult(
				wafService.Bot.GetAllBotRuleSets(params),
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
func registerGetBotRuleSet(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
        Use:   "GetBotRuleSet",
        Short: "GetBotRuleSet",
        Long:  "GetBotRuleSet",
        Run: func(cmd *cobra.Command, args []string) {
            wafService := createClient()
            var params bot.GetBotRuleSetParams
            internal.ParseInputParams(cmd, &params)
            result := internal.CheckResult(
				wafService.Bot.GetBotRuleSet(params),
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
func registerUpdateBotRuleSet(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
        Use:   "UpdateBotRuleSet",
        Short: "UpdateBotRuleSet",
        Long:  "UpdateBotRuleSet",
        Run: func(cmd *cobra.Command, args []string) {
            wafService := createClient()
            var params bot.UpdateBotRuleSetParams
            internal.ParseInputParams(cmd, &params)
            internal.Check(
				wafService.Bot.UpdateBotRuleSet(params),
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
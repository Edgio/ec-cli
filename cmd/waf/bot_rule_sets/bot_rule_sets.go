// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.
package bot_rule_sets

import (
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf/bot_rule_sets"
	"github.com/edgecast/ec-cli/cmd/internal"
	"github.com/spf13/cobra"
)

const (
	jsonUsage         = "input data formatted in JSON"
	jsonFlagName      = "input-params-json"
	jsonFlagShortName = "i"
	jsonFlagValue     = "{}"
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

// Root returns the bot_rule_sets root command
func Root() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "bot_rule_sets",
		Short: "bot_rule_sets",
		Long:  "bot_rule_sets",
		Run: func(cmd *cobra.Command, args []string) {
			internal.Check(
				cmd.Help(),
			)
		},
	}

	cmdAddBotRuleSet := &cobra.Command{
		Use:   "AddBotRuleSet",
		Short: "AddBotRuleSet",
		Long:  "AddBotRuleSet",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params *bot_rule_sets.AddBotRuleSetParams
			internal.ParseInputParams(cmd, params)
			result := internal.CheckResult(
				wafService.BotRuleSets.AddBotRuleSet(params),
			)
			internal.DisplayResponse(result)

		},
	}

	cmdAddBotRuleSet.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdAddBotRuleSet)

	cmdDeleteBotRuleSet := &cobra.Command{
		Use:   "DeleteBotRuleSet",
		Short: "DeleteBotRuleSet",
		Long:  "DeleteBotRuleSet",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params *bot_rule_sets.DeleteBotRuleSetParams
			internal.ParseInputParams(cmd, params)
			internal.Check(
				wafService.BotRuleSets.DeleteBotRuleSet(params),
			)

		},
	}

	cmdDeleteBotRuleSet.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdDeleteBotRuleSet)

	cmdGetAllBotRuleSets := &cobra.Command{
		Use:   "GetAllBotRuleSets",
		Short: "GetAllBotRuleSets",
		Long:  "GetAllBotRuleSets",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params *bot_rule_sets.GetAllBotRuleSetsParams
			internal.ParseInputParams(cmd, params)
			result := internal.CheckResult(
				wafService.BotRuleSets.GetAllBotRuleSets(params),
			)
			internal.DisplayResponse(result)

		},
	}

	cmdGetAllBotRuleSets.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdGetAllBotRuleSets)

	cmdGetBotRuleSet := &cobra.Command{
		Use:   "GetBotRuleSet",
		Short: "GetBotRuleSet",
		Long:  "GetBotRuleSet",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params *bot_rule_sets.GetBotRuleSetParams
			internal.ParseInputParams(cmd, params)
			result := internal.CheckResult(
				wafService.BotRuleSets.GetBotRuleSet(params),
			)
			internal.DisplayResponse(result)

		},
	}

	cmdGetBotRuleSet.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdGetBotRuleSet)

	cmdUpdateBotRuleSet := &cobra.Command{
		Use:   "UpdateBotRuleSet",
		Short: "UpdateBotRuleSet",
		Long:  "UpdateBotRuleSet",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params *bot_rule_sets.UpdateBotRuleSetParams
			internal.ParseInputParams(cmd, params)
			internal.Check(
				wafService.BotRuleSets.UpdateBotRuleSet(params),
			)

		},
	}

	cmdUpdateBotRuleSet.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdUpdateBotRuleSet)

	return rootCmd
}

// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.
package custom_rule_sets

import (
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf/custom_rule_sets"
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

// Root returns the custom_rule_sets root command
func Root() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "custom_rule_sets",
		Short: "custom_rule_sets",
		Long:  "custom_rule_sets",
		Run: func(cmd *cobra.Command, args []string) {
			internal.Check(
				cmd.Help(),
			)
		},
	}

	cmdAddCustomRuleSet := &cobra.Command{
		Use:   "AddCustomRuleSet",
		Short: "AddCustomRuleSet",
		Long:  "AddCustomRuleSet",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params *custom_rule_sets.AddCustomRuleSetParams
			internal.ParseInputParams(cmd, params)
			result := internal.CheckResult(
				wafService.CustomRuleSets.AddCustomRuleSet(params),
			)
			internal.DisplayResponse(result)

		},
	}

	cmdAddCustomRuleSet.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdAddCustomRuleSet)

	cmdDeleteCustomRuleSet := &cobra.Command{
		Use:   "DeleteCustomRuleSet",
		Short: "DeleteCustomRuleSet",
		Long:  "DeleteCustomRuleSet",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params *custom_rule_sets.DeleteCustomRuleSetParams
			internal.ParseInputParams(cmd, params)
			internal.Check(
				wafService.CustomRuleSets.DeleteCustomRuleSet(params),
			)

		},
	}

	cmdDeleteCustomRuleSet.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdDeleteCustomRuleSet)

	cmdGetAllCustomRuleSets := &cobra.Command{
		Use:   "GetAllCustomRuleSets",
		Short: "GetAllCustomRuleSets",
		Long:  "GetAllCustomRuleSets",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params *custom_rule_sets.GetAllCustomRuleSetsParams
			internal.ParseInputParams(cmd, params)
			result := internal.CheckResult(
				wafService.CustomRuleSets.GetAllCustomRuleSets(params),
			)
			internal.DisplayResponse(result)

		},
	}

	cmdGetAllCustomRuleSets.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdGetAllCustomRuleSets)

	cmdGetCustomRuleSet := &cobra.Command{
		Use:   "GetCustomRuleSet",
		Short: "GetCustomRuleSet",
		Long:  "GetCustomRuleSet",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params *custom_rule_sets.GetCustomRuleSetParams
			internal.ParseInputParams(cmd, params)
			result := internal.CheckResult(
				wafService.CustomRuleSets.GetCustomRuleSet(params),
			)
			internal.DisplayResponse(result)

		},
	}

	cmdGetCustomRuleSet.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdGetCustomRuleSet)

	cmdUpdateCustomRuleSet := &cobra.Command{
		Use:   "UpdateCustomRuleSet",
		Short: "UpdateCustomRuleSet",
		Long:  "UpdateCustomRuleSet",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params *custom_rule_sets.UpdateCustomRuleSetParams
			internal.ParseInputParams(cmd, params)
			internal.Check(
				wafService.CustomRuleSets.UpdateCustomRuleSet(params),
			)

		},
	}

	cmdUpdateCustomRuleSet.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdUpdateCustomRuleSet)

	return rootCmd
}

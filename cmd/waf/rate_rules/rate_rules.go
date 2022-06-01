// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.
package rate_rules

import (
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf/rate_rules"
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

// Root returns the rate_rules root command
func Root() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "rate_rules",
		Short: "rate_rules",
		Long:  "rate_rules",
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
			var params *rate_rules.AddRateRuleParams
			internal.ParseInputParams(cmd, params)
			result := internal.CheckResult(
				wafService.RateRules.AddRateRule(params),
			)
			internal.DisplayResponse(result)

		},
	}

	cmdAddRateRule.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdAddRateRule)

	cmdDeleteRateRule := &cobra.Command{
		Use:   "DeleteRateRule",
		Short: "DeleteRateRule",
		Long:  "DeleteRateRule",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params *rate_rules.DeleteRateRuleParams
			internal.ParseInputParams(cmd, params)
			internal.Check(
				wafService.RateRules.DeleteRateRule(params),
			)

		},
	}

	cmdDeleteRateRule.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdDeleteRateRule)

	cmdGetAllRateRules := &cobra.Command{
		Use:   "GetAllRateRules",
		Short: "GetAllRateRules",
		Long:  "GetAllRateRules",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params *rate_rules.GetAllRateRulesParams
			internal.ParseInputParams(cmd, params)
			result := internal.CheckResult(
				wafService.RateRules.GetAllRateRules(params),
			)
			internal.DisplayResponse(result)

		},
	}

	cmdGetAllRateRules.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdGetAllRateRules)

	cmdGetRateRule := &cobra.Command{
		Use:   "GetRateRule",
		Short: "GetRateRule",
		Long:  "GetRateRule",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params *rate_rules.GetRateRuleParams
			internal.ParseInputParams(cmd, params)
			result := internal.CheckResult(
				wafService.RateRules.GetRateRule(params),
			)
			internal.DisplayResponse(result)

		},
	}

	cmdGetRateRule.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdGetRateRule)

	cmdUpdateRateRule := &cobra.Command{
		Use:   "UpdateRateRule",
		Short: "UpdateRateRule",
		Long:  "UpdateRateRule",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params *rate_rules.UpdateRateRuleParams
			internal.ParseInputParams(cmd, params)
			internal.Check(
				wafService.RateRules.UpdateRateRule(params),
			)

		},
	}

	cmdUpdateRateRule.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdUpdateRateRule)

	return rootCmd
}

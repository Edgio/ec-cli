// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.
package managed_rules

import (
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf/managed_rules"
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

// Root returns the managed_rules root command
func Root() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "managed_rules",
		Short: "managed_rules",
		Long:  "managed_rules",
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
			var params *managed_rules.AddManagedRuleParams
			internal.ParseInputParams(cmd, params)
			result := internal.CheckResult(
				wafService.ManagedRules.AddManagedRule(params),
			)
			internal.DisplayResponse(result)

		},
	}

	cmdAddManagedRule.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdAddManagedRule)

	cmdDeleteManagedRule := &cobra.Command{
		Use:   "DeleteManagedRule",
		Short: "DeleteManagedRule",
		Long:  "DeleteManagedRule",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params *managed_rules.DeleteManagedRuleParams
			internal.ParseInputParams(cmd, params)
			internal.Check(
				wafService.ManagedRules.DeleteManagedRule(params),
			)

		},
	}

	cmdDeleteManagedRule.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdDeleteManagedRule)

	cmdGetAllManagedRules := &cobra.Command{
		Use:   "GetAllManagedRules",
		Short: "GetAllManagedRules",
		Long:  "GetAllManagedRules",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params *managed_rules.GetAllManagedRulesParams
			internal.ParseInputParams(cmd, params)
			result := internal.CheckResult(
				wafService.ManagedRules.GetAllManagedRules(params),
			)
			internal.DisplayResponse(result)

		},
	}

	cmdGetAllManagedRules.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdGetAllManagedRules)

	cmdGetManagedRule := &cobra.Command{
		Use:   "GetManagedRule",
		Short: "GetManagedRule",
		Long:  "GetManagedRule",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params *managed_rules.GetManagedRuleParams
			internal.ParseInputParams(cmd, params)
			result := internal.CheckResult(
				wafService.ManagedRules.GetManagedRule(params),
			)
			internal.DisplayResponse(result)

		},
	}

	cmdGetManagedRule.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdGetManagedRule)

	cmdUpdateManagedRule := &cobra.Command{
		Use:   "UpdateManagedRule",
		Short: "UpdateManagedRule",
		Long:  "UpdateManagedRule",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params *managed_rules.UpdateManagedRuleParams
			internal.ParseInputParams(cmd, params)
			internal.Check(
				wafService.ManagedRules.UpdateManagedRule(params),
			)

		},
	}

	cmdUpdateManagedRule.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdUpdateManagedRule)

	return rootCmd
}

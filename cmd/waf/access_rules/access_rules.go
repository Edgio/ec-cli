// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.
package access_rules

import (
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf/access_rules"
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

// Root returns the access_rules root command
func Root() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "access_rules",
		Short: "access_rules",
		Long:  "access_rules",
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
			var params *access_rules.AddAccessRuleParams
			internal.ParseInputParams(cmd, params)
			result := internal.CheckResult(
				wafService.AccessRules.AddAccessRule(params),
			)
			internal.DisplayResponse(result)

		},
	}

	cmdAddAccessRule.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdAddAccessRule)

	cmdDeleteAccessRule := &cobra.Command{
		Use:   "DeleteAccessRule",
		Short: "DeleteAccessRule",
		Long:  "DeleteAccessRule",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params *access_rules.DeleteAccessRuleParams
			internal.ParseInputParams(cmd, params)
			internal.Check(
				wafService.AccessRules.DeleteAccessRule(params),
			)

		},
	}

	cmdDeleteAccessRule.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdDeleteAccessRule)

	cmdGetAccessRule := &cobra.Command{
		Use:   "GetAccessRule",
		Short: "GetAccessRule",
		Long:  "GetAccessRule",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params *access_rules.GetAccessRuleParams
			internal.ParseInputParams(cmd, params)
			result := internal.CheckResult(
				wafService.AccessRules.GetAccessRule(params),
			)
			internal.DisplayResponse(result)

		},
	}

	cmdGetAccessRule.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdGetAccessRule)

	cmdGetAllAccessRules := &cobra.Command{
		Use:   "GetAllAccessRules",
		Short: "GetAllAccessRules",
		Long:  "GetAllAccessRules",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params *access_rules.GetAllAccessRulesParams
			internal.ParseInputParams(cmd, params)
			result := internal.CheckResult(
				wafService.AccessRules.GetAllAccessRules(params),
			)
			internal.DisplayResponse(result)

		},
	}

	cmdGetAllAccessRules.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdGetAllAccessRules)

	cmdUpdateAccessRule := &cobra.Command{
		Use:   "UpdateAccessRule",
		Short: "UpdateAccessRule",
		Long:  "UpdateAccessRule",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params *access_rules.UpdateAccessRuleParams
			internal.ParseInputParams(cmd, params)
			internal.Check(
				wafService.AccessRules.UpdateAccessRule(params),
			)

		},
	}

	cmdUpdateAccessRule.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdUpdateAccessRule)

	return rootCmd
}

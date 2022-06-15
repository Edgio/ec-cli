// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.
package managed

// This file is auto-generated. Modifications to this file may be overwritten.

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

// Root returns the managed root command.
func Root() *cobra.Command {
	rootCmd := createRootCmd()

	registerAddManagedRule(rootCmd)
	registerDeleteManagedRule(rootCmd)
	registerGetAllManagedRules(rootCmd)
	registerGetManagedRule(rootCmd)
	registerUpdateManagedRule(rootCmd)

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

func registerAddManagedRule(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
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

	cmd.Flags().StringP(
		"input-params-json",
		"i",
		"{}",
		"input data formatted in JSON")

	rootCmd.AddCommand(cmd)
}

func registerDeleteManagedRule(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
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

	cmd.Flags().StringP(
		"input-params-json",
		"i",
		"{}",
		"input data formatted in JSON")

	rootCmd.AddCommand(cmd)
}

func registerGetAllManagedRules(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
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

	cmd.Flags().StringP(
		"input-params-json",
		"i",
		"{}",
		"input data formatted in JSON")

	rootCmd.AddCommand(cmd)
}

func registerGetManagedRule(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
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

	cmd.Flags().StringP(
		"input-params-json",
		"i",
		"{}",
		"input data formatted in JSON")

	rootCmd.AddCommand(cmd)
}

func registerUpdateManagedRule(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
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

	cmd.Flags().StringP(
		"input-params-json",
		"i",
		"{}",
		"input data formatted in JSON")

	rootCmd.AddCommand(cmd)
}

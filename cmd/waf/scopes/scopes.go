// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.
package scopes

// This file is auto-generated. Modifications to this file may be overwritten.

import (
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf/scopes"
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

// Root returns the scopes root command.
func Root() *cobra.Command {
	rootCmd := createRootCmd()

	registerGetAllScopes(rootCmd)
	registerModifyAllScopes(rootCmd)

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

func registerGetAllScopes(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "GetAllScopes",
		Short: "GetAllScopes",
		Long:  "GetAllScopes",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params scopes.GetAllScopesParams
			internal.ParseInputParams(cmd, &params)
			result := internal.CheckResult(
				wafService.Scopes.GetAllScopes(params),
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

func registerModifyAllScopes(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "ModifyAllScopes",
		Short: "ModifyAllScopes",
		Long:  "ModifyAllScopes",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params scopes.Scopes
			internal.ParseInputParams(cmd, &params)
			result := internal.CheckResult(
				wafService.Scopes.ModifyAllScopes(params),
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

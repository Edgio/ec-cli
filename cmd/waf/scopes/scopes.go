// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.
package scopes

import (
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf"
	"github.com/EdgeCast/ec-sdk-go/edgecast/waf/scopes"
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

// Root returns the scopes root command
func Root() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "scopes",
		Short: "scopes",
		Long:  "scopes",
		Run: func(cmd *cobra.Command, args []string) {
			internal.Check(
				cmd.Help(),
			)
		},
	}

	cmdGetAllScopes := &cobra.Command{
		Use:   "GetAllScopes",
		Short: "GetAllScopes",
		Long:  "GetAllScopes",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params *scopes.GetAllScopesParams
			internal.ParseInputParams(cmd, params)
			result := internal.CheckResult(
				wafService.Scopes.GetAllScopes(params),
			)
			internal.DisplayResponse(result)

		},
	}

	cmdGetAllScopes.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdGetAllScopes)

	cmdModifyAllScopes := &cobra.Command{
		Use:   "ModifyAllScopes",
		Short: "ModifyAllScopes",
		Long:  "ModifyAllScopes",
		Run: func(cmd *cobra.Command, args []string) {
			wafService := createClient()
			var params *scopes.Scopes
			internal.ParseInputParams(cmd, params)
			result := internal.CheckResult(
				wafService.Scopes.ModifyAllScopes(params),
			)
			internal.DisplayResponse(result)

		},
	}

	cmdModifyAllScopes.Flags().StringP(
		jsonFlagName,
		jsonFlagShortName,
		jsonFlagValue,
		jsonUsage)

	rootCmd.AddCommand(cmdModifyAllScopes)

	return rootCmd
}

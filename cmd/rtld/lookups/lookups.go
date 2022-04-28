package lookups

import (
	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/rtld"
	"github.com/EdgeCast/ec-sdk-go/edgecast/rtld/lookups"
	"github.com/edgecast/ec-cli/cmd/internal"
	"github.com/spf13/cobra"
)

func createClient() *rtld.RealTimeLogDeliveryAPI {
	idsCredentials := edgecast.IDSCredentials{
		ClientID:     "",
		ClientSecret: "",
		Scope:        "",
	}
	sdkConfig := edgecast.NewSDKConfig()
	sdkConfig.IDSCredentials = idsCredentials

	return internal.CheckResult(
		rtld.New(sdkConfig),
	)
}

// Root returns the lookups root command
func Root() *cobra.Command { //nolint:funlen
	rootCmd := &cobra.Command{
		Use:   "lookups",
		Short: "lookups",
		Long:  "lookups",
		Run: func(cmd *cobra.Command, args []string) {
			internal.Check(
				cmd.Help(),
			)
		},
	}

	cmdLookupsGetAwsRegions := &cobra.Command{
		Use:   "LookupsGetAwsRegions",
		Short: "LookupsGetAwsRegions",
		Long:  "LookupsGetAwsRegions",
		Run: func(cmd *cobra.Command, args []string) {
			var params *lookups.LookupsGetAwsRegionsParams
			internal.ParseInputParams(cmd, params)

			rtldService := createClient()

			result := internal.CheckResult(
				rtldService.Lookups.LookupsGetAwsRegions(params),
			)

			internal.DisplayResponse(result)
		},
	}
	cmdLookupsGetAwsRegions.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
	rootCmd.AddCommand(cmdLookupsGetAwsRegions)

	cmdLookupsGetAzureAccessTypes := &cobra.Command{
		Use:   "LookupsGetAzureAccessTypes",
		Short: "LookupsGetAzureAccessTypes",
		Long:  "LookupsGetAzureAccessTypes",
		Run: func(cmd *cobra.Command, args []string) {
			var params *lookups.LookupsGetAzureAccessTypesParams
			internal.ParseInputParams(cmd, params)

			rtldService := createClient()

			result := internal.CheckResult(
				rtldService.Lookups.LookupsGetAzureAccessTypes(params),
			)

			internal.DisplayResponse(result)
		},
	}
	cmdLookupsGetAzureAccessTypes.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
	rootCmd.AddCommand(cmdLookupsGetAzureAccessTypes)

	cmdLookupsGetCustomItems := &cobra.Command{
		Use:   "LookupsGetCustomItems",
		Short: "LookupsGetCustomItems",
		Long:  "LookupsGetCustomItems",
		Run: func(cmd *cobra.Command, args []string) {
			var params *lookups.LookupsGetCustomItemsParams
			internal.ParseInputParams(cmd, params)

			rtldService := createClient()

			result := internal.CheckResult(
				rtldService.Lookups.LookupsGetCustomItems(params),
			)

			internal.DisplayResponse(result)
		},
	}
	cmdLookupsGetCustomItems.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
	rootCmd.AddCommand(cmdLookupsGetCustomItems)

	cmdLookupsGetDeliveryMethods := &cobra.Command{
		Use:   "LookupsGetDeliveryMethods",
		Short: "LookupsGetDeliveryMethods",
		Long:  "LookupsGetDeliveryMethods",
		Run: func(cmd *cobra.Command, args []string) {
			var params *lookups.LookupsGetDeliveryMethodsParams
			internal.ParseInputParams(cmd, params)

			rtldService := createClient()

			result := internal.CheckResult(
				rtldService.Lookups.LookupsGetDeliveryMethods(params),
			)

			internal.DisplayResponse(result)
		},
	}
	cmdLookupsGetDeliveryMethods.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
	rootCmd.AddCommand(cmdLookupsGetDeliveryMethods)

	cmdLookupsGetDownsamplingRates := &cobra.Command{
		Use:   "LookupsGetDownsamplingRates",
		Short: "LookupsGetDownsamplingRates",
		Long:  "LookupsGetDownsamplingRates",
		Run: func(cmd *cobra.Command, args []string) {
			var params *lookups.LookupsGetDownsamplingRatesParams
			internal.ParseInputParams(cmd, params)

			rtldService := createClient()

			result := internal.CheckResult(
				rtldService.Lookups.LookupsGetDownsamplingRates(params),
			)

			internal.DisplayResponse(result)
		},
	}
	cmdLookupsGetDownsamplingRates.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
	rootCmd.AddCommand(cmdLookupsGetDownsamplingRates)

	cmdLookupsGetFieldRl := &cobra.Command{
		Use:   "LookupsGetFieldRl",
		Short: "LookupsGetFieldRl",
		Long:  "LookupsGetFieldRl",
		Run: func(cmd *cobra.Command, args []string) {
			var params *lookups.LookupsGetFieldRlParams
			internal.ParseInputParams(cmd, params)

			rtldService := createClient()

			result := internal.CheckResult(
				rtldService.Lookups.LookupsGetFieldRl(params),
			)

			internal.DisplayResponse(result)
		},
	}
	cmdLookupsGetFieldRl.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
	rootCmd.AddCommand(cmdLookupsGetFieldRl)

	cmdLookupsGetFieldsCdn := &cobra.Command{
		Use:   "LookupsGetFieldsCdn",
		Short: "LookupsGetFieldsCdn",
		Long:  "LookupsGetFieldsCdn",
		Run: func(cmd *cobra.Command, args []string) {
			var params *lookups.LookupsGetFieldsCdnParams
			internal.ParseInputParams(cmd, params)

			rtldService := createClient()

			result := internal.CheckResult(
				rtldService.Lookups.LookupsGetFieldsCdn(params),
			)

			internal.DisplayResponse(result)
		},
	}
	cmdLookupsGetFieldsCdn.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
	rootCmd.AddCommand(cmdLookupsGetFieldsCdn)

	cmdLookupsGetFieldsWaf := &cobra.Command{
		Use:   "LookupsGetFieldsWaf",
		Short: "LookupsGetFieldsWaf",
		Long:  "LookupsGetFieldsWaf",
		Run: func(cmd *cobra.Command, args []string) {
			var params *lookups.LookupsGetFieldsWafParams
			internal.ParseInputParams(cmd, params)

			rtldService := createClient()

			result := internal.CheckResult(
				rtldService.Lookups.LookupsGetFieldsWaf(params),
			)

			internal.DisplayResponse(result)
		},
	}
	cmdLookupsGetFieldsWaf.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
	rootCmd.AddCommand(cmdLookupsGetFieldsWaf)

	cmdLookupsGetHTTPAuthenticationMethods := &cobra.Command{
		Use:   "LookupsGetHTTPAuthenticationMethods",
		Short: "LookupsGetHTTPAuthenticationMethods",
		Long:  "LookupsGetHTTPAuthenticationMethods",
		Run: func(cmd *cobra.Command, args []string) {
			var params *lookups.LookupsGetHTTPAuthenticationMethodsParams
			internal.ParseInputParams(cmd, params)

			rtldService := createClient()

			result := internal.CheckResult(
				rtldService.Lookups.LookupsGetHTTPAuthenticationMethods(params),
			)

			internal.DisplayResponse(result)
		},
	}
	cmdLookupsGetHTTPAuthenticationMethods.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
	rootCmd.AddCommand(cmdLookupsGetHTTPAuthenticationMethods)

	cmdLookupsGetLogFormats := &cobra.Command{
		Use:   "LookupsGetLogFormats",
		Short: "LookupsGetLogFormats",
		Long:  "LookupsGetLogFormats",
		Run: func(cmd *cobra.Command, args []string) {
			var params *lookups.LookupsGetLogFormatsParams
			internal.ParseInputParams(cmd, params)

			rtldService := createClient()

			result := internal.CheckResult(
				rtldService.Lookups.LookupsGetLogFormats(params),
			)

			internal.DisplayResponse(result)
		},
	}
	cmdLookupsGetLogFormats.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
	rootCmd.AddCommand(cmdLookupsGetLogFormats)

	cmdLookupsGetPlatforms := &cobra.Command{
		Use:   "LookupsGetPlatforms",
		Short: "LookupsGetPlatforms",
		Long:  "LookupsGetPlatforms",
		Run: func(cmd *cobra.Command, args []string) {
			var params *lookups.LookupsGetPlatformsParams
			internal.ParseInputParams(cmd, params)

			rtldService := createClient()

			result := internal.CheckResult(
				rtldService.Lookups.LookupsGetPlatforms(params),
			)

			internal.DisplayResponse(result)
		},
	}
	cmdLookupsGetPlatforms.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
	rootCmd.AddCommand(cmdLookupsGetPlatforms)

	cmdLookupsGetStatusCodes := &cobra.Command{
		Use:   "LookupsGetStatusCodes",
		Short: "LookupsGetStatusCodes",
		Long:  "LookupsGetStatusCodes",
		Run: func(cmd *cobra.Command, args []string) {
			var params *lookups.LookupsGetStatusCodesParams
			internal.ParseInputParams(cmd, params)

			rtldService := createClient()

			result := internal.CheckResult(
				rtldService.Lookups.LookupsGetStatusCodes(params),
			)

			internal.DisplayResponse(result)
		},
	}
	cmdLookupsGetStatusCodes.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
	rootCmd.AddCommand(cmdLookupsGetStatusCodes)

	return rootCmd
}

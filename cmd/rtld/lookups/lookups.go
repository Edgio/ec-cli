package lookups

import (
    "encoding/json"
    "fmt"
    "log"

    "github.com/EdgeCast/ec-sdk-go/edgecast/rtld/lookups"
    "github.com/spf13/cobra"
)

func parseInputParams[T any](c *cobra.Command, target *T) error {
    data := c.Flag("input-params-json").Value.String()

    if err := json.Unmarshal([]byte(data), target); err != nil {
        return fmt.Errorf("error parsing input params: %w", err)
    }

    return nil
}

func displayResponse(obj interface{}) {
    bytes, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		log.Fatal(err)
    }
    fmt.Printf("%s\n", string(bytes))
}


// Root returns the lookups root command
func Root() *cobra.Command {
    rootCmd := &cobra.Command{
        Use:   "lookups",
        Short: "lookups",
        Long: "lookups",
        Run: func(cmd *cobra.Command, args []string) {
            err := cmd.Help()
            if err != nil {
                log.Fatal(err)
            }
        },
    }
    
    
    cmdLookupsGetAwsRegions := &cobra.Command{
        Use:   "LookupsGetAwsRegions",
        Short: "LookupsGetAwsRegions",
        Long:  "LookupsGetAwsRegions",
        Run: func(cmd *cobra.Command, args []string) {
            var input lookups.LookupsGetAwsRegionsParams
            err := parseInputParams(cmd, &input)
            if err != nil {
                log.Fatal(err)
            }
            displayResponse(input)
        },
    }
    cmdLookupsGetAwsRegions.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
    rootCmd.AddCommand(cmdLookupsGetAwsRegions)
    
    cmdLookupsGetAzureAccessTypes := &cobra.Command{
        Use:   "LookupsGetAzureAccessTypes",
        Short: "LookupsGetAzureAccessTypes",
        Long:  "LookupsGetAzureAccessTypes",
        Run: func(cmd *cobra.Command, args []string) {
            var input lookups.LookupsGetAzureAccessTypesParams
            err := parseInputParams(cmd, &input)
            if err != nil {
                log.Fatal(err)
            }
            displayResponse(input)
        },
    }
    cmdLookupsGetAzureAccessTypes.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
    rootCmd.AddCommand(cmdLookupsGetAzureAccessTypes)
    
    cmdLookupsGetCustomItems := &cobra.Command{
        Use:   "LookupsGetCustomItems",
        Short: "LookupsGetCustomItems",
        Long:  "LookupsGetCustomItems",
        Run: func(cmd *cobra.Command, args []string) {
            var input lookups.LookupsGetCustomItemsParams
            err := parseInputParams(cmd, &input)
            if err != nil {
                log.Fatal(err)
            }
            displayResponse(input)
        },
    }
    cmdLookupsGetCustomItems.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
    rootCmd.AddCommand(cmdLookupsGetCustomItems)
    
    cmdLookupsGetDeliveryMethods := &cobra.Command{
        Use:   "LookupsGetDeliveryMethods",
        Short: "LookupsGetDeliveryMethods",
        Long:  "LookupsGetDeliveryMethods",
        Run: func(cmd *cobra.Command, args []string) {
            var input lookups.LookupsGetDeliveryMethodsParams
            err := parseInputParams(cmd, &input)
            if err != nil {
                log.Fatal(err)
            }
            displayResponse(input)
        },
    }
    cmdLookupsGetDeliveryMethods.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
    rootCmd.AddCommand(cmdLookupsGetDeliveryMethods)
    
    cmdLookupsGetDownsamplingRates := &cobra.Command{
        Use:   "LookupsGetDownsamplingRates",
        Short: "LookupsGetDownsamplingRates",
        Long:  "LookupsGetDownsamplingRates",
        Run: func(cmd *cobra.Command, args []string) {
            var input lookups.LookupsGetDownsamplingRatesParams
            err := parseInputParams(cmd, &input)
            if err != nil {
                log.Fatal(err)
            }
            displayResponse(input)
        },
    }
    cmdLookupsGetDownsamplingRates.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
    rootCmd.AddCommand(cmdLookupsGetDownsamplingRates)
    
    cmdLookupsGetFieldRl := &cobra.Command{
        Use:   "LookupsGetFieldRl",
        Short: "LookupsGetFieldRl",
        Long:  "LookupsGetFieldRl",
        Run: func(cmd *cobra.Command, args []string) {
            var input lookups.LookupsGetFieldRlParams
            err := parseInputParams(cmd, &input)
            if err != nil {
                log.Fatal(err)
            }
            displayResponse(input)
        },
    }
    cmdLookupsGetFieldRl.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
    rootCmd.AddCommand(cmdLookupsGetFieldRl)
    
    cmdLookupsGetFieldsCdn := &cobra.Command{
        Use:   "LookupsGetFieldsCdn",
        Short: "LookupsGetFieldsCdn",
        Long:  "LookupsGetFieldsCdn",
        Run: func(cmd *cobra.Command, args []string) {
            var input lookups.LookupsGetFieldsCdnParams
            err := parseInputParams(cmd, &input)
            if err != nil {
                log.Fatal(err)
            }
            displayResponse(input)
        },
    }
    cmdLookupsGetFieldsCdn.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
    rootCmd.AddCommand(cmdLookupsGetFieldsCdn)
    
    cmdLookupsGetFieldsWaf := &cobra.Command{
        Use:   "LookupsGetFieldsWaf",
        Short: "LookupsGetFieldsWaf",
        Long:  "LookupsGetFieldsWaf",
        Run: func(cmd *cobra.Command, args []string) {
            var input lookups.LookupsGetFieldsWafParams
            err := parseInputParams(cmd, &input)
            if err != nil {
                log.Fatal(err)
            }
            displayResponse(input)
        },
    }
    cmdLookupsGetFieldsWaf.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
    rootCmd.AddCommand(cmdLookupsGetFieldsWaf)
    
    cmdLookupsGetHTTPAuthenticationMethods := &cobra.Command{
        Use:   "LookupsGetHTTPAuthenticationMethods",
        Short: "LookupsGetHTTPAuthenticationMethods",
        Long:  "LookupsGetHTTPAuthenticationMethods",
        Run: func(cmd *cobra.Command, args []string) {
            var input lookups.LookupsGetHTTPAuthenticationMethodsParams
            err := parseInputParams(cmd, &input)
            if err != nil {
                log.Fatal(err)
            }
            displayResponse(input)
        },
    }
    cmdLookupsGetHTTPAuthenticationMethods.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
    rootCmd.AddCommand(cmdLookupsGetHTTPAuthenticationMethods)
    
    cmdLookupsGetLogFormats := &cobra.Command{
        Use:   "LookupsGetLogFormats",
        Short: "LookupsGetLogFormats",
        Long:  "LookupsGetLogFormats",
        Run: func(cmd *cobra.Command, args []string) {
            var input lookups.LookupsGetLogFormatsParams
            err := parseInputParams(cmd, &input)
            if err != nil {
                log.Fatal(err)
            }
            displayResponse(input)
        },
    }
    cmdLookupsGetLogFormats.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
    rootCmd.AddCommand(cmdLookupsGetLogFormats)
    
    cmdLookupsGetPlatforms := &cobra.Command{
        Use:   "LookupsGetPlatforms",
        Short: "LookupsGetPlatforms",
        Long:  "LookupsGetPlatforms",
        Run: func(cmd *cobra.Command, args []string) {
            var input lookups.LookupsGetPlatformsParams
            err := parseInputParams(cmd, &input)
            if err != nil {
                log.Fatal(err)
            }
            displayResponse(input)
        },
    }
    cmdLookupsGetPlatforms.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
    rootCmd.AddCommand(cmdLookupsGetPlatforms)
    
    cmdLookupsGetStatusCodes := &cobra.Command{
        Use:   "LookupsGetStatusCodes",
        Short: "LookupsGetStatusCodes",
        Long:  "LookupsGetStatusCodes",
        Run: func(cmd *cobra.Command, args []string) {
            var input lookups.LookupsGetStatusCodesParams
            err := parseInputParams(cmd, &input)
            if err != nil {
                log.Fatal(err)
            }
            displayResponse(input)
        },
    }
    cmdLookupsGetStatusCodes.Flags().StringP("input-params-json", "i", "{}", "input data formatted in JSON")
    rootCmd.AddCommand(cmdLookupsGetStatusCodes)
    
    return rootCmd
}

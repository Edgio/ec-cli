// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.
package internal

import (
	"fmt"
	"net/url"

	"github.com/EdgeCast/ec-sdk-go/edgecast"
	"github.com/EdgeCast/ec-sdk-go/edgecast/eclog"
	"github.com/spf13/viper"
)

// config stores all configurations of the application.
// The values are read by Viper from a config file or environment variables.
type Config struct {
	APIAddress       string `mapstructure:"API_ADDRESS"`
	APIAddressLegacy string `mapstructure:"API_ADDRESS_LEGACY"`
	APIToken         string `mapstructure:"API_TOKEN"`

	IDSAddress      string `mapstructure:"IDS_ADDRESS"`
	IDSClientID     string `mapstructure:"IDS_CLIENT_ID"`
	IDSClientSecret string `mapstructure:"IDS_CLIENT_SECRET"`
	IDSScope        string `mapstructure:"IDS_SCOPE"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() (*edgecast.SDKConfig, error) {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	var config *Config
	err = viper.Unmarshal(&config)

	if err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &edgecast.SDKConfig{
		BaseAPIURL:       stringToURL(config.APIAddress),
		BaseAPIURLLegacy: stringToURL(config.APIAddressLegacy),
		BaseIDSURL:       stringToURL(config.IDSAddress),
		Logger:           eclog.NewNullLogger(),
		APIToken:         config.APIToken,
		IDSCredentials: edgecast.IDSCredentials{
			ClientID:     config.IDSClientID,
			ClientSecret: config.IDSClientSecret,
			Scope:        config.IDSScope,
		},
		UserAgent: "ec-cli",
	}, nil
}

func stringToURL(s string) url.URL {
	u, _ := url.Parse(s)

	return *u
}

// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.
package util

import (
	"fmt"

	"github.com/spf13/viper"
)

// config stores all configurations of the application.
// The values are read by Viper from a config file or environment variables.
type Config struct {
	APIURLProd       string `mapstructure:"API_URL_PROD"`
	APIURLProdLegacy string `mapstructure:"API_URL_PROD_LEGACY"`
	IDSURLProd       string `mapstructure:"IDS_URL_PROD"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	var config *Config
	err = viper.Unmarshal(&config)

	if err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	return config, nil
}

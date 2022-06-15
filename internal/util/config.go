// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.
package util

import (
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
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}

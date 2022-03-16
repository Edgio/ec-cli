/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/edgecast/ec-cli/cmd"
)

func main() {

	/*
		config, err := util.LoadConfig(".")
		if err != nil {
			log.Fatal("cannot load config:", err)
		}
		fmt.Printf("%# v", pretty.Formatter(config.APIURLProd))
		fmt.Printf("%# v", pretty.Formatter(config.APIURLProdLegacy))
		fmt.Printf("%# v", pretty.Formatter(config.IDSURLProd))
	*/
	cmd.Execute()
}

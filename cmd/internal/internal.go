// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.
package internal

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func ParseInputParams[T any](c *cobra.Command, target *T) {
	data := c.Flag("input-params-json").Value.String()

	Check(
		json.Unmarshal([]byte(data), target),
	)
}

func DisplayResponse(obj any) {
	bytes := CheckResult(
		json.MarshalIndent(obj, "", "\t"),
	)
	fmt.Printf("%s\n", string(bytes)) //nolint: forbidigo
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckResult[T any](result T, err error) T { //nolint: ireturn
	if err != nil {
		log.Fatal(err)
	}

	return result
}

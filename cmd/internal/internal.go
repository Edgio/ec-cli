package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func ParseInputParams[T any](c *cobra.Command, target *T) {

	data := c.Flag("input-params-json").Value.String()
	Check(
		json.Unmarshal([]byte(data), target),
	)

	file := c.Flag("input-params-json-file").Value.String()

	if len(file) > 0 {
		if !IsFile(file) {
			fmt.Println("input error: params file does not exist.. ", file)

		}
		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
			return
		}
		Check(
			json.Unmarshal([]byte(content), target),
		)
		return
	}
}

func DisplayResponse(obj any) {
	bytes := CheckResult(
		json.MarshalIndent(obj, "", "\t"), //nolint: errchkjson
	)
	fmt.Printf("%s\n", string(bytes)) //nolint: forbidigo
}

func Check(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

func CheckResult[T any](result T, err error) T { //nolint: ireturn
	if err != nil {
		log.Fatal(err)
	}

	return result
}

// IsFile - check if the given path is file //TODO move this to cgk
func IsFile(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		return false
	}

	return !file.IsDir()
}

type FileHelper struct {
}

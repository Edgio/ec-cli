package util

import (
	"log"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	_, err := LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
}

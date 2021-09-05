package main

import (
	"log"
	"os"
	"testing"
)

func TestConfigInit(t *testing.T) {
	config := &CmdConfig{}
	config.constructor()

	if _, err := os.Stat(config.confFile); err != nil {
		log.Fatal(err)
	} else if os.IsNotExist(err) {
		log.Fatal(err)
		t.Fatalf("config file %v is not exist", config.confFile)
	}
}

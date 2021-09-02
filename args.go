package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func parseArgs() string {
	// define args
	definedArgs := [4]string{"init", "list", "add", "help"}
	// parse commandline args
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return "default"
	}
	if len(args) > 1 {
		fmt.Fprintln(os.Stderr, "args too much")
		os.Exit(1)
	}
	for _, v := range definedArgs {
		if args[0] == v {
			return args[0]
		}
	}
	fmt.Fprintln(os.Stderr, "args does not match(init|list|add|help)")
	os.Exit(1)
	return ""
}

func selectKey(config CmdConfig) string {
	confList, err := config.readConfig()
	if err != nil {
		os.Exit(1)
	}
	var keys []string
	for _, v := range confList {
		keys = append(keys, v["name"].(string))
	}
	prompt := promptui.Select{
		Label: "Select Key",
		Items: keys,
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Pronpt Failed %v\n", err)
	}
	return result
}

func help() {
	fmt.Println("GOTP Help")
}

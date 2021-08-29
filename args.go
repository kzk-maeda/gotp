package main

import (
	"flag"
	"fmt"
	"os"
)

func parseArgs() string {
	// define args
	definedArgs := [3]string{"init", "list", "add"}
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
	fmt.Fprintln(os.Stderr, "args does not match(init|list|add)")
	os.Exit(1)
	return ""
}

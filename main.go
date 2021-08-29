package main

import (
	"fmt"
	"strconv"

	"github.com/atotto/clipboard"
)

func main() {
	// parse args
	arg := parseArgs()
	fmt.Println(arg)
	config := &CmdConfig{}
	config.constructor()

	switch arg {
	case "init":
		config.init()
	case "add":
		config.addConfig()
	case "list":
		config.listConfig()
	case "help":
		help()
	case "default":
		// culc otp
		fmt.Println("GOTP")
		totp := execTOTP("ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEF", 30)
		fmt.Println(totp)
		clipboard.WriteAll(strconv.Itoa(totp))
	}
}

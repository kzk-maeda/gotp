package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/atotto/clipboard"
)

func main() {
	// parse args
	arg := parseArgs()
	// fmt.Println(arg)
	config := &CmdConfig{}
	config.constructor()

	switch arg {
	case "init":
		config.init()
	case "add":
		err := config.addConfig()
		if err != nil {
			os.Exit(1)
		}
	case "list":
		config.listConfig()
	case "help":
		help()
	case "default":
		// culc otp
		fmt.Println("GOTP")
		key := selectKey(*config)
		secret, _ := config.getSecretFromName(key)
		// fmt.Println(secret)
		totp := execTOTP(secret, 30)
		fmt.Println(totp)
		clipboard.WriteAll(strconv.Itoa(int(totp)))
	}
}

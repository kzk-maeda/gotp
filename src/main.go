package main

import (
	"fmt"
	"strconv"
	"./otp"
	"github.com/atotto/clipboard"
)

func main() {
	fmt.Println("GOTP")
	totp := otp.TOTP("ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEF", 30)
	fmt.Println(totp)
	clipboard.WriteAll(strconv.Itoa(totp))
}
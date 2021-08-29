package main

import (
	"fmt"
	"strconv"

	"github.com/atotto/clipboard"
)

func main() {
	fmt.Println("GOTP")
	totp := execTOTP("ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEF", 30)
	fmt.Println(totp)
	clipboard.WriteAll(strconv.Itoa(totp))
}

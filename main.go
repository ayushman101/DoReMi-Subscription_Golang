package main

import (
	"os"
)

const (
	START string = "START_SUBSCRIPTION"
	ADD   string = "ADD_SUBSCRIPTION"
	TOPUP string = "ADD_TOPUP"
	PRINT string = "PRINT_RENEWAL_DETAILS"
)

var user User

func main() {
	cliArgs := os.Args[1:]

	entrypoint(cliArgs)
}

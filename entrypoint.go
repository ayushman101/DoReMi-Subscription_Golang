package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func entrypoint(cliArgs []string) {
	filePath := cliArgs[0]
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening the input file")

		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		args := scanner.Text()
		argList := strings.Fields(args)

		switchCommand(argList)
	}
}

func switchCommand(argList []string) {

	switch argList[0] {
	case START:
		//start subscription handler
		StartSubscriptionHandler(argList)
	case ADD:
		//add subscription handler
		addSubscriptionHandler(argList)
	case PRINT:
		//print handler
		printDetailsHandler()
	case TOPUP:
		//add topup handler
		addTopUpHandler(argList)
	default:
		fmt.Println("Invalid command")

	}
}

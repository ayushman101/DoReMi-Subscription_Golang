package main

import (
	"testing"
	"time"
)

func TestEntrypoint(t *testing.T) {
	cliArgs := []string{
		"input2.txt",
	}

	entrypoint(cliArgs)

	cliArgs = []string{
		"inpu.txt",
	}

	entrypoint(cliArgs)
}

func TestStartSubscription(t *testing.T) {
	var argList []string = []string{
		"START_SUBSCRIPTION",
		"20-02-2022",
	}
	StartSubscriptionHandler(argList)
}

func TestAddSubscription(t *testing.T) {

	startDate, _ := time.Parse("2006-01-02", "2006-01-02")

	user = getNewDefaultUser(startDate)

	var argList []string = []string{
		"ADD_SUBSCRIPTION",
		"MUSIC",
		"FREE",
	}
	addSubscriptionHandler(argList)

}

func TestAddTopUp(t *testing.T) {

	user = getTestUser()
	argList := []string{
		"ADD_TOPUP",
		"FOUR_DEVICE",
		"3",
	}

	addTopUpHandler(argList)

}

func TestPrint(t *testing.T) {
	user = getTestUser()

	printDetailsHandler()

	user.printRenewalAmount()

}

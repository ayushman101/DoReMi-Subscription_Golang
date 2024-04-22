package main

import (
	"fmt"
	"strings"
	"time"
)

func StartSubscriptionHandler(argList []string) {

	date := strings.Split(argList[1], "-")

	//get the date and Validate it
	startDate, err := time.Parse("2006-01-02", date[2]+"-"+date[1]+"-"+date[0])

	if err != nil {
		fmt.Println("INVALID_DATE")
		return
	}

	user = getNewDefaultUser(startDate)

	// fmt.Println(user)
}

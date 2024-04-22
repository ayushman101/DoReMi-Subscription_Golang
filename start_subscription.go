package main

import (
	"fmt"
	"strings"
	"time"
)

func StartSubscriptionHandler(argList []string) {

	date := strings.Split(argList[1], "-")

	//get the date and Validate it
	startDate, err := time.Parse(time.DateOnly, date[2]+"-"+date[1]+"-"+date[0])

	if err != nil {
		fmt.Println("INVALID_DATE")
		return
	}

	user = User{
		StartDate: startDate,
		Music: Subscription{
			Category: MUSIC,
			Plan:     NONE,
		},
		Video: Subscription{
			Category: VIDEO,
			Plan:     NONE,
		},
		Podcast: Subscription{
			Category: PODCAST,
			Plan:     NONE,
		},
		TopUp: TopUp{
			Months:  1,
			Devices: 1,
		},
	}

	// fmt.Println(user)
}

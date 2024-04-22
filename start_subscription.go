package main

import (
	"fmt"
	"strings"
	"time"
)

func StartSubscriptionHandler(argList []string) {

	date := strings.Split(argList[1], "-")

	// day, _ := strconv.Atoi(date[0])
	// month, _ := strconv.Atoi(date[1])
	// year, _ := strconv.Atoi(date[2])

	// startDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	startDate, err := time.Parse(time.DateOnly, date[2]+"-"+date[1]+"-"+date[0])

	if err != nil {
		fmt.Println("INVALID_DATE")
		return
	}
	//TODO: validate the date

	// fmt.Println(startDate)

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

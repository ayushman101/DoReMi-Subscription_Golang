package main

import (
	"fmt"
	"strconv"
	"strings"
)

func StartSubscriptionHandler(argList []string) {

	date := strings.Split(argList[1], "-")
	fmt.Println(date)

	day, _ := strconv.Atoi(date[0])
	month, _ := strconv.Atoi(date[1])
	year, _ := strconv.Atoi(date[2])

	user = User{
		StartDate: Date{
			uint8(day),
			uint8(month),
			uint16(year),
		},
		Music: Subscription{
			Category: MUSIC,
			Plan:     NONE,
		},
		Video: Subscription{
			Category: VIDEO,
			Plan:     NONE,
		},
		Podcast: Subscription{
			Category: VIDEO,
			Plan:     NONE,
		},
		TopUp: TopUp{
			Months:  1,
			Devices: 1,
		},
	}

	fmt.Println(user)
}

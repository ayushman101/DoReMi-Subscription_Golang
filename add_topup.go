package main

import (
	"fmt"
	"strconv"
)

func addTopUpHandler(argList []string) {
	if user.TopUpPlans != DEFAULT {
		fmt.Println("ADD_TOPUP_FAILED DUPLICATE_TOPUP")
	}

	if user.Music.Plan == NONE && user.Video.Plan == NONE && user.Podcast.Plan == NONE {
		fmt.Println("ADD_TOPUP_FAILED SUBSCRIPTIONS_NOT_FOUND")
		return
	}

	months, _ := strconv.Atoi(argList[2])

	topup := NewTopUpPlan(argList[1], uint8(months))

	user.TopUp = topup

}

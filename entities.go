package main

import (
	"fmt"
	"time"
)

type User struct {
	StartDate time.Time
	Music     Subscription
	Video     Subscription
	Podcast   Subscription
	TopUp
}

func (user *User) AddSubscription(sub Subscription) {

	switch sub.Category {
	case MUSIC:
		validateAndAddSub(sub, &user.Music)
	case VIDEO:
		validateAndAddSub(sub, &user.Video)
	case PODCAST:
		validateAndAddSub(sub, &user.Podcast)
	}
}

func validateAndAddSub(sub Subscription, userSub *Subscription) {

	if userSub.Plan == NONE {
		*userSub = sub
	} else {
		fmt.Println("ADD_SUBSCRIPTION_FAILED DUPLICATE_CATEGORY")
	}

}

func getTestUser() User {
	startDate, _ := time.Parse("2006-01-02", "2006-01-02")

	return User{
		StartDate: startDate,
		Music:     getNewSubscription(MUSIC, FREE),
		Video:     getNewSubscription(VIDEO, PREMIUM),
		Podcast:   getNewSubscription(PODCAST, PERSONAL),
		TopUp:     getNewDefaultTopUp(TEN_DEVICE, 1, 4, 400),
	}
}

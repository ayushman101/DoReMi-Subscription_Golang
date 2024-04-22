package main

import (
	"fmt"
	"time"
)

type Category uint8

const (
	MUSIC Category = iota
	VIDEO
	PODCAST
)

func (c Category) String() string {
	switch c {
	case MUSIC:
		return "MUSIC"
	case VIDEO:
		return "VIDEO"
	case PODCAST:
		return "PODCAST"
	default:
		return "UNKNOWN"
	}
}

type Plan uint8

const (
	NONE Plan = iota
	FREE
	PERSONAL
	PREMIUM
)

func (p Plan) String() string {
	switch p {
	case FREE:
		return "FREE"
	case PERSONAL:
		return "PERSONAL"
	case PREMIUM:
		return "PREMIUM"
	default:
		return "NONE"
	}
}

type Subscription struct {
	Category
	Plan
	Months       uint8
	Amount       uint32
	Renewal_Date time.Time
}

//TODO :: COMPLETE THIS FUNCTION

func getNewSubscription(cat Category, pl Plan) Subscription {
	return Subscription{
		Category: cat,
		Plan:     pl,
	}
}

func getAmountForSubscription(p Plan, c Category) uint32 {
	if p == PERSONAL {
		if c == VIDEO {
			return uint32(200)
		}
		return uint32(100)
	}

	if p == PREMIUM {
		if c == MUSIC {
			return uint32(250)
		}
		if c == VIDEO {
			return uint32(500)
		}
		return uint32(300)
	}

	return uint32(0)
}

type TopUpPlans uint8

const (
	DEFAULT TopUpPlans = iota
	FOUR_DEVICE
	TEN_DEVICE
)

func (t TopUpPlans) String() string {
	switch t {
	case FOUR_DEVICE:
		return "FOUR_DEVICE"
	case TEN_DEVICE:
		return "TEN_DEVICE"
	default:
		return "DEFAULT"
	}
}

type TopUp struct {
	TopUpPlans
	Months  uint8 // default to 1
	Devices uint8 // can be 1, 4, or 8 and default to 1
	Amount  uint16
}

func getNewDefaultTopUp(tup TopUpPlans, months uint8, devices uint8, amount uint16) TopUp {
	return TopUp{
		TopUpPlans: tup,
		Months:     months,
		Devices:    devices,
		Amount:     amount,
	}
}

type User struct {
	StartDate time.Time
	Music     Subscription
	Video     Subscription
	Podcast   Subscription
	TopUp
}

func (u User) printRenewalAmount() {
	fmt.Printf("RENEWAL_AMOUNT ")
	amount := user.Music.Amount + user.Podcast.Amount + user.Video.Amount + uint32(user.TopUp.Amount)
	fmt.Println(amount)
}

func getNewDefaultUser(startDate time.Time) User {
	return User{
		StartDate: startDate,
		Music:     getNewSubscription(MUSIC, NONE),
		Video:     getNewSubscription(VIDEO, NONE),
		Podcast:   getNewSubscription(PODCAST, NONE),
		TopUp:     getNewDefaultTopUp(DEFAULT, 1, 1, 0),
	}
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

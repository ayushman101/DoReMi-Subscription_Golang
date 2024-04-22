package main

import "time"

func getNewSubscription(cat Category, pl Plan) Subscription {
	return Subscription{
		Category: cat,
		Plan:     pl,
	}
}

var getAmount map[string]uint32 = map[string]uint32{
	"MUSIC FREE":       0,
	"MUSIC PERSONAL":   100,
	"MUSIC PREMIUM":    250,
	"VIDEO FREE":       0,
	"VIDEO PERSONAL":   200,
	"VIDEO PREMIUM":    500,
	"PODCAST FREE":     0,
	"PODCAST PERSONAL": 100,
	"PODCAST PREMIUM":  300,
}

func getNewDefaultTopUp(tup TopUpPlans, months uint8, devices uint8, amount uint16) TopUp {
	return TopUp{
		TopUpPlans: tup,
		Months:     months,
		Devices:    devices,
		Amount:     amount,
	}
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

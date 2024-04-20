package main

import "fmt"

type Date struct {
	Day   uint8
	Month uint8
	Year  uint16
}

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

func getCategoryFromString(str string) Category {
	switch str {
	case "MUSIC":
		return MUSIC
	case "VIDEO":
		return VIDEO
	case "PODCAST":
		return PODCAST
	default:
		return MUSIC
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

func getPlanFromString(str string) Plan {
	switch str {
	case "FREE":
		return FREE
	case "PERSONAL":
		return PERSONAL
	case "PREMIUM":
		return PREMIUM
	default:
		return NONE
	}
}

type Subscription struct {
	Category
	Plan
	Months       uint8
	Amount       uint32
	Renewal_Date Date
}

//TODO :: COMPLETE THIS FUNCTION

// func (sub *Subscription) CalRenewableDate(startDate Date) {

// 	sub.Renewal_Date.Month = (startDate.Month + sub.Months) % 12

// 	if startDate.Month+sub.Months > 12 {
// 		sub.Renewal_Date.Year = startDate.Year + 1
// 	}

// }

func getMonthsFromPlan(p Plan) uint8 {
	switch p {
	case FREE:
		return uint8(1)
	case PERSONAL:
		return uint8(1)
	case PREMIUM:
		return uint8(3)
	default:
		return uint8(0)
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

type TopUp struct {
	Months  uint8 // default to 1
	Devices uint8 // can be 1, 4, or 8 and default to 1
}

type User struct {
	StartDate Date
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

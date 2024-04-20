package main

import "time"

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
	FREE Plan = iota
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
		return "UNKNOWN"
	}
}

type Subscription struct {
	Category
	Plan
	Months       uint8
	Amount       uint32
	Renewal_Date time.Time
}

type TopUp struct {
	Months  uint8 // default to 1
	Devices uint8 // can be 1, 4, or 8 and default to 1
}

type User struct {
	StartDate time.Time
	Music     Category
	Video     Category
	Podcast   Category
	TopUp
}

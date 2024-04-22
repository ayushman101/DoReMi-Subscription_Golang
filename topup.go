package main

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

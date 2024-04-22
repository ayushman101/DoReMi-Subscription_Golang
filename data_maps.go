package main

var categoryMap map[string]Category = map[string]Category{
	"MUSIC":   MUSIC,
	"VIDEO":   VIDEO,
	"PODCAST": PODCAST,
}

var planMap map[string]Plan = map[string]Plan{
	"FREE":     FREE,
	"PERSONAL": PERSONAL,
	"PREMIUM":  PREMIUM,
}

var planMonthsMap map[string]uint8 = map[string]uint8{
	"FREE":     1,
	"PERSONAL": 1,
	"PREMIUM":  3,
}

func NewTopUpPlan(plan string, months uint8) TopUp {
	if plan == FOUR_DEVICE.String() {
		return TopUp{
			TopUpPlans: FOUR_DEVICE,
			Devices:    4,
			Months:     months,
			Amount:     50 * uint16(months),
		}
	}

	return TopUp{
		TopUpPlans: TEN_DEVICE,
		Devices:    10,
		Months:     months,
		Amount:     100 * uint16(months),
	}

}

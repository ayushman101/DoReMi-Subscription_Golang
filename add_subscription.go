package main

import "fmt"

func addSubscriptionHandler(argList []string) {

	if user.StartDate.Year() == 1 {
		fmt.Println("ADD_SUBSCRIPTION_FAILED INVALID_DATE")
		return
	}

	sub := Subscription{

		Category: categoryMap[argList[1]],

		Plan: planMap[argList[2]],

		Months: planMonthsMap[planMap[argList[2]].String()],

		Amount: getAmountForSubscription(planMap[argList[2]], categoryMap[argList[1]]),

		Renewal_Date: user.StartDate.AddDate(0, int(planMonthsMap[planMap[argList[2]].String()]), -10),
	}

	user.AddSubscription(sub)

}

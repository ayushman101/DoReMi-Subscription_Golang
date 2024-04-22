package main

import "fmt"

func addSubscriptionHandler(argList []string) {

	if user.StartDate.Year() == 1 {
		fmt.Println("INVALID DATE")
		return
	}

	plan := getPlanFromString(argList[2])
	category := getCategoryFromString(argList[1])
	months := getMonthsFromPlan(plan)
	endDate := user.StartDate.AddDate(0, int(months), -10)

	sub := Subscription{
		Category:     category,
		Plan:         plan,
		Months:       months,
		Amount:       getAmountForSubscription(plan, category),
		Renewal_Date: endDate,
	}

	user.AddSubscription(sub)

	// fmt.Println(user)
}

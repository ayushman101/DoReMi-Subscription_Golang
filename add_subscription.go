package main

import "fmt"

func addSubscriptionHandler(argList []string) {

	plan := getPlanFromString(argList[2])
	category := getCategoryFromString(argList[1])

	sub := Subscription{
		Category: category,
		Plan:     plan,
		Months:   getMonthsFromPlan(plan),
		Amount:   getAmountForSubscription(plan, category),
	}

	user.AddSubscription(sub)

	fmt.Println(user)
}

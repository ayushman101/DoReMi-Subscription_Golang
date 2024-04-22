package main

import "fmt"

func printDetailsHandler() {
	if user.Music.Plan == NONE && user.Video.Plan == NONE && user.Podcast.Plan == NONE {
		fmt.Println("SUBSCRIPTIONS_NOT_FOUND")
		return
	}

	printSub(user.Music)
	printSub(user.Video)
	printSub(user.Podcast)

	user.printRenewalAmount()

}

func printSub(sub Subscription) {
	if sub.Plan != NONE {
		fmt.Printf("RENEWAL_REMINDER %s %d-%02d-%d\n", sub.Category.String(), sub.Renewal_Date.Day(), sub.Renewal_Date.Month(), sub.Renewal_Date.Year())
	}

}

func (u User) printRenewalAmount() {
	fmt.Printf("RENEWAL_AMOUNT ")
	amount := user.Music.Amount + user.Podcast.Amount + user.Video.Amount + uint32(user.TopUp.Amount)
	fmt.Println(amount)
}

package main

import "fmt"

func printDetailsHandler(argList []string) {
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
		fmt.Printf("RENEWAL_REMINDER %s %d-%d-%d\n", sub.Category.String(), user.Podcast.Renewal_Date.Day(), user.Podcast.Renewal_Date.Month(), user.Podcast.Renewal_Date.Year())
	}

}

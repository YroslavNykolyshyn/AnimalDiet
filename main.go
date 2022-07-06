package main

import "fmt"

const (
	feedAmountDog = 10
	feedAmountCat = 7
	feedAmountCow = 25
)

func main() {
	var totalFeedAmount int
	animals := []Animal{
		Dog{
			Name:       "Bob",
			Weight:     7,
			FeedAmount: feedAmountDog,
		},
		Cat{
			Name:       "Kuki",
			Weight:     7,
			FeedAmount: feedAmountCat,
		},
		Cow{
			Name:       "Star",
			Weight:     50,
			FeedAmount: feedAmountCow,
		},
	}
	for _, a := range animals {
		sum := a.getWeight() * a.getFeedAmount()
		totalFeedAmount += sum
		fmt.Printf("Name %s , Wieght %d , FeedAmount %d \n", a.getName(), a.getWeight(), a.getFeedAmount())
	}
	fmt.Printf("Total amount of feed: %d", totalFeedAmount)
}

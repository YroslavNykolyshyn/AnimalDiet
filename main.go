package main

import "fmt"

const (
	feedAmountDog = 10
	feedAmountCat = 7
	feedAmountCow = 25
)

func main() {
	var totalFeedAmount int
	animalFarm := []Animal{
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
		Dog{
			Name:       "Joy",
			Weight:     5,
			FeedAmount: feedAmountDog,
		},
		Cat{
			Name:       "Monika",
			Weight:     3,
			FeedAmount: feedAmountCat,
		},
		Cow{
			Name:       "Miya",
			Weight:     67,
			FeedAmount: feedAmountCow,
		},
	}
	for _, a := range animalFarm {
		feedAmount := a.getWeight() * a.getFeedAmount()
		totalFeedAmount += feedAmount
		fmt.Printf(
			"Name %s , Weight %d , FeedAmount %d \n Next animal->\n",
			a.getName(),
			a.getWeight(),
			feedAmount,
		)
	}
	fmt.Printf("Total amount of feed: %d", totalFeedAmount)
}

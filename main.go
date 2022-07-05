package main

import "fmt"

type avgWeight interface {
	getAvgWeight() int
}

func FeedAmount(a avgWeight, weight int) int {
	fa := a.getAvgWeight() * weight
	return fa
}

func totalFeedAmount(feedAmCow, feedAmCat, feedAmDog int) int {
	feedAm := feedAmDog + feedAmCow + feedAmCat
	return feedAm
}

func main() {
	dog := Dog{
		Name:       "Bob",
		Weight:     5,
		FeedAmount: 10,
	}

	var awDog avgWeight
	awDog = dog
	feedAmDog := FeedAmount(awDog, dog.Weight)
	fmt.Printf("Dog name: %s"+
		" Amount of food to feed a dog: %d \n",
		dog.Name, feedAmDog)

	cat := Cat{
		Name:       "Kuki",
		Weight:     3,
		FeedAmount: 7,
	}
	var awCat avgWeight
	awCat = cat
	feedAmCat := FeedAmount(awCat, cat.Weight)
	fmt.Printf("Cat name: %s"+
		" Amount of food to feed a cat: %d \n",
		cat.Name, feedAmCat)

	cow := Cow{
		Name:       "Star",
		Weight:     100,
		FeedAmount: 25,
	}
	var awCow avgWeight
	awCow = cow
	feedAmCow := FeedAmount(awCow, cow.Weight)
	fmt.Printf("Cow name: %s"+
		" Amount of food to feed a cow: %d \n",
		cow.Name, feedAmCow)

	fmt.Printf("Total amount of feed: %d", totalFeedAmount(feedAmCow, feedAmCat, feedAmDog))

}

//Привіт

package main

type Animal interface {
	AnimalName
	AnimalWeight
	AnimalFeedAmount
}
type AnimalName interface {
	getName() string
}

type AnimalWeight interface {
	getWeight() int
}

type AnimalFeedAmount interface {
	getFeedAmount() int
}

type Dog struct {
	Name       string
	Weight     int
	FeedAmount int
}

func (d Dog) getName() string {
	return d.Name
}

func (d Dog) getWeight() int {
	return d.Weight
}

func (d Dog) getFeedAmount() int {
	return d.FeedAmount
}

type Cat struct {
	Name       string
	Weight     int
	FeedAmount int
}

func (c Cat) getName() string {
	return c.Name
}

func (c Cat) getWeight() int {
	return c.Weight
}

func (c Cat) getFeedAmount() int {
	return c.FeedAmount
}

type Cow struct {
	Name       string
	Weight     int
	FeedAmount int
}

func (cow Cow) getName() string {
	return cow.Name
}

func (cow Cow) getWeight() int {
	return cow.Weight
}

func (cow Cow) getFeedAmount() int {
	return cow.FeedAmount
}

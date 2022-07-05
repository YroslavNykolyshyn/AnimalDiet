package main

type Dog struct {
	Name       string
	Weight     int
	FeedAmount int
}
type Cat struct {
	Name       string
	Weight     int
	FeedAmount int
}
type Cow struct {
	Name       string
	Weight     int
	FeedAmount int
}

func (d Dog) getAvgWeight() int {
	return d.FeedAmount
}

func (c Cat) getAvgWeight() int {
	return c.FeedAmount
}

func (co Cow) getAvgWeight() int {
	return co.FeedAmount
}

//Привіт

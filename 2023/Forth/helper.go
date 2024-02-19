package main

import (
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printSeperator() {
	fmt.Println("--------------------------------------------------------------")
}

func (c *Cards) display() {
	for _, value := range *c {
		value.display()
	}
}

func (c *Card) display() {
	fmt.Printf("ID: %d - > card numbers %v; winning numbers %v\n", c.ID, c.NumberYouHave, c.WinningNumbers)
}

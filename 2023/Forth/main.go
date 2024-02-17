package main

import (
	"bufio"
	"fmt"
	"os"
)

type Card struct {
	ID             int
	WinningNumbers []int
	NumberYouHave  []int
}

func main() {
	file, err := os.Open("input/sample_input1.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	cards := []Card{}
	for scanner.Scan() {
		line := scanner.Text()
		current_card := parse_card(line)
		cards = append(cards, current_card)
	}

	fmt.Printf("%v", cards)
}

func parse_card(line string) Card {
	// Need to get the ID after the word "Card"
	// Need to get the string after ':' until I hit a '|' then break it on space
	// Need to get the rest of the string then break it on space
	// Create the card and send it back
	return Card{}
}

func (c *Card) get_card_points() int {
	return 0
}

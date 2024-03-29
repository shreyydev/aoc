package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	ID             int
	WinningNumbers []int
	NumberYouHave  []int
	Count          int
}

type Cards []Card

func main() {
	file, err := os.Open("input/input2.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	var cards Cards
	for scanner.Scan() {
		line := scanner.Text()
		current_card := parse_card(line)
		cards = append(cards, current_card)
	}
	// cards.display()
	cards.part1()
	cards.part2()
}

func (c *Cards) part1() {
	sum := 0
	for _, card := range *c {
		sum += card.get_card_points()
	}
	fmt.Printf("Sum: %v\n", sum)
}

func (c *Cards) part2() {
	for i := 0; i < len(*c); i++ {
		card := (*c)[i]
		cards_won := card.get_matching_winnning_numbers()
		for j := i + 1; j < i+1+cards_won; j++ {
			(*c)[j].Count += card.Count
		}
	}
	sum := 0
	for _, card := range *c {
		sum += card.Count
	}
	fmt.Printf("Total cards: %d", sum)
}

func (c *Card) get_matching_winnning_numbers() int {
	matching_nums := 0
	for _, card_num := range c.NumberYouHave {
		if slices.Contains(c.WinningNumbers, card_num) {
			matching_nums++
		}
	}
	return matching_nums
}

func parse_card(line string) Card {
	// Need to get the ID after the word "Card"
	fields := strings.FieldsFunc(line, func(r rune) bool {
		if r == ' ' || r == ':' {
			return true
		}
		return false
	})
	ID_Int, _ := strconv.Atoi(fields[1])

	// Need to get the string after ':' until I hit a '|' then break it on space
	var winning_numbers []int
	var i int
	for i = 2; fields[i] != "|"; i++ {
		num, _ := strconv.Atoi(fields[i])
		winning_numbers = append(winning_numbers, num)
	}
	// Need to get the rest of the string then break it on space
	var numbers_you_have []int
	for j := i + 1; j < len(fields); j++ {
		num, _ := strconv.Atoi(fields[j])
		numbers_you_have = append(numbers_you_have, num)
	}

	return Card{
		ID:             ID_Int,
		WinningNumbers: winning_numbers,
		NumberYouHave:  numbers_you_have,
		Count:          1,
	}
}

func (c *Card) get_card_points() int {
	points := 0
	for _, card_num := range c.NumberYouHave {
		if slices.Contains(c.WinningNumbers, card_num) {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}
	return points
}

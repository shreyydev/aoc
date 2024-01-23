package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"unicode"
)

type Point struct {
	line     int
	position int
}

type CharMap map[Point]rune

type NumberMap map[int][]Point
type FinalNumberMap map[int]int

// Part1 answer: 537732

func main() {
	file, err := os.Open("input/input2.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	line_num := 0
	char_map := CharMap{}
	for scanner.Scan() {
		line := scanner.Text()
		char_map.addCharsFromLine(line, line_num)
		line_num++
	}

	// part1(char_map)
	part2(char_map)

}

func part2(char_map CharMap) {
	sum := 0
	number_map := createNumberMap(char_map)

	sorted_char_maps_keys := char_map.get_sorted_keys()

	symbols_found := make(map[string]int)

	// Iterate over each of the coordinated in a sorted manner (0,0; 0,1; 0,2; ...)
	for _, point := range sorted_char_maps_keys {
		current_char := char_map[point] // the current char
		// check if the current char is a symbol
		if current_char == '*' {
			symbols_found[string(current_char)] = symbols_found[string(current_char)] + 1
			// Get the surrounding points that are around the current coordinate
			adj_points := getAdjacentPoints(point)
			// for each of the surrounding point
			nums_encountered := make(map[int]bool)
			for adj_point := range adj_points {
				// for each number we parsed from the input
			num_loop:
				for num, points_arr := range number_map {
					// check if the surrouding point is contained in the current number array
					if _, ok := nums_encountered[num]; slices.Contains(points_arr, adj_point) && !ok {
						nums_encountered[num] = true
						break num_loop
					}
				}
			}
			if len(nums_encountered) == 2 {
				temp := 1
				for num := range nums_encountered {
					temp *= num
				}
				sum += temp
			}
		}
	}

	fmt.Printf("Part 2 => Sum: %d\n", sum)
	// final_number_map.display()
	// fmt.Printf("Symbols found: %v\n", symbols_found)
	// fmt.Printf("Char Map: %v\n", char_map)
	// fmt.Printf("Sorted Char Map: %v\n", sorted_char_maps_keys)
	// fmt.Printf("Number Map: %v\n", number_map)

}

func part1(char_map CharMap) {
	sum := 0
	number_map := createNumberMap(char_map)

	final_number_map := createFinalNumberMap(number_map)
	sorted_char_maps_keys := char_map.get_sorted_keys()

	symbols_found := make(map[string]int)

	// Iterate over each of the coordinated in a sorted manner (0,0; 0,1; 0,2; ...)
	for _, point := range sorted_char_maps_keys {
		current_char := char_map[point] // the current char
		// check if the current char is a symbol
		if current_char != '.' && !unicode.IsNumber(current_char) {
			symbols_found[string(current_char)] = symbols_found[string(current_char)] + 1
			// Get the surrounding points that are around the current coordinate
			adj_points := getAdjacentPoints(point)
			// for each of the surrounding point
			nums_encountered := make(map[int]bool)
			for adj_point := range adj_points {
				// for each number we parsed from the input
			num_loop:
				for num, points_arr := range number_map {
					// check if the surrouding point is contained in the current number array
					if _, ok := nums_encountered[num]; slices.Contains(points_arr, adj_point) && !ok {
						nums_encountered[num] = true
						final_number_map[num] += 1
						break num_loop
					}
				}
			}
		}
	}

	for num, count := range final_number_map {
		sum += (num * count)
	}

	fmt.Printf("Part 1 => Sum: %d\n", sum)
	// final_number_map.display()
	// fmt.Printf("Symbols found: %v\n", symbols_found)
	// fmt.Printf("Char Map: %v\n", char_map)
	// fmt.Printf("Sorted Char Map: %v\n", sorted_char_maps_keys)
	// fmt.Printf("Number Map: %v\n", number_map)
}

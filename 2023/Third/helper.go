package main

import (
	"fmt"
	"sort"
	"strconv"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func printSeperator() {
	fmt.Println("--------------------------------------------------------------")
}

func createFinalNumberMap(number_map NumberMap) FinalNumberMap {
	final_number_map := FinalNumberMap{}
	for k := range number_map {
		final_number_map[k] = 0
	}
	return final_number_map
}

func getAdjacentPoints(symbol_point Point) map[Point]int {
	points := map[Point]int{
		{
			line:     max(0, symbol_point.line-1),
			position: max(0, symbol_point.position-1),
		}: 1,
		{
			line:     max(0, symbol_point.line-1),
			position: max(0, symbol_point.position),
		}: 1,
		{
			line:     max(0, symbol_point.line-1),
			position: max(0, symbol_point.position+1),
		}: 1,
		{
			line:     max(0, symbol_point.line),
			position: max(0, symbol_point.position-1),
		}: 1,
		{
			line:     max(0, symbol_point.line),
			position: max(0, symbol_point.position+1),
		}: 1,
		{
			line:     max(0, symbol_point.line+1),
			position: max(0, symbol_point.position-1),
		}: 1,
		{
			line:     max(0, symbol_point.line+1),
			position: max(0, symbol_point.position),
		}: 1,
		{
			line:     max(0, symbol_point.line+1),
			position: max(0, symbol_point.position+1),
		}: 1,
	}
	delete(points, symbol_point)
	return points
}

// * Point
func (p *Point) display() {
	fmt.Printf("(line: %d, postion: %d)", p.line, p.position)
}

// * Final Number Map

func (f *FinalNumberMap) display() {
	fv := *f
	sorted_keys := f.get_sorted_keys()
	for _, number := range sorted_keys {
		val := fv[number]
		fmt.Printf("%d: %v\n", number, val)
	}
}

func (f *FinalNumberMap) get_sorted_keys() []int {
	mv := *f
	keys := make([]int, 0, len(mv))
	for k := range mv {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

// * CharMap
func (m *CharMap) get_sorted_keys() []Point {
	mv := *m
	keys := make([]Point, 0, len(mv))
	for k := range mv {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		if keys[i].line != keys[j].line {
			return keys[i].line < keys[j].line
		}
		return keys[i].position < keys[j].position
	})
	return keys
}

func createNumberMap(char_map CharMap) NumberMap {
	number_map := NumberMap{}
	accumalator := 0
	points_accumalator := []Point{}

	keys := char_map.get_sorted_keys()

	for _, point := range keys {
		current_char := char_map[point]
		if unicode.IsDigit(current_char) {
			num, _ := strconv.Atoi(string(current_char))
			accumalator = accumalator*10 + num
			points_accumalator = append(points_accumalator, point)
		} else {
			if accumalator != 0 {
				val, ok := number_map[accumalator]
				if ok {
					val = append(val, points_accumalator...)
				} else {
					val = points_accumalator
				}
				number_map[accumalator] = val
				points_accumalator = []Point{}
				accumalator = 0
			}
		}
	}
	return number_map
}

func (m *CharMap) addCharsFromLine(line string, y_value int) {
	for x_value, char := range line {
		(*m)[Point{position: x_value, line: y_value}] = char
	}
}

func (m *CharMap) display() {
	mv := *m
	keys := mv.get_sorted_keys()

	for _, k := range keys {
		k.display()
		fmt.Printf(": %c\n", mv[k])
	}
}

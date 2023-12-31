package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	line     int
	position int
}

type CharMap map[Point]string

func main() {
	fmt.Println("Third")
	file, err := os.Open("input/sample_input1.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	line_num := 1
	sum := 0
	char_map := CharMap{}
	for scanner.Scan() {
		line := scanner.Text()

		char_map.addCharsFromLine(line, line_num-1)

		line_num++
	}
	char_map.display()
	fmt.Printf("Sum: %d\n", sum)
}

// TODO A function that gives me the points that I need to check to find an adjacent point

func (m *CharMap) addCharsFromLine(line string, y_value int) {
	for x_value, char := range line {
		(*m)[Point{position: x_value, line: y_value}] = string(char)
	}
}

func (m *CharMap) display() {
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

	for _, k := range keys {
		k.display()
		fmt.Printf(": %s\n", mv[k])
	}
}

func (p *Point) display() {
	fmt.Printf("(line: %d, postion: %d)", p.line+1, p.position+1)
}

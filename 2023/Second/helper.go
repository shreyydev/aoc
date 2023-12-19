package main

import (
	"fmt"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getNumUntil(line string, startPosition int, stopChar string) int {
	num := 0
	for i := startPosition; string(line[i]) != stopChar; i++ {
		id, err := strconv.Atoi(string(line[i]))
		check(err)
		num = num*10 + id
	}
	return num
}

func printSeperator() {
	fmt.Println("--------------------------------------------------------------")
}

func (g Game) display() {
	fmt.Printf("Game id: %d\n", g.ID)
	for setNum, setInfo := range g.Set {
		fmt.Printf("Set %d: ", setNum+1)
		setInfo.display()
	}
	printSeperator()
}

func (s SetInfo) display() {
	fmt.Printf("{Red: %d, Green: %d, Blue: %d}\n", s.Red, s.Green, s.Blue)
}

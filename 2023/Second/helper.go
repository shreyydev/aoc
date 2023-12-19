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

func printSeperator() {
	fmt.Println("--------------------------------------------------------------")
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

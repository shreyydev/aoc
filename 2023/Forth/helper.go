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

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input/sample_input1.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

}

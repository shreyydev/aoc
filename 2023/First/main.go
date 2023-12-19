package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var alphaNumbers = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9, "ten": 10}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("real_input2.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	sumWithWords := 0
	for scanner.Scan() {
		// sum += extractNumNumerical(scanner.Text())
		sumWithWords += extractNumWordsCapable(scanner.Text())
	}
	fmt.Printf("Total sum: %d\n", sum)
	fmt.Printf("Total sum with words: %d\n", sumWithWords)
}

func extractNumWordsCapable(line string) int {
	left := firstNumFromLeft(line)
	right := firstNumFromRight(line)
	return (left * 10) + right
}

func firstNumFromLeft(line string) int {
	lineLength := len(line)
	for i, char := range line {
		if unicode.IsDigit(char) {
			num, _ := strconv.Atoi(string(char))
			return num
		}
		var sb strings.Builder
		sb.WriteString(string(char))
	inner:
		for j := i + 1; j < lineLength; j++ {
			sb.WriteString(string(line[j]))
			word := sb.String()
			if len(word) <= 5 {
				if containsKey(alphaNumbers, word) {
					return alphaNumbers[word]
				}
			} else {
				break inner
			}
		}
	}
	return -1
}

func firstNumFromRight(line string) int {
	line = Reverse(line)
	lineLength := len(line)
	for i, char := range line {
		if unicode.IsDigit(char) {
			num, _ := strconv.Atoi(string(char))
			return num
		}
		var sb strings.Builder
		sb.WriteString(string(char))
	inner:
		for j := i + 1; j < lineLength; j++ {
			var sb2 strings.Builder
			sb2.WriteString(string(line[j]))
			sb2.WriteString(sb.String())
			sb.Reset()
			sb.WriteString(sb2.String())
			word := sb.String()
			if len(word) <= 5 {
				if containsKey(alphaNumbers, word) {
					return alphaNumbers[word]
				}
			} else {
				break inner
			}
		}
	}
	return -1
}

func containsKey[M ~map[K]V, K comparable, V any](m M, k K) bool {
	_, ok := m[k]
	return ok
}

// * Functions used in problem 1
// func extractNumNumerical(line string) int {
// 	left := firstNumFromLeft(line) //! function has been changed in the problem 2
// 	right := firstNumFromLeft(Reverse(line))
// 	return (left * 10) + right
// }

// func firstNumFromLeft(line string) int {
// 	for _, char := range line {
// 		if unicode.IsDigit(char) {
// 			num, _ := strconv.Atoi(string(char))
// 			return num
// 		}
// 	}
// 	return -1
// }

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	ID  int
	Set map[int]SetInfo
}

type SetInfo struct {
	Red   int
	Green int
	Blue  int
}

const gameIdFieldIndex = 1
const set1FieldIndex = 2

// const maxRedCubes = 12
// const maxGreenCubes = 13
// const maxBlueCubes = 14

func main() {
	file, err := os.Open("input2.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// sum := 0
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	game := parseGame(line)
	// 	if isGamePossible(game) {
	// 		sum += game.ID
	// 	}
	// }
	// fmt.Printf("Sum of Game IDs: %d\n", sum)

	sumMinRequiredSet := 0
	for scanner.Scan() {
		line := scanner.Text()
		game := parseGame(line)
		minRequiredSet := minRequiredSetToPlayGame(game)
		sumMinRequiredSet += minRequiredSet.Red * minRequiredSet.Blue * minRequiredSet.Green
	}
	fmt.Printf("Sum of Game IDs: %d\n", sumMinRequiredSet)

}

func parseGame(line string) Game {
	stringWords := strings.Fields(line)

	id := getNumUntil(stringWords[gameIdFieldIndex], 0, ":")
	game := Game{ID: id, Set: make(map[int]SetInfo)}

	currentSetIndex := set1FieldIndex
	for i := 0; currentSetIndex < len(stringWords)-1; i++ {
		game.Set[i], currentSetIndex = getSetInfo(stringWords, currentSetIndex)
	}

	return game
}

func getSetInfo(stringWords []string, startIndex int) (SetInfo, int) {
	setInfo := SetInfo{}
	lastIndex := startIndex
	for i := startIndex; i < len(stringWords); i++ {
		if count, err := strconv.Atoi(stringWords[i]); err == nil {
			colorIndex := i + 1
			if strings.Contains(stringWords[colorIndex], "red") {
				setInfo.Red = count
			} else if strings.Contains(stringWords[colorIndex], "blue") {
				setInfo.Blue = count
			} else {
				setInfo.Green = count
			}
			i++
			lastIndex = i
			if strings.Contains(stringWords[colorIndex], ";") {
				break
			}
		}
	}
	return setInfo, lastIndex
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

// func isGamePossible(game Game) bool {
// 	for _, set := range game.Set {
// 		if set.Red > maxRedCubes || set.Blue > maxBlueCubes || set.Green > maxGreenCubes {
// 			return false
// 		}
// 	}
// 	return true
// }

func minRequiredSetToPlayGame(game Game) SetInfo {
	minSetRequired := SetInfo{Red: 0, Blue: 0, Green: 0}
	for _, set := range game.Set {
		if set.Red > minSetRequired.Red {
			minSetRequired.Red = set.Red
		}
		if set.Blue > minSetRequired.Blue {
			minSetRequired.Blue = set.Blue
		}
		if set.Green > minSetRequired.Green {
			minSetRequired.Green = set.Green
		}
	}
	return minSetRequired
}

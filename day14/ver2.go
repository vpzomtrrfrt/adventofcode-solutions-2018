package main

import "fmt"
import "log"
import "os"
import "io/ioutil"
import "strings"

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	str := strings.TrimSpace(string(bytes))
	chars := []rune(str)
	searchPattern := make([]int, len(chars))
	for idx, char := range chars {
		searchPattern[idx] = int(char - 48)
	}

	scoreboard := []int{3, 7}

	currentA := 0
	currentB := 1

	for {
		scoreA := scoreboard[currentA]
		scoreB := scoreboard[currentB]

		sum := scoreA + scoreB

		var newScores []int
		if sum >= 10 {
			newScores = []int{sum / 10, sum % 10}
		} else {
			newScores = []int{sum}
		}

		scoreboard = append(scoreboard, newScores...)

		currentA += scoreA + 1
		currentB += scoreB + 1

		for currentA >= len(scoreboard) {
			currentA -= len(scoreboard)
		}
		for currentB >= len(scoreboard) {
			currentB -= len(scoreboard)
		}

		for o := max(0, len(scoreboard) - len(searchPattern) - 2); o < len(scoreboard) - len(searchPattern); o++ {
			found := true
			for i, score := range searchPattern {
				idx := i + o
				if scoreboard[idx] != score {
					found = false
					break
				}
			}
			if found {
				fmt.Println(o)
				return
			}
		}
	}
}

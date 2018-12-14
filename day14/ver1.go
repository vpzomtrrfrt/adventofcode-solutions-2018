package main

import "fmt"
import "log"
import "os"
import "io/ioutil"
import "strings"
import "strconv"

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	str := strings.TrimSpace(string(bytes))
	input, err := strconv.Atoi(str)
	if err != nil { log.Fatal(err) }

	scoreboard := []int{3, 7}

	currentA := 0
	currentB := 1

	for len(scoreboard) < input + 10 {
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
	}
	for _, score := range scoreboard[input:(input + 10)] {
		fmt.Print(score)
	}
	fmt.Println()
}

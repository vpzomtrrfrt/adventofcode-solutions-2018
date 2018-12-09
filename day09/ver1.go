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

	input := string(bytes)

	i1 := strings.Index(input, " ")
	players_str := input[:i1]
	players, err := strconv.Atoi(players_str)
	if err != nil { log.Fatal(err) }

	i2 := strings.Index(input, "worth") + 6
	i3 := strings.Index(input[i2:], " ") + i2
	marbles_str := input[i2:i3]
	marbles, err := strconv.Atoi(marbles_str)
	if err != nil { log.Fatal(err) }

	circle := []int{0}

	i := 1

	current := 0

	scores := make(map[int]int)

	for i <= marbles {
		if i % 23 == 0 {
			currentPlayer := i % players
			points := i

			current -= 7
			if current < 0 {
				current += len(circle)
			}
			points += circle[current]

			circle = append(circle[:current], circle[(current + 1):]...)

			scores[currentPlayer] += points
		} else {
			current += 2
			for current > len(circle) {
				current -= len(circle)
			}

			newCircle := make([]int, 0)
			newCircle = append(newCircle, circle[:current]...)
			newCircle = append(newCircle, i)
			newCircle = append(newCircle, circle[current:]...)
			circle = newCircle
		}

		i++
	}

	highest := 0

	for _, score := range scores {
		if score > highest {
			highest = score
		}
	}

	fmt.Println(highest)
}

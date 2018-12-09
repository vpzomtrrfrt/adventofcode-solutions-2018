package main

import "fmt"
import "log"
import "os"
import "io/ioutil"
import "strings"
import "strconv"
import "container/list"

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

	marbles *= 100

	circle := list.New()
	circle.PushBack(0)

	i := 1

	current := circle.Front()

	scores := make(map[int]int)

	for i <= marbles {
		if i % 23 == 0 {
			currentPlayer := i % players
			points := i

			for j := 0; j < 7; j++ {
				current = current.Prev()
				if current == nil {
					current = circle.Back()
				}
			}
			points += current.Value.(int)

			taken := current
			current = current.Next()

			circle.Remove(taken)

			scores[currentPlayer] += points
		} else {
			for j := 0; j < 2; j++ {
				current = current.Next()
				if current == nil {
					current = circle.Front()
				}
			}

			circle.InsertBefore(i, current)
			current = current.Prev()
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

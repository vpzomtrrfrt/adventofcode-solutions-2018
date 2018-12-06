package main

import "fmt"
import "log"
import "os"
import "io/ioutil"
import "unicode"
import "strings"

func process(remaining []rune) []rune {
	var i = 0

	for i < len(remaining) - 1 {
		current := remaining[i]
		next := remaining[i + 1]

		if current != next && unicode.ToLower(current) == unicode.ToLower(next) {
			remaining = append(remaining[:i], remaining[(i + 2):]...)

			if i > 0 {
				i--
			}
		} else {
			i++
		}
	}

	return remaining
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var start = []rune(strings.TrimSpace(string(bytes)))
	letters := "abcdefghijklmnopqrstuvwxyz"

	var lowScore = len(start)

	for _, letter := range letters {
		upper := unicode.ToUpper(letter)

		// based on https://stackoverflow.com/a/28714120/2533397
		var current = make([]rune, len(start))
		w := 0
		for _, char := range start {
			if char != letter && char != upper {
				current[w] = char
				w++
			}
		}
		current = current[:w]
		var remains = process(current)
		var score = len(remains)

		if score < lowScore {
			lowScore = score
		}
	}

	fmt.Println(lowScore)
}

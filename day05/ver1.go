package main

import "fmt"
import "log"
import "os"
import "io/ioutil"
import "unicode"
import "strings"

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var remaining = []rune(strings.TrimSpace(string(bytes)))

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

	fmt.Println(len(remaining))
}

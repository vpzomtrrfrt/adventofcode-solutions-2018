package main

import "fmt"
import "log"
import "os"
import "io/ioutil"
import "strings"

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	spl := strings.Split(string(bytes), "\n")

	count2 := 0
	count3 := 0

	for _, line := range spl {
		counts := make(map[rune]int)
		for _, char := range line {
			if _, found := counts[char]; found {
				counts[char] += 1
			} else {
				counts[char] = 1
			}
		}

		var found2 = false
		var found3 = false

		for _, count := range counts {
			if count == 2 {
				found2 = true
			}
			if count == 3 {
				found3 = true
			}
		}

		if found2 {
			count2 += 1
		}
		if found3 {
			count3 += 1
		}
	}

	result := count2 * count3
	fmt.Println(result)
}

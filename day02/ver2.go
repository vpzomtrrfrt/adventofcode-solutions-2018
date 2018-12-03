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

	for prev_row := 0; prev_row < len(spl); prev_row++ {
		prev := spl[prev_row]
		prev_ary := []rune(prev)

		for curr_row := prev_row + 1; curr_row < len(spl); curr_row++ {
			curr := spl[curr_row]
			curr_ary := []rune(curr)

			var diffidx = -1

			for i := 0; i < len(curr_ary); i++ {
				curr_val := curr_ary[i]
				prev_val := prev_ary[i]

				if curr_val != prev_val {
					if diffidx == -1 {
						diffidx = i
					} else {
						diffidx = -2
					}
				}
			}

			if diffidx >= 0 {
				fmt.Println(string(curr_ary[0:diffidx]) + string(curr_ary[(diffidx + 1):]))
				return
			}
		}
	}
}

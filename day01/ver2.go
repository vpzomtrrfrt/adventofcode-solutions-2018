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

	spl := strings.Split(string(bytes), "\n")
	total := 0

	seen := make(map[int]int)
	seen[0] = 0

	for {
		for i := 0; i < len(spl); i++ {
			line := spl[i]
			if line == "" {
				continue
			}
			value, err := strconv.Atoi(line)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			} else {
				total += value
				if _, found := seen[total]; found {
					fmt.Println(total)
					return
				}
				seen[total] = 0
			}
		}
	}
}

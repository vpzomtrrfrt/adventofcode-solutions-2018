package main

import "fmt"
import "log"
import "os"
import "io/ioutil"
import "strings"
import "sort"

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	spl := strings.Split(string(bytes), "\n")

	requirements := make(map[string]*[]string)

	for _, line := range spl {
		if line == "" {
			continue;
		}

		requirement := line[5:6]
		id := line[36:37]

		if list, found := requirements[id]; found {
			(*list) = append(*list, requirement)
		} else {
			newList := []string{requirement}
			requirements[id] = &newList
		}

		if _, found := requirements[requirement]; !found {
			newList := make([]string, 0)
			requirements[requirement] = &newList
		}
	}

	steps := make([]string, 0, len(requirements))
	for key := range requirements {
		steps = append(steps, key)
	}
	sort.Slice(steps, func(i, j int) bool { return steps[i] < steps[j] })

	completed := make(map[string]bool)

	for {
		var anyWaiting = false
		for _, step := range steps {
			if _, found := completed[step]; found {
				continue
			}

			anyWaiting = true

			reqs := requirements[step]
			var canStart = true
			for _, req := range (*reqs) {
				if _, found := completed[req]; !found {
					canStart = false
					break
				}
			}

			if canStart {
				fmt.Print(step)
				completed[step] = true
				break
			}
		}
		if !anyWaiting {
			break
		}
	}
}

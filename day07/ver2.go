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

	requirements := make(map[rune]*[]rune)

	for _, line := range spl {
		if line == "" {
			continue;
		}

		line := []rune(line)

		requirement := line[5]
		id := line[36]

		if list, found := requirements[id]; found {
			(*list) = append(*list, requirement)
		} else {
			newList := []rune{requirement}
			requirements[id] = &newList
		}

		if _, found := requirements[requirement]; !found {
			newList := make([]rune, 0)
			requirements[requirement] = &newList
		}
	}

	steps := make([]rune, 0, len(requirements))
	for key := range requirements {
		steps = append(steps, key)
	}
	sort.Slice(steps, func(i, j int) bool { return steps[i] < steps[j] })

	completed := make(map[rune]bool)

	running := make(map[rune]int)

	elfCount := 5
	timeOffset := 60

	var totalTime = -1

	for {
		totalTime++

		for id, time := range running {
			if time <= 1 {
				completed[id] = true
				delete(running, id)
			} else {
				running[id] = time - 1
			}
		}

		if len(running) < elfCount {
			var anyWaiting = len(running) > 0
			for _, step := range steps {
				if len(running) >= elfCount {
					break
				}

				if _, found := completed[step]; found {
					continue
				}
				if _, found := running[step]; found {
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
					time := int(step - 64) + timeOffset
					running[step] = time
				}
			}
			if !anyWaiting {
				break
			}
		}
	}

	fmt.Println(totalTime)
}

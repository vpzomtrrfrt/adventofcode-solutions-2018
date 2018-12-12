package main

import "fmt"
import "log"
import "os"
import "io/ioutil"
import "strings"
import "math"

type ruleCondition struct {
	left2 bool
	left1 bool
	current bool
	right1 bool
	right2 bool
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	spl := strings.Split(string(bytes), "\n\n")

	initial_line := spl[0]
	rest := spl[1]
	rules_spl := strings.Split(rest, "\n")

	state := make(map[int]bool)

	initial_i1 := strings.Index(initial_line, ": ")
	initial_str := initial_line[(initial_i1 + 2):]

	for idx, char := range initial_str {
		if char == '#' {
			state[idx] = true
		}
	}

	life_rules := make(map[ruleCondition]bool)

	for _, line := range rules_spl {
		if line == "" { continue }

		result := line[len(line) - 1]
		if result == '#' {
			condition := ruleCondition {
				left2: line[0] == '#',
				left1: line[1] == '#',
				current: line[2] == '#',
				right1: line[3] == '#',
				right2: line[4] == '#',
			}
			life_rules[condition] = true
		}
	}

	var lastSum int
	var lastSumDiff int
	var stableCount int

	genCount := 50000000000

	for generation := 1; generation <= genCount; generation++ {
		new_state := make(map[int]bool)

		minPos := math.MaxInt32
		maxPos := math.MinInt32

		for key, value := range state {
			if value {
				if key > maxPos {
					maxPos = key
				}
				if key < minPos {
					minPos = key
				}
			}
		}

		for i := minPos - 2; i < maxPos + 2; i++ {
			condition := ruleCondition {
				left2: state[i - 2],
				left1: state[i - 1],
				current: state[i],
				right1: state[i + 1],
				right2: state[i + 2],
			}
			if life_rules[condition] {
				new_state[i] = true
			}
		}

		state = new_state

		var sum = 0
		for key, value := range state {
			if value {
				sum += key
			}
		}

		sumDiff := sum - lastSum
		if lastSumDiff == sumDiff {
			stableCount++
			if stableCount > 2 {
				// fudge time

				remaining := genCount - generation

				sum += sumDiff * remaining
				generation += remaining
			}
		} else {
			stableCount = 0
		}

		lastSum = sum
		lastSumDiff = sumDiff
	}

	fmt.Println(lastSum)
}

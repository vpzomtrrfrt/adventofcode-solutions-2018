package main

import "fmt"
import "log"
import "os"
import "io/ioutil"
import "strings"
import "strconv"
import "sort"

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	spl := strings.Split(string(bytes), "\n")
	sort.Strings(spl)

	var lastSleep = 0
	var currentGuard = 0
	var sleeps = make(map[int]*[60]int)
	for _, line := range spl {
		if line == "" {
			continue;
		}
		minute_str := line[15:17]
		minute, err := strconv.Atoi(minute_str)

		if err != nil { log.Fatal(err) }

		evt_type := line[19:20]

		if evt_type == "G" {
			// guard begins shift

			i1 := 26
			i2 := strings.Index(line[i1:], " ") + i1

			id_str := line[i1:i2]
			id, err := strconv.Atoi(id_str)
			if err != nil { log.Fatal(err) }

			currentGuard = id
		} else if evt_type == "f" {
			// falls asleep
			lastSleep = minute
		} else if evt_type == "w" {
			// wakes up
			var record *[60]int
			if record_, found := sleeps[currentGuard]; found {
				record = record_
			} else {
				var newRecord [60]int
				record = &newRecord
				sleeps[currentGuard] = record
			}

			for i := lastSleep; i < minute; i++ {
				record[i] += 1
			}
		}
	}

	var topMinuteCount = -1
	var mostConsistentGuard = -1
	var mostConsistentMinute = -1

	for guard, record := range sleeps {
		for minute, count := range record {
			if count > topMinuteCount {
				topMinuteCount = count
				mostConsistentGuard = guard
				mostConsistentMinute = minute
			}
		}
	}
	
	fmt.Println(mostConsistentGuard * mostConsistentMinute)
}

package main

import "fmt"
import "log"
import "os"
import "io/ioutil"
import "strings"
import "strconv"

type point struct {
	x int
	y int
}

func iabs(value int) int {
	if value < 0 {
		return -value
	} else {
		return value
	}
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	spl := strings.Split(string(bytes), "\n")

	points := make([]point, 0, len(spl))

	var maxX = 0
	var maxY = 0

	for _, line := range spl {
		if line == "" {
			continue;
		}

		i1 := strings.Index(line, ",")
		x_str := line[:i1]
		y_str := strings.TrimSpace(line[(i1 + 1):])

		x, err := strconv.Atoi(x_str)
		if err != nil { log.Fatal(err) }

		y, err := strconv.Atoi(y_str)
		if err != nil { log.Fatal(err) }

		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}

		points = append(points, point {
			x: x,
			y: y,
		})
	}

	var totalArea = 0

	border := 200
	threshold := 10000

	for x := -border; x < maxX + border; x++ {
		for y := -border; y < maxY + border; y++ {
			var totalDist = 0

			for _, point := range points {
				dist := iabs(x - point.x) + iabs(y - point.y)

				totalDist += dist
			}

			if totalDist < threshold {
				totalArea++
			}
		}
	}

	fmt.Println(totalArea)
}

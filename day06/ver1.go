package main

import "fmt"
import "log"
import "os"
import "io/ioutil"
import "strings"
import "strconv"
import "math"

type point struct {
	x int
	y int
}

type area struct {
	center point
	count int
	onEdge bool
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

	var results = make([]*area, len(points))
	for idx, center := range points {
		results[idx] = &area {
			center: center,
			count: 0,
			onEdge: false,
		}
	}

	border := 2

	for x := -border; x < maxX + border; x++ {
		for y := -border; y < maxY + border; y++ {
			var shortestDist = math.MaxInt32
			var winningArea = -1
			var hasTie = false

			for idx, area := range results {
				dist := iabs(x - area.center.x) + iabs(y - area.center.y)

				if dist == shortestDist {
					hasTie = true
				} else if dist < shortestDist {
					hasTie = false
					shortestDist = dist
					winningArea = idx
				}
			}

			if !hasTie {
				results[winningArea].count++
				if x == -border || y == -border || x == maxX + border - 1 || y == maxY + border - 1 {
					results[winningArea].onEdge = true
				}
			}
		}
	}

	var largestArea = 0

	for _, area := range results {
		if !area.onEdge {
			if area.count > largestArea {
				largestArea = area.count
			}
		}
	}

	fmt.Println(largestArea)
}

package main

import "fmt"
import "log"
import "os"
import "io/ioutil"
import "strings"
import "strconv"

// This is really slow.

type point struct {
	x int
	y int
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	str := strings.TrimSpace(string(bytes))
	serial, err := strconv.Atoi(str)
	if err != nil { log.Fatal(err) }

	powerMap := make(map[point]int)
	for y := 1; y <= 300; y++ {
		for x := 1; x <= 300; x++ {
			pt := point {
				x: x,
				y: y,
			}

			rack := x + 10
			var power = rack * y
			power += serial
			power *= rack
			power = (power % 1000) / 100
			power -= 5

			powerMap[pt] = power
		}
	}

	var maxPower int
	var maxPowerPoint point
	var maxPowerSize int

	for y := 1; y <= 300; y++ {
		for x := 1; x <= 300; x++ {
			corner := point { x: x, y: y }
			var squarePower = powerMap[corner]

			if squarePower > maxPower {
				maxPower = squarePower
				maxPowerPoint = corner
				maxPowerSize = 1
			}

			for size := 2; size <= (300 + 1 - y); size++ {
				x1 := x + size - 1
				for sy := 0; sy < size; sy++ {
					pt := point { x: x1, y: y + sy }
					squarePower += powerMap[pt]
				}

				y1 := y + size - 1
				for sx := 0; sx < size - 1; sx++ {
					pt := point { x: x + sx, y: y1 }
					squarePower += powerMap[pt]
				}

				if squarePower > maxPower {
					maxPower = squarePower
					maxPowerPoint = corner
					maxPowerSize = size
				}
			}
		}
	}

	fmt.Printf("%d,%d,%d\n", maxPowerPoint.x, maxPowerPoint.y, maxPowerSize)
}

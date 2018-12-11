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

	maxPower := 0
	var maxPowerPoint point

	for y := 1; y <= 300 - 2; y++ {
		for x := 1; x <= 300 - 2; x++ {
			squarePower :=
				powerMap[point { x: x, y: y }] +
				powerMap[point { x: x + 1, y: y }] +
				powerMap[point { x: x + 2, y: y }] +
				powerMap[point { x: x, y: y + 1 }] +
				powerMap[point { x: x + 1, y: y + 1 }] +
				powerMap[point { x: x + 2, y: y + 1 }] +
				powerMap[point { x: x, y: y + 2 }] +
				powerMap[point { x: x + 1, y: y + 2 }] +
				powerMap[point { x: x + 2, y: y + 2 }]

			if squarePower > maxPower {
				maxPower = squarePower
				maxPowerPoint = point { x: x, y: y }
			}
		}
	}

	fmt.Printf("%d,%d\n", maxPowerPoint.x, maxPowerPoint.y)
}

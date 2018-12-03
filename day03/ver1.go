package main

import "fmt"
import "log"
import "os"
import "io/ioutil"
import "strings"
import "strconv"

type Point struct {
	x int
	y int
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	spl := strings.Split(string(bytes), "\n")

	cloth := make(map[Point]int)

	for _, line := range spl {
		if line == "" {
			continue
		}
		i2 := strings.Index(line, ",")
		i1 := strings.LastIndex(line[:i2], " ")

		x_str := line[(i1 + 1):i2]
		x, err := strconv.Atoi(x_str)

		if err != nil { log.Fatal(err) }

		i3 := strings.Index(line[i2:], ":") + i2

		y_str := line[(i2 + 1):i3]
		y, err := strconv.Atoi(y_str)

		if err != nil { log.Fatal(err) }

		i4 := strings.Index(line[i3:], "x") + i3

		width_str := line[(i3 + 2):i4]
		width, err := strconv.Atoi(width_str)

		if err != nil { log.Fatal(err) }

		height_str := line[(i4 + 1):]
		height, err := strconv.Atoi(height_str)

		if err != nil { log.Fatal(err) }

		// fmt.Println(x, y, width, height)

		for xo := 0; xo < width; xo++ {
			for yo := 0; yo < height; yo++ {
				pt := Point {
					x: x + xo,
					y: y + yo,
				}
				if _, found := cloth[pt]; found {
					cloth[pt] += 1
				} else {
					cloth[pt] = 1
				}
			}
		}
	}

	var total = 0
	for _, count := range cloth {
		if count >= 2 {
			total++
		}
	}

	fmt.Println(total)
}

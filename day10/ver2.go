package main

import "fmt"
import "log"
import "os"
import "io/ioutil"
import "strings"
import "strconv"

type vector2 struct {
	x int
	y int
}

type star struct {
	position vector2
	velocity vector2
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bytes), "\n")

	stars := make([]*star, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}

		i1 := strings.Index(line, "<")
		i2 := strings.Index(line[i1:], ",") + i1
		i3 := strings.Index(line[i2:], ">") + i2

		i4 := strings.Index(line[i3:], "<") + i3
		i5 := strings.Index(line[i4:], ",") + i4
		i6 := strings.Index(line[i5:], ">") + i5

		posX_str := strings.TrimSpace(line[(i1 + 1):i2])
		posY_str := strings.TrimSpace(line[(i2 + 2):i3])
		velX_str := strings.TrimSpace(line[(i4 + 1):i5])
		velY_str := strings.TrimSpace(line[(i5 + 2):i6])

		posX, err := strconv.Atoi(posX_str)
		if err != nil { log.Fatal(err) }
		posY, err := strconv.Atoi(posY_str)
		if err != nil { log.Fatal(err) }
		velX, err := strconv.Atoi(velX_str)
		if err != nil { log.Fatal(err) }
		velY, err := strconv.Atoi(velY_str)
		if err != nil { log.Fatal(err) }

		position := vector2 {
			x: posX,
			y: posY,
		}
		velocity := vector2 {
			x: velX,
			y: velY,
		}

		star := star {
			position: position,
			velocity: velocity,
		}

		stars = append(stars, &star)
	}

	lastArea := -1
	wentDown := false

	i := -1

	for {
		minX := 0
		minY := 0
		maxX := 0
		maxY := 0

		sky := make(map[vector2]bool)
		for _, star := range stars {
			sky[star.position] = true
			if star.position.x > maxX {
				maxX = star.position.x
			}
			if star.position.x < minX {
				minX = star.position.x
			}
			if star.position.y > maxY {
				maxY = star.position.y
			}
			if star.position.y < minY {
				minY = star.position.y
			}
			star.position = vector2 {
				x: star.position.x + star.velocity.x,
				y: star.position.y + star.velocity.y,
			}
		}

		area := (maxX - minX) * (maxY - minY)

		if wentDown && area > lastArea {
			fmt.Println(i)
			return
		}
		if area < lastArea {
			wentDown = true
		} else {
			wentDown = false
		}

		lastArea = area
		i++
	}
}

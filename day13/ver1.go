package main

import "fmt"
import "log"
import "os"
import "io/ioutil"
import "strings"
import "sort"

type point struct {
	x int
	y int
}

type direction int8
var (
	directionEAST	direction = 0
	directionNORTH	direction = 1
	directionWEST	direction = 2
	directionSOUTH	direction = 3
)

type rotation int8
var (
	rotationLEFT	rotation = 4
	rotationNONE	rotation = 5
	rotationRIGHT	rotation = 6
)

type pointType int8
var (
	pointTypeCONTINUE		pointType = 7
	pointTypeINTERSECTION		pointType = 8
	pointTypeCORNER_NORTHWEST	pointType = 9
	pointTypeCORNER_NORTHEAST	pointType = 10
)

type cart struct {
	pos point
	dir direction
	nextRotation rotation
}

func (cart *cart) moveForward() {
	switch cart.dir {
	case directionEAST:
		cart.pos.x++
	case directionNORTH:
		cart.pos.y--
	case directionWEST:
		cart.pos.x--
	case directionSOUTH:
		cart.pos.y++
	default:
		panic("invalid direction")
	}
}

func rotate(dir direction, rot rotation) direction {
	if rot == rotationNONE {
		return dir
	} else if rot == rotationLEFT {
		switch dir {
		case directionEAST:
			return directionNORTH
		case directionNORTH:
			return directionWEST
		case directionWEST:
			return directionSOUTH
		case directionSOUTH:
			return directionEAST
		default:
			panic("invalid direction")
		}
	} else if rot == rotationRIGHT {
		switch dir {
		case directionEAST:
			return directionSOUTH
		case directionNORTH:
			return directionEAST
		case directionWEST:
			return directionNORTH
		case directionSOUTH:
			return directionWEST
		default:
			panic("invalid direction")
		}
	} else {
		panic("invalid rotation")
	}
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bytes), "\n")

	trackMap := make(map[point]pointType)
	carts := make([]*cart, 0)

	for y, line := range lines {
		if line == "" {
			continue
		}

		for x, char := range line {
			if char != ' ' {
				pt := point {
					x: x,
					y: y,
				}
				charType := pointTypeCONTINUE

				if char == '+' {
					charType = pointTypeINTERSECTION
				} else if char == '/' {
					charType = pointTypeCORNER_NORTHWEST
				} else if char == '\\' {
					charType = pointTypeCORNER_NORTHEAST
				}

				var newCartDir direction = -1
				if char == '>' {
					newCartDir = directionEAST
				} else if char == '^' {
					newCartDir = directionNORTH
				} else if char == '<' {
					newCartDir = directionWEST
				} else if char == 'v' {
					newCartDir = directionSOUTH
				}

				if newCartDir != -1 {
					carts = append(carts, &cart {
						pos: pt,
						dir: newCartDir,
						nextRotation: rotationLEFT,
					})
				}

				trackMap[pt] = charType
			}
		}
	}

	for {
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].pos.y < carts[j].pos.y {
				return true
			} else if carts[i].pos.y > carts[j].pos.y {
				return false
			} else {
				return carts[i].pos.x < carts[j].pos.x
			}
		})

		for _, cart := range carts {
			cart.moveForward()

			for _, cart2 := range carts {
				if cart != cart2 && cart.pos == cart2.pos {
					// crash!!
					fmt.Printf("%d,%d\n", cart.pos.x, cart.pos.y)
					return
				}
			}

			ptType := trackMap[cart.pos]
			if ptType == pointTypeINTERSECTION {
				cart.dir = rotate(cart.dir, cart.nextRotation)
				switch cart.nextRotation {
				case rotationLEFT:
					cart.nextRotation = rotationNONE
				case rotationNONE:
					cart.nextRotation = rotationRIGHT
				case rotationRIGHT:
					cart.nextRotation = rotationLEFT
				}
			} else if ptType == pointTypeCORNER_NORTHWEST {
				switch cart.dir {
				case directionEAST:
					cart.dir = directionNORTH
				case directionNORTH:
					cart.dir = directionEAST
				case directionWEST:
					cart.dir = directionSOUTH
				case directionSOUTH:
					cart.dir = directionWEST
				}
			} else if ptType == pointTypeCORNER_NORTHEAST {
				switch cart.dir {
				case directionEAST:
					cart.dir = directionSOUTH
				case directionNORTH:
					cart.dir = directionWEST
				case directionWEST:
					cart.dir = directionNORTH
				case directionSOUTH:
					cart.dir = directionEAST
				}
			}
		}
	}
}

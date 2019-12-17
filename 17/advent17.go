package main

import (
	"fmt"
	"image"
	"strconv"
	"strings"

	"github.com/hluk/advent-of-code-2019-golang"
	"github.com/hluk/advent-of-code-2019-golang/intcode"
)

type Pos = image.Point

const (
	AreaScaffold = '#'
	AreaFloor    = '.'
)

type Tile struct{}
type Area = map[Pos]Tile

func toChar(d Pos) byte {
	if d.X == 1 {
		return '>'
	}
	if d.X == -1 {
		return '<'
	}
	if d.Y == 1 {
		return 'v'
	}
	if d.Y == -1 {
		return '^'
	}
	panic("Bad direction")
}

func printArea(a Area, robot Pos, d Pos) {
	minX, maxX := 0, 0
	minY, maxY := 0, 0
	for k, _ := range a {
		minX = adv.Min(minX, k.X)
		minY = adv.Min(minY, k.Y)
		maxX = adv.Max(maxX, k.X)
		maxY = adv.Max(maxY, k.Y)
	}

	pos := Pos{}
	for pos.Y = minY; pos.Y <= maxY; pos.Y++ {
		for pos.X = minX; pos.X <= maxX; pos.X++ {
			_, ok := a[pos]
			if ok {
				if robot == pos {
					dir := toChar(d)
					fmt.Printf("\033[32;1;4m%v\033[0m", string(dir))
				} else {
					fmt.Printf("#")
				}
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println("")
	}
}

func GetArea2(chOut chan intcode.Value) (Area, Pos, Pos) {
	area := Area{}

	pos := Pos{}
	robot := Pos{}
	d := Pos{}

	for {
		if cmd, ok := <-chOut; ok {
			b := byte(cmd)
			if b == 10 {
				pos.X = 0
				pos.Y++
			} else {
				if b != AreaFloor {
					if b != AreaScaffold {
						robot = pos
						switch b {
						case '^':
							d = Pos{0, -1}
						case 'v':
							d = Pos{0, 1}
						case '<':
							d = Pos{-1, 0}
						case '>':
							d = Pos{1, 0}
						case 'X':
							fmt.Println("Robot in space!")
							d = Pos{0, 0}
						}
					}
					area[pos] = Tile{}
				}
				pos.X++
			}
		} else {
			break
		}
	}

	return area, robot, d
}

func GetArea(r []intcode.Value) (Area, Pos, Pos) {
	chIn := make(chan intcode.Value)
	chOut := make(chan intcode.Value)
	go intcode.Run(r, chIn, chOut)
	return GetArea2(chOut)
}

func RunRobot(r []intcode.Value, area Area, robot Pos, d0 Pos) intcode.Value {
	pos := robot
	d := d0

	moves := ""
	move := 0

	for {
		newPos := pos.Add(d)
		if _, ok := area[newPos]; ok {
			pos = newPos
			move++
		} else {
			if move != 0 {
				moves += strconv.Itoa(move)
			}
			move = 0
			right := Pos{-d.Y, d.X}
			left := Pos{d.Y, -d.X}
			if canMove(pos, right, area) {
				d = right
				moves += "R"
			} else if canMove(pos, left, area) {
				d = left
				moves += "L"
			} else {
				break
			}
		}
	}

	prog := moves
	routines := make([]string, 3)

	for i := range routines {
		for j := 2; j < len(moves) && strings.Count(moves[j:], moves[:j]) > 1; j++ {
			routines[i] = moves[:j]
		}
		moves = strings.ReplaceAll(moves, routines[i], "")
		prog = strings.ReplaceAll(prog, routines[i], string('A'+i))
	}

	chIn := make(chan intcode.Value)
	chOut := make(chan intcode.Value)
	go intcode.Run(r, chIn, chOut)

	readPrompt(chOut)

	for i, c := range prog {
		if i != 0 {
			chIn <- ','
		}
		chIn <- intcode.Value(c)
	}
	chIn <- 10

	for _, routine := range routines {
		readPrompt(chOut)
		for i, c := range routine {
			if c >= 'A' && i != 0 {
				chIn <- ','
			}
			chIn <- intcode.Value(c)
			if c >= 'A' && i != len(routine) {
				chIn <- ','
			}
		}
		chIn <- 10
	}

	readPrompt(chOut)

	debug := false
	if debug {
		chIn <- 'y'
		chIn <- 10
		readPrompt(chOut)
		return -1
	}

	chIn <- 'n'
	chIn <- 10

	return readPrompt(chOut)
}

func readPrompt(chOut chan intcode.Value) intcode.Value {
	c := <-chOut
	for ; c != ':' && c != '?' && c < 256; c = <-chOut {
		fmt.Printf(string(c))
	}
	fmt.Printf(string(c))
	nl := <-chOut
	if nl == 10 {
		fmt.Println(" <INPUT>")
	}
	return c
}

func canMove(pos Pos, d Pos, area Area) bool {
	_, ok := area[pos.Add(d)]
	return ok
}

func isIntersection(pos Pos, area Area) bool {
	return canMove(pos, Pos{0, 1}, area) &&
		canMove(pos, Pos{1, 0}, area) &&
		canMove(pos, Pos{0, -1}, area) &&
		canMove(pos, Pos{-1, 0}, area)
}

func findIntersections(area Area) int {
	intersect := 0

	for pos := range area {
		if isIntersection(pos, area) {
			intersect += pos.X * pos.Y
		}
	}

	return intersect
}

func main() {
	r := intcode.LoadRegisters("advent17.txt")
	area, robot, d := GetArea(r)
	//printArea(area, robot, d)

	intersect := findIntersections(area)
	fmt.Println(intersect)

	r[0] = 2
	dust := RunRobot(r, area, robot, d)
	fmt.Println(dust)
}

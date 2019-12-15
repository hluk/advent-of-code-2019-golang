package main

import (
	"fmt"
	"image"
	"os"
	"strconv"

	"github.com/hluk/advent-of-code-2019-golang"
	"github.com/hluk/advent-of-code-2019-golang/intcode"
)

type Pos = image.Point
type Area = map[Pos]byte

const (
	DroidHitWall = 0
	DroidMoved   = 1
	DroidArrived = 2
)

const (
	AreaWall    = '#'
	AreaVisited = '.'
	AreaOxygen  = 'O'
)

const (
	North = 1
	South = 2
	West  = 3
	East  = 4
)

func toDir(d Pos) int {
	if d.X == 1 {
		return East
	}
	if d.X == -1 {
		return West
	}
	if d.Y == 1 {
		return South
	}
	if d.Y == -1 {
		return North
	}
	panic("Bad direction")
}

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

func next(d Pos) Pos {
	return Pos{-d.Y, d.X}
}

func printAreaNear(a Area, pos0 Pos, x byte, wall Pos) {
	limit := Pos{50, 32}

	cols := os.Getenv("X")
	rows := os.Getenv("Y")
	limitX, err1 := strconv.Atoi(cols)
	limitY, err2 := strconv.Atoi(rows)
	if limitX > 0 && err1 == nil && limitY > 0 && err2 == nil {
		limit.X = limitX / 2
		limit.Y = limitY / 2
	}

	pos := pos0.Sub(limit)
	max := pos0.Add(limit)
	for ; pos.Y < max.Y; pos.Y++ {
		for pos.X = pos0.X - limit.X; pos.X < max.X; pos.X++ {
			if pos0 == pos {
				fmt.Printf("\033[32;1;4m%v\033[0m", string(x))
			} else if wall == pos {
				fmt.Printf("\033[31;1;4m#\033[0m")
			} else {
				v, ok := a[pos]
				if ok {
					fmt.Printf("%s", string(v))
				} else {
					fmt.Printf(" ")
				}
			}
		}
		fmt.Println("")
	}
}

func printArea(a Area) {
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
			v, ok := a[pos]
			if ok {
				fmt.Printf("%s", string(v))
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println("")
	}
}

func runDroid(r []Value) (Area, Pos) {
	area := Area{}

	pos := Pos{0, 0}
	d := Pos{0, -1}
	path := []Pos{}
	finish := Pos{0, 0}
	backtrack := false

	chIn := make(chan Value)
	chOut := make(chan Value)
	go intcode.Run(r, chIn, chOut)

	for {
		for i := 0; i < 4 && area[pos.Add(d)] != 0; i++ {
			d = next(d)
		}
		backtrack = false
		if area[pos.Add(d)] != 0 {
			backtrack = true
			if len(path) == 0 {
				break
			}
			pos0 := path[len(path)-1]
			d = Pos{pos0.X - pos.X, pos0.Y - pos.Y}
		}
		dir := toDir(d)
		chIn <- Value(dir)

		if cmd, ok := <-chOut; ok {
			if cmd == DroidHitWall {
				wall := pos.Add(d)
				area[wall] = AreaWall
				//printAreaNear(area, pos, toChar(d), wall)
			} else if cmd == DroidMoved || cmd == DroidArrived {
				if backtrack {
					if len(path) == 0 {
						break
					}
					path = path[:len(path)-1]
				} else {
					path = append(path, pos)
				}

				pos = pos.Add(d)
				area[pos] = AreaVisited
				//printAreaNear(area, pos, toChar(d), pos)

				if cmd == DroidArrived {
					finish = pos
				}
			}
		} else {
			break
		}
	}

	return area, finish
}

func exploreArea(area Area, pos Pos, visit func(Pos, int) bool) {
	ps := []Pos{pos}
	ms := []int{0}

	for len(ps) != 0 {
		pos := ps[0]
		ps = ps[1:]

		m := ms[0]
		ms = ms[1:]

		if area[pos] != '.' {
			continue
		}

		area[pos] = AreaOxygen
		if !visit(pos, m) {
			break
		}

		d := Pos{0, -1}
		ps = append(ps, pos.Add(d))
		d = next(d)
		ps = append(ps, pos.Add(d))
		d = next(d)
		ps = append(ps, pos.Add(d))
		d = next(d)
		ps = append(ps, pos.Add(d))
		ms = append(ms, m+1)
		ms = append(ms, m+1)
		ms = append(ms, m+1)
		ms = append(ms, m+1)
	}
}

func distance(area Area, finish Pos) int {
	area2 := Area{}
	for k, v := range area {
		area2[k] = v
	}

	finishDistance := 0
	exploreArea(area2, Pos{0, 0}, func(pos Pos, distance int) bool {
		if pos == finish {
			finishDistance = distance
			return false
		}
		return true
	})
	return finishDistance
}

func fillTime(area Area, start Pos) int {
	area2 := Area{}
	for k, v := range area {
		area2[k] = v
	}

	minutes := 0
	exploreArea(area2, start, func(pos Pos, distance int) bool {
		minutes = distance
		return true
	})
	return minutes
}

func main() {
	r := intcode.LoadRegisters("advent15.txt")

	area, finish := runDroid(r)
	fmt.Println(distance(area, finish))
	fmt.Println(fillTime(area, finish))
}

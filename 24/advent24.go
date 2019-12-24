package main

import (
	"fmt"
	"io/ioutil"
	"math/bits"
	"strings"
)

type Area uint32

const Rows int = 5
const Columns int = 5

func BugMask(x, y int) Area {
	return 1 << (y*Rows + x)
}

func (a Area) Bug1(x, y int) int {
	return int((a >> (y*Rows + x)) & 1)
}

func (a Area) CountBugsForX(x int) int {
	c := 0
	for y := 0; y < Rows; y++ {
		c += a.Bug1(x, y)
	}
	return c
}

func (a Area) CountBugsForY(y int) int {
	c := 0
	for x := 0; x < Columns; x++ {
		c += a.Bug1(x, y)
	}
	return c
}

func (a Area) HasBug(x, y int) bool {
	return a.Bug1(x, y) == 1
}

func (a Area) CountBug(x, y int) int {
	if 0 <= y && y < Rows && 0 <= x && x < Columns && a.HasBug(x, y) {
		return 1
	}
	return 0
}

func (a Area) AdjecentBugs(x, y int) int {
	return a.CountBug(x-1, y) +
		a.CountBug(x+1, y) +
		a.CountBug(x, y-1) +
		a.CountBug(x, y+1)
}

func (a Area) Bug(x, y int) Area {
	return a | BugMask(x, y)
}

func (a Area) ToString() string {
	str := ""
	for y := 0; y < Rows; y++ {
		for x := 0; x < Columns; x++ {
			if a.HasBug(x, y) {
				str += "#"
			} else {
				str += "."
			}
		}
		str += "\n"
	}
	return str
}

func LoadArea(path string) Area {
	var a Area

	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	dat = dat[:len(dat)]

	txt := string(dat)
	txt = strings.TrimRight(txt, "\n")
	lines := strings.Split(txt, "\n")
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				a = a.Bug(x, y)
			}
		}
	}

	return a
}

func FirstSeen(a Area) Area {
	seen := map[Area]struct{}{}

	minutes := 0
loop:
	for {
		var b Area
		for y := 0; y < Rows; y++ {
			for x := 0; x < Columns; x++ {
				if a.HasBug(x, y) {
					if a.AdjecentBugs(x, y) == 1 {
						b = b.Bug(x, y)
					}
				} else {
					c := a.AdjecentBugs(x, y)
					if c == 1 || c == 2 {
						b = b.Bug(x, y)
					}
				}
			}
		}
		minutes++

		a = b
		if _, ok := seen[a]; ok {
			break loop
		}
		seen[a] = struct{}{}
	}

	return a
}

type Levels map[int]Area

func (l Levels) AdjecentBugs(i, x, y int) int {
	c := l[i].AdjecentBugs(x, y)

	outer := l[i-1]
	if x == 0 {
		c += outer.Bug1(1, 2)
	}

	if y == 0 {
		c += outer.Bug1(2, 1)
	}

	if x == Columns-1 {
		c += outer.Bug1(3, 2)
	}

	if y == Rows-1 {
		c += outer.Bug1(2, 3)
	}

	inner := l[i+1]
	if x == 1 && y == 2 {
		c += inner.CountBugsForX(0)
	}

	if x == 2 && y == 1 {
		c += inner.CountBugsForY(0)
	}

	if x == 3 && y == 2 {
		c += inner.CountBugsForX(Columns - 1)
	}

	if x == 2 && y == 3 {
		c += inner.CountBugsForY(Rows - 1)
	}

	return c
}

func CountAfter(maxMinutes int, a Area) int {
	levels := Levels{}
	levels[0] = a
	levels[-1] |= 0
	levels[1] |= 0

	for minutes := 0; minutes < maxMinutes; minutes++ {
		levels2 := Levels{}

		for i, a := range levels {
			var b Area

			for y := 0; y < Rows; y++ {
				for x := 0; x < Columns; x++ {
					if x == 2 && y == 2 {
						continue
					}
					if a.HasBug(x, y) {
						if levels.AdjecentBugs(i, x, y) == 1 {
							b = b.Bug(x, y)
						}
					} else {
						c := levels.AdjecentBugs(i, x, y)
						if c == 1 || c == 2 {
							b = b.Bug(x, y)
						}
					}
				}
			}

			levels2[i] = b
		}

		levels = levels2

		// Create another neighboring levels if needed
		for i, a := range levels {
			if a.CountBugsForX(0) != 0 ||
				a.CountBugsForY(0) != 0 ||
				a.CountBugsForX(Columns-1) != 0 ||
				a.CountBugsForY(Rows-1) != 0 {
				levels[i-1] |= 0
			}

			if a.AdjecentBugs(2, 2) != 0 {
				levels[i+1] |= 0
			}
		}
	}

	count := 0
	for _, level := range levels {
		count += bits.OnesCount32(uint32(level))
	}
	return count
}

func main() {
	a := LoadArea("advent24.txt")
	fmt.Println(FirstSeen(a))

	fmt.Println(CountAfter(200, a))
}

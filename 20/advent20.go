package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"strings"

	"github.com/hluk/advent-of-code-2019-golang"
)

type Pos = image.Point
type Tile = [2]byte
type Area = map[Pos]Tile
type Ports = map[Tile][2]Pos

func Floor() Tile {
	return Tile{'.', '.'}
}

type QueueItem struct {
	pos   Pos
	steps int
	level int
}
type Queue = []QueueItem

func queueItem(q QueueItem, p Pos, x, y int) QueueItem {
	newPos := p.Add(Pos{x, y})
	return QueueItem{newPos, q.steps + 1, q.level}
}

func Explore(area Area, ports Ports, withLevels bool) int {
	startTile := Tile{'A', 'A'}
	start := ports[startTile][0]
	queue := Queue{{start, -1, 0}}

	finishTile := Tile{'Z', 'Z'}

	type Visited struct {
		level int
		pos   Pos
	}
	visitedMap := map[Visited]int{}

	ds := [4]Pos{
		Pos{0, -1},
		Pos{0, 1},
		Pos{1, 0},
		Pos{-1, 0},
	}

	outX := 0
	outY := 0
	for p := range area {
		outX = adv.Max(outX, p.X)
		outY = adv.Max(outY, p.Y)
	}

	for {
		q := queue[0]
		queue = queue[1:]

		pos := q.pos
		tile, ok := area[pos]
		if !ok {
			continue
		}

		visited := Visited{q.level, pos}
		_, ok2 := visitedMap[visited]
		if ok2 {
			continue
		}
		visitedMap[visited] = q.steps

		if tile[0] >= 'A' && tile != startTile {
			if tile == finishTile {
				if q.level != 0 {
					continue
				}
				return q.steps - 1
			}

			if withLevels {
				if pos.X == 1 || pos.Y == 1 || pos.X == outX || pos.Y == outY {
					q.level--
					if q.level < 0 {
						continue
					}
				} else {
					q.level++
				}
			}

			pos2 := ports[tile]
			if pos2[0] == pos {
				pos = pos2[1]
			} else {
				pos = pos2[0]
			}

			for _, d := range ds {
				if area[pos.Add(d)][0] == '.' {
					pos = pos.Add(d)
					break
				}
			}
		}

		queue = append(
			queue,
			queueItem(q, pos, 0, -1),
			queueItem(q, pos, 1, 0),
			queueItem(q, pos, 0, 1),
			queueItem(q, pos, -1, 0),
		)
	}
}

func isPort(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func LoadArea(path string) (Area, Ports) {
	a := Area{}
	ports := Ports{}

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
			b := byte(c)
			pos := Pos{x, y}
			if b == '.' {
				a[pos] = Floor()
			} else if isPort(b) {
				port := Tile{}
				if x == 1 {
					port = Tile{line[0], b}
				} else if y == 1 {
					port = Tile{lines[0][x], b}
				} else if x > 1 && line[x-1] == '.' {
					port = Tile{b, line[x+1]}
				} else if y > 1 && lines[y-1][x] == '.' {
					port = Tile{b, lines[y+1][x]}
				} else if x+1 < len(line) && line[x+1] == '.' {
					port = Tile{line[x-1], b}
				} else if y+1 < len(lines) && lines[y+1][x] == '.' {
					port = Tile{lines[y-1][x], b}
				}
				if port != (Tile{}) {
					pos2 := ports[port]
					if pos2[0] == (Pos{}) {
						pos2[0] = pos
					} else {
						pos2[1] = pos
					}
					ports[port] = pos2
					a[pos] = port
				}
			}
		}
	}

	return a, ports
}

func main() {
	a, p := LoadArea("advent20.txt")

	steps1 := Explore(a, p, false)
	fmt.Println(steps1)

	steps2 := Explore(a, p, true)
	fmt.Println(steps2)
}

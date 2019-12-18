package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"strings"
)

type Pos = image.Point
type Area = map[Pos]byte
type Keys struct {
	keys int
}

type QueueItem struct {
	pos   Pos
	steps int
	keys  Keys
}
type Queue = []QueueItem

type Pos4 = [4]image.Point
type QueueItem4 struct {
	pos4  Pos4
	keys  Keys
	steps [4]int
}
type Queue4 = []QueueItem4

func (k Keys) ToString() string {
	result := ""
	for key := 'a'; key <= 'z'; key++ {
		if k.Has(byte(key)) {
			result += string(key)
		}
	}
	return result
}

func (k *Keys) Add(key byte) {
	i := key - 'a'
	k.keys = k.keys | (1 << i)
}

func (k Keys) Unlock(door byte) bool {
	i := door - 'A'
	return (k.keys & (1 << i)) != 0
}

func (k Keys) Has(key byte) bool {
	i := key - 'a'
	return (k.keys & (1 << i)) != 0
}

func queueItem(q QueueItem, x, y int) QueueItem {
	newPos := q.pos.Add(Pos{x, y})
	return QueueItem{newPos, q.steps + 1, q.keys}
}

func Explore(area Area, allKeys Keys, start Pos) int {
	queue := Queue{{start, 0, Keys{}}}

	visitedMap := map[Pos]map[Keys]int{}

	for {
		q := queue[0]
		queue = queue[1:]

		v, ok := area[q.pos]
		if !ok {
			continue
		}

		visitedKeyMap, ok := visitedMap[q.pos]
		if ok {
			visitedSteps, ok2 := visitedKeyMap[q.keys]
			if ok2 && visitedSteps <= q.steps {
				continue
			}
		} else {
			visitedKeyMap = map[Keys]int{}
			visitedMap[q.pos] = visitedKeyMap
		}
		visitedKeyMap[q.keys] = q.steps

		if v >= 'A' && v <= 'Z' {
			if !q.keys.Unlock(v) {
				continue
			}
		} else if v >= 'a' && v <= 'z' {
			if !q.keys.Has(v) {
				keys := q.keys
				keys.Add(v)
				if allKeys.keys == keys.keys {
					return q.steps
				}
				q.keys = keys
			}
		}

		queue = append(
			queue,
			queueItem(q, 0, -1),
			queueItem(q, 1, 0),
			queueItem(q, 0, 1),
			queueItem(q, -1, 0),
		)
	}
}

func Explore4(area Area, allKeys Keys, starts Pos4) int {
	queue := Queue4{{starts, Keys{}, [4]int{}}}

	type Visited struct {
		pos4 Pos4
		keys Keys
	}
	visitedMap := map[Visited]int{}

	ds := [4]Pos{
		Pos{0, -1},
		Pos{0, 1},
		Pos{1, 0},
		Pos{-1, 0},
	}
	minSteps := 2318

	for len(queue) != 0 {
		q := queue[0]
		queue = queue[1:]

		allSteps := q.steps[0] + q.steps[1] + q.steps[2] + q.steps[3]
		if allSteps >= minSteps {
			continue
		}

		// check if already visited
		visited := Visited{q.pos4, q.keys}
		visitedSteps, ok := visitedMap[visited]
		if ok && visitedSteps <= allSteps {
			continue
		}
		visitedMap[visited] = allSteps

		for i := 0; i < 4; i++ {
			for _, d := range ds {
				keys := q.keys

				type PosDiff struct {
					p     Pos
					d     Pos
					steps int
				}
				ps := []PosDiff{{q.pos4[i], d, 0}}
				for len(ps) != 0 {
					pd := ps[0]
					ps = ps[1:]

					p := pd.p
					d := pd.d
					steps := pd.steps

					p2 := p.Add(d)
					a2 := area[p2]

					for a2 != 0 {
						if 'a' <= a2 && a2 <= 'z' {
							if !keys.Has(a2) {
								keys2 := keys
								keys2.Add(a2)
								stepsKey := steps + 1

								if keys2 == allKeys {
									usedSteps := allSteps + stepsKey
									if usedSteps < minSteps {
										minSteps = usedSteps
									}
								}
								q2 := q
								q2.pos4[i] = p2
								q2.steps[i] += stepsKey
								q2.keys = keys2
								queue = append(queue, q2)
								break
							}
						} else if 'A' <= a2 && a2 <= 'Z' {
							if !keys.Unlock(a2) {
								break
							}
						}

						p = p2
						steps++

						dl := Pos{d.Y, -d.X}
						dr := Pos{-d.Y, d.X}

						p2 = p.Add(d)
						pl := p.Add(dl)
						pr := p.Add(dr)
						al := area[pl]
						ar := area[pr]
						a2 = area[p2]

						if al != 0 {
							ps = append(ps, PosDiff{p, dl, steps})
						}

						if ar != 0 {
							ps = append(ps, PosDiff{p, dr, steps})
						}
					}
				}
			}
		}
	}
	return minSteps
}

func LoadArea(path string) (Area, Keys, Pos4) {
	a := Area{}
	keys := Keys{}
	starts := Pos4{}
	startCount := 0

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
			if b == '#' {
				continue
			}
			if b == '@' {
				starts[startCount] = Pos{x, y}
				startCount++
				b = '.'
			}
			a[Pos{x, y}] = b
			if b >= 'a' && b <= 'z' {
				keys.Add(b)
			}
		}
	}

	return a, keys, starts
}

func Split(a Area, start Pos) Pos4 {
	delete(a, start.Add(Pos{1, 0}))
	delete(a, start.Add(Pos{-1, 0}))
	delete(a, start.Add(Pos{0, 1}))
	delete(a, start.Add(Pos{0, -1}))
	return Pos4{
		start.Add(Pos{1, 1}),
		start.Add(Pos{1, -1}),
		start.Add(Pos{-1, 1}),
		start.Add(Pos{-1, -1}),
	}
}

func main() {
	a, keys, starts := LoadArea("advent18.txt")

	steps := Explore(a, keys, starts[0])
	fmt.Println(steps)

	starts = Split(a, starts[0])
	steps = Explore4(a, keys, starts)
	fmt.Println(steps)
}

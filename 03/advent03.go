package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Point = image.Point

type Wire struct {
	data map[Point]int
}

func (w Wire) Get(x, y int) int {
	v, ok := w.data[Point{x, y}]
	if ok {
		return v
	}
	return 0
}

func (w *Wire) Set(x, y, d int) {
	xy := Point{x, y}
	v, ok := w.data[xy]
	if ok {
		d = Min(d, v)
	}
	w.data[xy] = d
}

func Crossings(w1, w2 Wire) []Point {
	result := []Point{}
	for k := range w1.data {
		if _, ok := w2.data[k]; ok {
			if k.X == 0 && k.Y == 0 {
				continue
			}
			result = append(result, k)
		}
	}
	return result
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Distance(p Point) int {
	return Abs(p.X) + Abs(p.Y)
}

func CrossingsShortestDistance(w1, w2 Wire) int {
	cs := Crossings(w1, w2)
	d := Distance(cs[0])
	for _, c := range cs {
		d = Min(d, Distance(c))
	}
	return d
}

func CrossingsShortestSignal(w1, w2 Wire) int {
	result := math.MaxInt32
	for k := range w1.data {
		if v2, ok := w2.data[k]; ok {
			if k.X == 0 && k.Y == 0 {
				continue
			}
			result = Min(result, w1.data[k]+v2)
		}
	}
	return result
}

func LoadWire(txt string) Wire {
	w := Wire{}
	w.data = make(map[Point]int)
	dirs := strings.Split(txt, ",")
	x, y, d := 0, 0, 0

	addWire := func(v *int, dv int, amount int) {
		stop := *v + dv*(amount+1)
		for *v += dv; *v != stop; *v += dv {
			d++
			w.Set(x, y, d)
		}
		*v -= dv
	}

	for _, dir := range dirs {
		where := dir[0]
		amount, err := strconv.Atoi(dir[1:])
		if err != nil {
			panic(err)
		}

		switch where {
		case 'R':
			addWire(&x, 1, amount)
		case 'L':
			addWire(&x, -1, amount)
		case 'U':
			addWire(&y, 1, amount)
		case 'D':
			addWire(&y, -1, amount)
		}
	}
	return w
}

func main() {
	dat, err := ioutil.ReadFile("advent03.txt")
	if err != nil {
		panic(err)
	}

	txt := string(dat)
	txt = strings.TrimRight(txt, "\n")
	lines := strings.Split(txt, "\n")
	w1 := LoadWire(lines[0])
	w2 := LoadWire(lines[1])

	fmt.Println(CrossingsShortestDistance(w1, w2))
	fmt.Println(CrossingsShortestSignal(w1, w2))
}

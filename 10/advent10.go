//revive:disable:exported
package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

type Pos = image.Point
type MapData = map[Pos]struct{}

type Map struct {
	Map MapData
	W   int
	H   int
	D   []Pos
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func CreateMap() Map {
	return Map{MapData{}, 0, 0, []Pos{}}
}

func (m *Map) Set(p Pos) {
	m.W = Max(m.W, p.X+1)
	m.H = Max(m.H, p.Y+1)
	m.Map[p] = struct{}{}
}

func (m *Map) Unset(p Pos) {
	delete(m.Map, p)
}

func (m Map) Get(p Pos) bool {
	_, ok := m.Map[p]
	return ok
}

func (m Map) VisibleAsteroidInLOS(p0, pd Pos) Pos {
	p := p0.Add(pd)
	for ; 0 <= p.Y && p.Y < m.H && 0 <= p.X && p.X < m.W; p = p.Add(pd) {
		if m.Get(p) {
			return p
		}
	}
	return p0
}

func Angle(p Pos) float64 {
	rad := math.Atan2(float64(p.Y), float64(p.X))
	deg := rad*(180/math.Pi) + 90
	deg = math.Mod(deg+360, 360)
	return deg
}

func (m Map) Deltas() []Pos {
	ds := []Pos{}
	for xd := -m.W + 1; xd < m.W; xd++ {
		for yd := -m.H + 1; yd < m.H; yd++ {
			gcd := GCD(xd, yd)
			if Abs(gcd) != 1 || gcd == 0 {
				continue
			}
			ds = append(ds, Pos{xd, yd})
		}
	}
	sort.Slice(ds, func(i, j int) bool {
		return Angle(ds[i]) < Angle(ds[j])
	})
	return ds
}

func (m Map) Asteroids(p0 Pos, ch chan Pos) {
	for _, d := range m.D {
		p := m.VisibleAsteroidInLOS(p0, d)
		if p != p0 {
			ch <- p
		}
	}
	close(ch)
}

func (m Map) VisibleAsteroids(p0 Pos) int {
	ch := make(chan Pos)
	go m.Asteroids(p0, ch)
	count := 0
	for range ch {
		count++
	}
	return count
}

func LoadMap(path string) Map {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	dat = dat[:len(dat)-1]

	txt := string(dat)
	txt = strings.TrimRight(txt, "\n")
	lines := strings.Split(txt, "\n")
	m := CreateMap()
	for y, line := range lines {
		for x, v := range line {
			if v == '#' {
				m.Set(Pos{x, y})
			}
		}
	}
	m.D = m.Deltas()
	return m
}

func main() {
	m := LoadMap("advent10.txt")
	visibleMax := 0
	pMax := Pos{0, 0}
	for p := range m.Map {
		visible := m.VisibleAsteroids(p)
		if visibleMax < visible {
			visibleMax = visible
			pMax = p
		}
	}
	fmt.Println(pMax, visibleMax)

	ch := make(chan Pos)
	count := 0
	go m.Asteroids(pMax, ch)
	for p := range ch {
		count++
		if count == 200 {
			fmt.Println(p, p.X*100+p.Y)
			break
		}
		m.Unset(p)
	}
}

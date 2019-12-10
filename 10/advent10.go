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
type Hits = map[Pos][]Pos

type Map struct {
	Map MapData
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
	return Map{MapData{}}
}

func (m *Map) Set(p Pos) {
	m.Map[p] = struct{}{}
}

func (m Map) Get(p Pos) bool {
	_, ok := m.Map[p]
	return ok
}

func Angle(p Pos) float64 {
	rad := math.Atan2(float64(p.Y), float64(p.X))
	deg := rad*(180/math.Pi) + 90
	deg = math.Mod(deg+360, 360)
	return deg
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
	return m
}

func (m *Map) Hits(p0 Pos) Hits {
	hits := Hits{}
	for p := range m.Map {
		d := p.Sub(p0)
		scale := Abs(GCD(d.X, d.Y))
		if scale == 0 {
			continue
		}
		d = Pos{d.X / scale, d.Y / scale}
		hits[d] = append(hits[d], p)
	}
	return hits
}

func Distance(p0, p1 Pos) int {
	d := p1.Sub(p0)
	return Abs(d.X) + Abs(d.Y)
}

func HitAt(hits Hits, p0 Pos, count int) Pos {
	ds := []Pos{}
	for d := range hits {
		ds = append(ds, d)
	}
	sort.Slice(ds, func(i, j int) bool {
		return Angle(ds[i]) < Angle(ds[j])
	})

	for i := 0; ; i++ {
		for _, d := range ds {
			ps := hits[d]
			if i < len(ps) {
				count--
				if count == 0 {
					sort.Slice(ps, func(i, j int) bool {
						return Distance(p0, ps[i]) < Distance(p0, ps[j])
					})
					return ps[i]
				}
			}
		}
	}
}

func main() {
	m := LoadMap("advent10.txt")

	hitsMax := Hits{}
	pMax := Pos{0, 0}
	for p := range m.Map {
		hits := m.Hits(p)
		if len(hitsMax) < len(hits) {
			hitsMax = hits
			pMax = p
		}
	}
	fmt.Println(pMax, len(hitsMax))

	p := HitAt(hitsMax, pMax, 200)
	fmt.Println(p.X*100 + p.Y)
}

//revive:disable:exported
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Pos struct {
	X, Y, Z int
}

type Moon struct {
	P Pos
	V Pos
}

type Moons struct {
	a Moon
	b Moon
	c Moon
	d Moon
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Sign(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return (a * b) / GCD(a, b)
}

func (p Pos) Energy() int {
	return Abs(p.X) + Abs(p.Y) + Abs(p.Z)
}

func (p *Pos) Add(other Pos) {
	p.X += other.X
	p.Y += other.Y
	p.Z += other.Z
}

func (m Moon) PotentialEnergy() int {
	return m.P.Energy()
}

func (m Moon) KineticEnergy() int {
	return m.V.Energy()
}

func (m Moon) Energy() int {
	return m.PotentialEnergy() * m.KineticEnergy()
}

func parseComponent(comp string) int {
	str := strings.Split(comp, "=")[1]
	v, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return v
}

func (m Moons) TotalEnergy() int {
	return m.a.Energy() + m.b.Energy() + m.c.Energy() + m.d.Energy()
}

func Accel(a Moon, b Moon, c Moon, d Moon) Pos {
	return Pos{
		Sign(b.P.X, a.P.X) + Sign(c.P.X, a.P.X) + Sign(d.P.X, a.P.X),
		Sign(b.P.Y, a.P.Y) + Sign(c.P.Y, a.P.Y) + Sign(d.P.Y, a.P.Y),
		Sign(b.P.Z, a.P.Z) + Sign(c.P.Z, a.P.Z) + Sign(d.P.Z, a.P.Z),
	}
}

func (m *Moons) Simulate() {
	aa := Accel(m.a, m.b, m.c, m.d)
	ab := Accel(m.b, m.c, m.d, m.a)
	ac := Accel(m.c, m.d, m.a, m.b)
	ad := Accel(m.d, m.a, m.b, m.c)

	m.a.V.Add(aa)
	m.b.V.Add(ab)
	m.c.V.Add(ac)
	m.d.V.Add(ad)

	m.a.P.Add(m.a.V)
	m.b.P.Add(m.b.V)
	m.c.P.Add(m.c.V)
	m.d.P.Add(m.d.V)
}

func LoadMoons(path string) Moons {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	dat = dat[:len(dat)-1]

	txt := string(dat)
	txt = strings.TrimRight(txt, "\n")
	lines := strings.Split(txt, "\n")

	m := make([]Moon, len(lines))
	for i, line := range lines {
		line = strings.TrimRight(line, "\n")
		line = strings.TrimRight(line, ">")
		t := strings.Split(line, ", ")
		m[i].P = Pos{
			parseComponent(t[0]),
			parseComponent(t[1]),
			parseComponent(t[2]),
		}
	}

	return Moons{m[0], m[1], m[2], m[3]}
}

func AccelAxis(a, b, c, d int) int {
	return Sign(b, a) + Sign(c, a) + Sign(d, a)
}

func SimulateAxis(m, v []int) {
	v[0] += AccelAxis(m[0], m[1], m[2], m[3])
	v[1] += AccelAxis(m[1], m[0], m[2], m[3])
	v[2] += AccelAxis(m[2], m[0], m[1], m[3])
	v[3] += AccelAxis(m[3], m[0], m[1], m[2])

	m[0] += v[0]
	m[1] += v[1]
	m[2] += v[2]
	m[3] += v[3]
}

func FindAxisLoop(m [4]int) int {
	m0 := m
	v0 := [4]int{}
	v := v0
	for i := 0; ; i++ {
		SimulateAxis(m[:], v[:])
		if m0 == m && v0 == v {
			return i + 1
		}
	}
}

func (m Moons) FindLoop() int {
	x := FindAxisLoop([4]int{m.a.P.X, m.b.P.X, m.c.P.X, m.d.P.X})
	y := FindAxisLoop([4]int{m.a.P.Y, m.b.P.Y, m.c.P.Y, m.d.P.Y})
	z := FindAxisLoop([4]int{m.a.P.Z, m.b.P.Z, m.c.P.Z, m.d.P.Z})
	return LCM(x, LCM(y, z))
}

func main() {
	moons := LoadMoons("advent12.txt")

	m := moons
	for i := 0; i < 1000; i++ {
		m.Simulate()
	}
	fmt.Println(m.TotalEnergy())

	fmt.Println(moons.FindLoop())
}

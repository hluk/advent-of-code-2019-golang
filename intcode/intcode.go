package intcode

import (
	"fmt"
	"image"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/hluk/advent-of-code-2019-golang"
)

type Value = int64
type OpCode = int64
type Mode = int64

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

type Op struct {
	Op    OpCode
	Mode1 Mode
	Mode2 Mode
	Mode3 Mode
}

type Memory struct {
	Data map[Value]Value
}

func CreateMemory(r []Value) Memory {
	m := Memory{map[Value]Value{}}
	for i, v := range r {
		m.Data[Value(i)] = v
	}
	return m
}

func (m Memory) Get(i Value) Value {
	if i < 0 {
		panic("Bad memory address")
	}
	return m.Data[i]
}

func (m *Memory) Set(i Value, v Value, mode Mode, addrBase Value) {
	if mode == 2 {
		i += addrBase
	}
	if i < 0 {
		panic("Bad memory address")
	}
	m.Data[i] = v
}

func LoadOp(opcode OpCode) Op {
	mode3 := opcode / 10000
	mode2 := opcode % 10000 / 1000
	mode1 := opcode % 1000 / 100
	op := opcode % 100
	return Op{op, mode1, mode2, mode3}
}

func LoadArg(value Value, mode Mode, addrBase Value, m Memory) Value {
	if mode == 1 {
		return value
	}

	if mode == 2 {
		return m.Get(value + addrBase)
	}

	return m.Get(value)
}

func (op *Op) Load1(ip Value, addrBase Value, m Memory) Value {
	return LoadArg(m.Get(ip+1), op.Mode1, addrBase, m)
}

func (op *Op) Load2(ip Value, addrBase Value, m Memory) Value {
	return LoadArg(m.Get(ip+2), op.Mode2, addrBase, m)
}

func Run(r []Value, chIn, chOut chan Value) {
	var ip Value
	var addrBase Value
	m := CreateMemory(r)
	for {
		op := LoadOp(m.Get(ip))
		switch op.Op {
		case 1, 2:
			a := op.Load1(ip, addrBase, m)
			b := op.Load2(ip, addrBase, m)
			c := m.Get(ip + 3)
			if op.Op == 1 {
				m.Set(c, a+b, op.Mode3, addrBase)
			} else {
				m.Set(c, a*b, op.Mode3, addrBase)
			}
			ip += 4

		case 3:
			a := m.Get(ip + 1)
			m.Set(a, <-chIn, op.Mode1, addrBase)
			ip += 2

		case 4:
			a := op.Load1(ip, addrBase, m)
			chOut <- a
			ip += 2

		case 5, 6:
			a := op.Load1(ip, addrBase, m)
			if (a == 0) == (op.Op == 5) {
				ip += 3
			} else {
				ip = op.Load2(ip, addrBase, m)
			}

		case 7:
			a := op.Load1(ip, addrBase, m)
			b := op.Load2(ip, addrBase, m)
			c := m.Get(ip + 3)
			if a < b {
				m.Set(c, 1, op.Mode3, addrBase)
			} else {
				m.Set(c, 0, op.Mode3, addrBase)
			}
			ip += 4

		case 8:
			a := op.Load1(ip, addrBase, m)
			b := op.Load2(ip, addrBase, m)
			c := m.Get(ip + 3)
			if a == b {
				m.Set(c, 1, op.Mode3, addrBase)
			} else {
				m.Set(c, 0, op.Mode3, addrBase)
			}
			ip += 4

		case 9:
			a := op.Load1(ip, addrBase, m)
			addrBase += a
			ip += 2

		case 99:
			close(chOut)
			return

		default:
			panic("invalid op code")
		}
	}
}

func LoadRegisters(path string) []Value {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	dat = dat[:len(dat)-1]

	txt := string(dat)
	txt = strings.TrimRight(txt, "\n")
	strOps := strings.Split(txt, ",")

	r := make([]Value, len(strOps))
	for i, strOp := range strOps {
		op, err := strconv.ParseInt(strOp, 10, 64)
		if err != nil {
			panic(err)
		}

		r[i] = op
	}

	return r
}

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

	chIn := make(chan Value, 1)
	chOut := make(chan Value)
	go Run(r, chIn, chOut)

	for {
		if cmd, ok := <-chOut; ok {
			if cmd == -1 {
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
			} else if cmd == DroidHitWall {
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
	r := LoadRegisters("advent15.txt")

	area, finish := runDroid(r)
	fmt.Println(distance(area, finish))
	fmt.Println(fillTime(area, finish))
}

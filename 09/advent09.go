//revive:disable:exported
package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Value = int64
type OpCode = int64
type Mode = int64

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

func RunWithInput(input Value, r []Value) {
	chIn := make(chan Value)
	chOut := make(chan Value)
	go Run(r, chIn, chOut)
	chIn <- input
	for {
		if v, ok := <-chOut; ok {
			fmt.Println(v)
		} else {
			break
		}
	}
}

func main() {
	dat, err := ioutil.ReadFile("advent09.txt")
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

	RunWithInput(1, r)
	RunWithInput(2, r)
}

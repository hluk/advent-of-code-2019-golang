package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Op struct {
	Op    int
	Mode1 bool
	Mode2 bool
	Mode3 bool
}

func LoadOp(opcode int) Op {
	mode3 := opcode/10000 > 0
	mode2 := opcode%10000/1000 > 0
	mode1 := opcode%1000/100 > 0
	op := opcode % 100
	return Op{op, mode1, mode2, mode3}
}

func LoadArg(value int, mode bool, r []int) int {
	if mode {
		return value
	}

	return r[value]
}

func (op *Op) Load1(ip int, r []int) int {
	return LoadArg(r[ip+1], op.Mode1, r)
}

func (op *Op) Load2(ip int, r []int) int {
	return LoadArg(r[ip+2], op.Mode2, r)
}

// Run op codes on r registers.
func Run(r []int, input []int) []int {
	outputs := []int{}
	ip := 0
	for {
		op := LoadOp(r[ip])
		switch op.Op {
		case 1, 2:
			a := op.Load1(ip, r)
			b := op.Load2(ip, r)
			c := r[ip+3]
			if op.Op == 1 {
				r[c] = a + b
			} else {
				r[c] = a * b
			}
			ip += 4

		case 3:
			a := r[ip+1]
			r[a], input = input[0], input[1:]
			ip += 2

		case 4:
			a := op.Load1(ip, r)
			outputs = append(outputs, a)
			ip += 2

		case 5, 6:
			a := op.Load1(ip, r)
			if (a == 0) == (op.Op == 5) {
				ip += 3
			} else {
				ip = op.Load2(ip, r)
			}

		case 7:
			a := op.Load1(ip, r)
			b := op.Load2(ip, r)
			c := r[ip+3]
			if a < b {
				r[c] = 1
			} else {
				r[c] = 0
			}
			ip += 4

		case 8:
			a := op.Load1(ip, r)
			b := op.Load2(ip, r)
			c := r[ip+3]
			if a == b {
				r[c] = 1
			} else {
				r[c] = 0
			}
			ip += 4

		case 99:
			return outputs

		default:
			panic("invalid op code")
		}
	}
}

func main() {
	dat, err := ioutil.ReadFile("advent05.txt")
	if err != nil {
		panic(err)
	}

	txt := string(dat)
	txt = strings.TrimRight(txt, "\n")
	strOps := strings.Split(txt, ",")
	r := make([]int, len(strOps))
	for i, strOp := range strOps {
		op, err := strconv.Atoi(strOp)
		if err != nil {
			panic(err)
		}

		r[i] = op
	}

	r1 := make([]int, len(r))
	copy(r1, r)
	result := Run(r1, []int{1})
	fmt.Println(result)

	r2 := make([]int, len(r))
	copy(r2, r)
	result = Run(r2, []int{5})
	fmt.Println(result)
}

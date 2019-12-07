//revive:disable:exported
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

func Run(r []int, chIn, chOut, chExit chan int) {
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
			r[a] = <-chIn
			ip += 2

		case 4:
			a := op.Load1(ip, r)
			chOut <- a
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
			if chExit != nil {
				chExit <- 1
			}
			return

		default:
			panic("invalid op code")
		}
	}
}

func Permutations(orig []int) func() []int {
	n := len(orig)
	result := make([]int, n)

	// offsets in a Fisher-Yates shuffle algorithm
	// https://stackoverflow.com/a/30230552/454171
	p := make([]int, n-1)

	next := func() []int {
		if p[0] < n {
			copy(result, orig)
			for i, v := range p {
				result[i], result[i+v] = result[i+v], result[i]
			}

			for i := n - 2; i >= 0; i-- {
				if i == 0 || p[i] < n-i-1 {
					p[i]++
					break
				}
				p[i] = 0
			}

			return result
		}
		return nil
	}
	return next
}

func MaxThrust(r, orig []int) int {
	max := 0
	n := 5
	p := Permutations(orig)
	for {
		pp := p()
		if pp == nil {
			break
		}

		ch := make([](chan int), n)
		for i := 0; i < n; i++ {
			ch[i] = make(chan int, 100)
			ch[i] <- pp[i]
		}

		chExit := make(chan int)
		for i := 0; i < n; i++ {
			rr := make([]int, len(r))
			copy(rr, r)
			if i == 0 {
				ch[n-1] <- 0
				go Run(rr, ch[n-1], ch[0], nil)
			} else if i == n-1 {
				go Run(rr, ch[n-2], ch[i], chExit)
			} else {
				go Run(rr, ch[i-1], ch[i], nil)
			}
		}

		<-chExit

		out := <-ch[4]
		if out > max {
			max = out
		}
	}

	return max
}

func main() {
	dat, err := ioutil.ReadFile("advent07.txt")
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

	fmt.Println(MaxThrust(r, []int{0, 1, 2, 3, 4}))
	fmt.Println(MaxThrust(r, []int{5, 6, 7, 8, 9}))
}

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Run op codes on r registers.
func Run(r []int) {
	ip := 0
	for {
		switch r[ip] {
		case 1:
			a := r[ip+1]
			b := r[ip+2]
			c := r[ip+3]
			r[c] = r[a] + r[b]
			ip += 4
		case 2:
			a := r[ip+1]
			b := r[ip+2]
			c := r[ip+3]
			r[c] = r[a] * r[b]
			ip += 4
		case 99:
			return
		default:
			panic("invalid op code")
		}
	}
}

// FixRun fixes 1 2 op codes and runs the program.
func FixRun(noun int, verb int, rr []int) int {
	r := make([]int, len(rr))
	copy(r, rr)
	r[1] = noun
	r[2] = verb
	Run(r)
	return r[0]
}

func main() {
	dat, err := ioutil.ReadFile("advent02.txt")
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

	result := FixRun(12, 2, r)
	fmt.Println(result)

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			result = FixRun(noun, verb, r)
			if result == 19690720 {
				fmt.Printf("%d (%d %d)\n", 100*noun+verb, noun, verb)
			}
		}
	}
}

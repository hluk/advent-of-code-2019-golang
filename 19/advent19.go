package main

import (
	"fmt"

	"github.com/hluk/advent-of-code-2019-golang/intcode"
)

func Pull(r []intcode.Value, x, y int) bool {
	chIn := make(chan intcode.Value)
	chOut := make(chan intcode.Value)
	go intcode.Run(r, chIn, chOut)
	chIn <- intcode.Value(x)
	chIn <- intcode.Value(y)
	pull := <-chOut
	return pull == 1
}

func Fits(r []intcode.Value, s, x, y int) bool {
	xy := Pull(r, x, y)
	x0 := Pull(r, x+s-1, y)
	x1 := Pull(r, x+s, y)
	y0 := Pull(r, x, y+s-1)
	y1 := Pull(r, x, y+s)
	return xy && x0 && y0 && !x1 && !y1
}

func main() {
	r := intcode.LoadRegisters("advent19.txt")

	area := 0
	x0, y0 := 0, 0
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			if Pull(r, x, y) {
				area++
				x0, y0 = x, y
			}
		}
	}
	fmt.Println(area)

	x := x0
	y := y0
	s := 100
	for {
		for !Pull(r, x, y) {
			x++
		}

		if Pull(r, x+s-1, y) {
			for Pull(r, x+s, y) {
				x++
			}
		} else {
			y++
		}

		if Pull(r, x, y+s-1) {
			for Pull(r, x, y+s) {
				y++
			}
		} else {
			y++
		}

		if Fits(r, s, x, y) {
			break
		}
	}

	fmt.Println(x*10000 + y)
}

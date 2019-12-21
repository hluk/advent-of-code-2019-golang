package main

import (
	"fmt"

	"github.com/hluk/advent-of-code-2019-golang/intcode"
)

func readPrompt(chOut chan intcode.Value) intcode.Value {
	c := <-chOut
	for c == 10 {
		c = <-chOut
	}

	if c > 255 {
		return c
	}

	c0 := c
	for ; c != 10 && c != '@'; c = <-chOut {
		fmt.Printf(string(c))
	}
	fmt.Println()

	return c0
}

func do(inst string, chIn chan intcode.Value) {
	fmt.Printf("> ")
	for _, r := range inst {
		fmt.Printf(string(r))
		chIn <- intcode.Value(r)
	}
	fmt.Println()
	chIn <- 10
}

func walk(r []intcode.Value, walk bool) {
	chIn := make(chan intcode.Value)
	chOut := make(chan intcode.Value)
	go intcode.Run(r, chIn, chOut)

	c := readPrompt(chOut)

	do("NOT A J", chIn)

	do("NOT B T", chIn)
	do("OR T J", chIn)

	do("NOT C T", chIn)
	do("OR T J", chIn)

	do("AND D J", chIn)

	if walk {
		do("WALK", chIn)
	} else {
		do("NOT E T", chIn)
		do("NOT T T", chIn)
		do("OR H T", chIn)
		do("AND T J", chIn)

		do("RUN", chIn)
	}

	if c > 255 {
		fmt.Println(c)
		return
	}

	for {
		c := <-chOut

		if c > 255 {
			fmt.Println(c)
			return
		}

		fmt.Printf(string(c))
	}
}

func main() {
	r := intcode.LoadRegisters("advent21.txt")

	walk(r, true)
	walk(r, false)
}

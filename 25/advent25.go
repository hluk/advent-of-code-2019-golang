package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/hluk/advent-of-code-2019-golang/intcode"
)

func main() {
	r := intcode.LoadRegisters("advent25.txt")

	chIn := make(chan intcode.Value)
	chOut := make(chan intcode.Value)
	chRead := make(chan intcode.Value)
	go intcode.RunWithRead(r, chIn, chOut, chRead)

	reader := bufio.NewReader(os.Stdin)

	initial := []string{
		"w", "n",
		"take wreath",

		"n", "n",
		"take spool of cat6",
		"s", "s",

		"e",
		"take fixed point",
		"w",

		"s", "e", "e",
		"take sand",
		"w", "s",
		"take ornament",
		"n", "w",

		"s",
		//"take giant electromagnet",

		//"e",
		//"take escape pod",

		"e", "e", "e",
		"take space law space brochure",

		"s",
		"take fuel cell",

		"s",
		"drop ornament", // too heavy
		"drop fuel cell",
		"drop spool of cat6",
		//"drop wreath",
		//"drop space law space brochure",
		//"drop fixed point",
		//"drop sand",
		"w",

		"inv",
	}
	for {
		select {
		case c := <-chOut:
			fmt.Printf(string(c))

		case <-chRead:
			fmt.Printf("> ")
			input := ""

			if len(initial) == 0 {
				input, _ = reader.ReadString('\n')
				if len(input) == 0 {
					return
				}
				input = input[:len(input)-1]
			} else {
				input = initial[0]
				initial = initial[1:]
				fmt.Println(input)
			}

			switch input {
			case "s":
				input = "south"
			case "n":
				input = "north"
			case "w":
				input = "west"
			case "e":
				input = "east"
			}

			for _, r := range input {
				chIn <- intcode.Value(r)
				<-chRead
			}
			chIn <- 10
		}
	}
}

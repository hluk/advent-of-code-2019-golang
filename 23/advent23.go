package main

import (
	"fmt"
	"reflect"

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

func main() {
	r := intcode.LoadRegisters("advent23.txt")

	chIn := [50]chan intcode.Value{}
	chOut := [50]chan intcode.Value{}
	chRead := [50]chan intcode.Value{}

	for i := 0; i < 50; i++ {
		chIn[i] = make(chan intcode.Value, 1000)
		chOut[i] = make(chan intcode.Value, 1000)
		chRead[i] = make(chan intcode.Value, 1000)
		go intcode.RunWithRead(r, chIn[i], chOut[i], chRead[i])
		<-chRead[i]
		chIn[i] <- intcode.Value(i)
	}

	cases := make([]reflect.SelectCase, 100)
	for i, ch := range chOut {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
	}
	for i, ch := range chRead {
		cases[50+i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
	}

	empty := 0
	var natX, natY intcode.Value
	var lastSentY intcode.Value = -1
	for {
		i, addr, ok := reflect.Select(cases)
		fmt.Println(i, addr, ok)
		if ok {
			if i < 50 {
				j := addr.Int()
				x := <-chOut[i]
				y := <-chOut[i]
				empty = 0
				if j == 255 {
					fmt.Println("NAT", x, y)
					natX = x
					natY = y
				} else {
					fmt.Println(i, "->", j, ":", x, y)
					chIn[j] <- x
					chIn[j] <- y
				}
			} else {
				i -= 50
				if empty > 1000 && i == 0 {
					fmt.Println("Restarting work")

					if lastSentY == natY {
						if lastSentY >= 17062 {
							fmt.Println("Too high")
						} else {
							fmt.Println(natY)
							break
						}
					}
					lastSentY = natY

					chIn[0] <- natX
					chIn[0] <- natY
					empty = 0
				} else {
					fmt.Println(i, "<-", -1)
					chIn[i] <- -1
					empty++
				}
			}
		} else {
			fmt.Println("Channel", i, "closed")
		}
	}
}

package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	w := 25
	h := 6

	dat, err := ioutil.ReadFile("advent08.txt")
	if err != nil {
		panic(err)
	}
	dat = dat[:len(dat)-1]

	size := w * h
	layers := len(dat) / size

	digits := make([][]int, layers)
	for i := 0; i < layers; i++ {
		digits[i] = make([]int, 3)
	}
	for i := 0; i < len(dat); i++ {
		layer := i / size
		digits[layer][dat[i]-'0']++
	}

	digits0 := digits[0]
	for i := 1; i < layers; i++ {
		if digits[i][0] < digits0[0] {
			digits0 = digits[i]
		}
	}

	fmt.Println(digits0[1] * digits0[2])

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			for i := 0; i < layers; i++ {
				pos := i*size + y*w + x
				value := dat[pos]
				if value == '0' {
					fmt.Printf(".")
					break
				} else if value == '1' {
					fmt.Printf("#")
					break
				}
			}
		}
		fmt.Println()
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

/*

Fuel required to launch a given module is based on its mass. Specifically, to
find the fuel required for a module, take its mass, divide by three, round
down, and subtract 2.

*/
func Fuel(mass int) int {
	return mass/3 - 2
}

/*

...calculate its fuel and add it to the total. Then, treat the fuel amount you
just calculated as the input mass and repeat the process, continuing until a
fuel requirement is zero or negative.

*/
func Fuel2(mass int) int {
	sum := 0
	for {
		fuel := Fuel(mass)
		if fuel <= 0 {
			return sum
		}
		sum += fuel

		mass = Fuel(fuel)
		if mass <= 0 {
			return sum
		}
		sum += mass
	}
}

func main() {
	dat, err := ioutil.ReadFile("advent01.txt")
	if err != nil {
		panic(err)
	}

	txt := string(dat)
	lines := strings.Split(txt, "\n")
	strValues := lines[:len(lines)-1]
	sum := 0
	for _, value := range strValues {
		mass, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		fuel := Fuel(mass)
		sum += fuel
	}
	fmt.Println(sum)

	sum = 0
	for _, value := range strValues {
		mass, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		fuel := Fuel2(mass)
		sum += fuel
	}
	fmt.Println(sum)
}

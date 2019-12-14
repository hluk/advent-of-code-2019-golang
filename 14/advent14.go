//revive:disable:exported
package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type Component struct {
	amount int
	name   string
}

type Reaction struct {
	outputAmount int
	inputs       []Component
}

type ReactionMap = map[string]Reaction

func ParseReactionComponent(component string) Component {
	c := strings.Split(component, " ")
	amount, err := strconv.Atoi(c[0])
	if err != nil {
		panic(err)
	}
	return Component{amount, c[1]}
}

func LoadReactions(path string) ReactionMap {
	r := ReactionMap{}

	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	dat = dat[:len(dat)]

	txt := string(dat)
	txt = strings.TrimRight(txt, "\n")
	lines := strings.Split(txt, "\n")
	for _, line := range lines {
		lr := strings.Split(line, " => ")

		components := strings.Split(lr[0], ", ")
		inputs := make([]Component, len(components))
		for i, component := range components {
			inputs[i] = ParseReactionComponent(component)
		}

		output := ParseReactionComponent(lr[1])
		r[output.name] = Reaction{output.amount, inputs}
	}

	return r
}

func MinOre(component Component, r ReactionMap) int {
	need := map[string]int{component.name: component.amount}
	loop := true
	for loop {
		loop = false
		for name, amount := range need {
			if name != "ORE" && amount > 0 {
				rr := r[name]
				amount := (amount-1)/rr.outputAmount + 1
				need[name] -= rr.outputAmount * amount

				for _, c := range rr.inputs {
					need[c.name] += c.amount * amount
				}
				loop = true
			}
		}
	}
	return need["ORE"]
}

func MaxFuel(output Component, r ReactionMap) int {
	return sort.Search(output.amount, func(n int) bool {
		return MinOre(Component{n + 1, output.name}, r) > output.amount
	})
}

func main() {
	r := LoadReactions("advent14.txt")

	n1 := MinOre(Component{1, "FUEL"}, r)
	fmt.Println(n1)

	n2 := MaxFuel(Component{1000000000000, "FUEL"}, r)
	fmt.Println(n2)
}

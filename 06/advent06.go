//revive:disable:exported
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Orbits = map[string]string

func CreateOrbits(lines []string) Orbits {
	orbits := make(Orbits)
	for _, line := range lines {
		x := strings.Split(line, ")")
		parent := x[0]
		name := x[1]
		orbits[name] = parent
	}
	return orbits
}

func ForEachParent(name string, orbits Orbits, fn func(string)) int {
	count := 0
	parent, ok := orbits[name], true
	for ok {
		fn(parent)
		parent, ok = orbits[parent]
	}
	return count
}

func CountOrbitsFor(name string, orbits Orbits) int {
	count := 0
	ForEachParent(name, orbits, func(_ string) { count++ })
	return count
}

func CountOrbits(orbits Orbits) int {
	count := 0
	for name := range orbits {
		count += CountOrbitsFor(name, orbits)
	}
	return count
}

func Parents(name string, orbits Orbits) []string {
	parents := []string{}
	ForEachParent(name, orbits, func(parent string) {
		parents = append(parents, parent)
	})
	return parents
}

func IndexOf(a string, arr []string) int {
	for i, x := range arr {
		if x == a {
			return i
		}
	}
	return -1
}

func CommonParentDistance(a, b string, orbits Orbits) (int, int) {
	parentsA := Parents(a, orbits)
	parentsB := Parents(b, orbits)
	for i, x := range parentsA {
		j := IndexOf(x, parentsB)
		if j != -1 {
			return i, j
		}
	}
	panic("No common parent found")
}

func CountTransfers(from, to string, orbits Orbits) int {
	i, j := CommonParentDistance(from, to, orbits)
	return i + j
}

func LoadOrbits(path string) Orbits {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	txt := string(dat)
	txt = strings.TrimRight(txt, "\n")
	lines := strings.Split(txt, "\n")
	return CreateOrbits(lines)
}

func main() {
	orbits := LoadOrbits("advent06.txt")
	fmt.Println(CountOrbits(orbits))
	fmt.Println(CountTransfers("YOU", "SAN", orbits))
}

package main

import (
	"fmt"
	"testing"
)

func ExamplePermutations() {
	p := Permutations([]int{1, 2})
	fmt.Println(p())
	fmt.Println(p())
	fmt.Println(p())
	// Output:
	// [1 2]
	// [2 1]
	// []
}

func BenchmarkPermutations(t *testing.B) {
	p := Permutations([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	for {
		pp := p()
		if pp == nil {
			break
		}
	}
}

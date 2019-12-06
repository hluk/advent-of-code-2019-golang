package main

import (
	"fmt"
	"testing"
)

func ExampleLoadOrbits() {
	orbits := LoadOrbits("advent06_test1.txt")
	fmt.Println(orbits["B"])
	fmt.Println(orbits["C"])
	fmt.Println(orbits["D"])
	fmt.Println(orbits["I"])
	fmt.Println(orbits["L"])
	// Output:
	// COM
	// B
	// C
	// D
	// K
}

func TestLoadOrbits(t *testing.T) {
	orbits := LoadOrbits("advent06_test1.txt")
	actual := orbits["B"]
	expected := "COM"
	if actual != expected {
		t.Errorf("Actual %v; Expected: %v", actual, expected)
	}
}

func TestCountOrbitsFor(t *testing.T) {
	orbits := LoadOrbits("advent06_test1.txt")

	actual := CountOrbitsFor("D", orbits)
	expected := 3
	if actual != expected {
		t.Errorf("Actual %v; Expected: %v", actual, expected)
	}

	actual = CountOrbitsFor("L", orbits)
	expected = 7
	if actual != expected {
		t.Errorf("Actual %v; Expected: %v", actual, expected)
	}
}

func TestCountOrbits(t *testing.T) {
	orbits := LoadOrbits("advent06_test1.txt")
	actual := CountOrbits(orbits)
	expected := 42
	if actual != expected {
		t.Errorf("Actual %v; Expected: %v", actual, expected)
	}
}

func TestCommonParentDistance(t *testing.T) {
	orbits := LoadOrbits("advent06_test2.txt")
	actual1, actual2 := CommonParentDistance("YOU", "SAN", orbits)
	expected1, expected2 := 3, 1
	if actual1 != expected1 || actual2 != expected2 {
		t.Errorf("Actual %v,%v; Expected: %v,%v", actual1, actual2, expected1, expected2)
	}
}

func TestPart2(t *testing.T) {
	orbits := LoadOrbits("advent06_test2.txt")
	actual := CountTransfers("YOU", "SAN", orbits)
	expected := 4
	if actual != expected {
		t.Errorf("Actual %v; Expected: %v", actual, expected)
	}
}

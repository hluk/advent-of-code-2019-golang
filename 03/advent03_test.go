package main

import (
	"reflect"
	"sort"
	"testing"
)

func sortCoords(slice [][2]int) {
	sort.Slice(slice, func(i, j int) bool {
		a, b := slice[i], slice[j]
		if a[0] == b[0] {
			return a[1] < b[1]
		}
		return a[0] < b[0]
	})
}

func TestLoadWire(t *testing.T) {
	w := LoadWire("R8,U5,L5,D3")
	if 0 == w.Get(1, 0) {
		t.Errorf("Failed")
	}
	if 0 == w.Get(2, 0) {
		t.Errorf("Failed")
	}
	if 0 == w.Get(7, 0) {
		t.Errorf("Failed")
	}
	if 0 == w.Get(8, 0) {
		t.Errorf("Failed")
	}
	if 0 != w.Get(9, 0) {
		t.Errorf("Failed")
	}

	if 0 == w.Get(8, 1) {
		t.Errorf("Failed")
	}
	if 0 == w.Get(8, 5) {
		t.Errorf("Failed")
	}
	if 0 != w.Get(8, 6) {
		t.Errorf("Failed")
	}
}

func TestCrossings(t *testing.T) {
	w1 := LoadWire("R8,U5,L5,D3")
	w2 := LoadWire("U7,R6,D4,L4")
	got := Crossings(w1, w2)
	expected := [][2]int{{6, 5}, {3, 3}}
	sortCoords(got)
	sortCoords(expected)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Actual: %v; Expected: %v", got, expected)
	}
}

func TestCrossingsShortestDistance(t *testing.T) {
	w1 := LoadWire("R8,U5,L5,D3")
	w2 := LoadWire("U7,R6,D4,L4")
	got := CrossingsShortestDistance(w1, w2)
	expected := 6
	if got != expected {
		t.Errorf("Actual: %v; Expected: %v", got, expected)
	}

	w1 = LoadWire("R75,D30,R83,U83,L12,D49,R71,U7,L72")
	w2 = LoadWire("U62,R66,U55,R34,D71,R55,D58,R83")
	got = CrossingsShortestDistance(w1, w2)
	expected = 159
	if got != expected {
		t.Errorf("Actual: %v; Expected: %v", got, expected)
	}

	w1 = LoadWire("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51")
	w2 = LoadWire("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
	got = CrossingsShortestDistance(w1, w2)
	expected = 135
	if got != expected {
		t.Errorf("Actual: %v; Expected: %v", got, expected)
	}
}

func TestCrossingsShortestSignal(t *testing.T) {
	w1 := LoadWire("R8,U5,L5,D3")
	w2 := LoadWire("U7,R6,D4,L4")
	got := CrossingsShortestSignal(w1, w2)
	expected := 30
	if got != expected {
		t.Errorf("Actual: %v; Expected: %v", got, expected)
	}

	w1 = LoadWire("R75,D30,R83,U83,L12,D49,R71,U7,L72")
	w2 = LoadWire("U62,R66,U55,R34,D71,R55,D58,R83")
	got = CrossingsShortestSignal(w1, w2)
	expected = 610
	if got != expected {
		t.Errorf("Actual: %v; Expected: %v", got, expected)
	}

	w1 = LoadWire("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51")
	w2 = LoadWire("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
	got = CrossingsShortestSignal(w1, w2)
	expected = 410
	if got != expected {
		t.Errorf("Actual: %v; Expected: %v", got, expected)
	}
}

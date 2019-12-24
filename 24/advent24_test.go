package main

import (
	"testing"

	"github.com/hluk/advent-of-code-2019-golang"
)

func TestAreaBug(t *testing.T) {
	var a Area

	adv.AssertEq(a.HasBug(0, 0), false, t)

	adv.AssertEq(a.HasBug(2, 2), false, t)

	a = a.Bug(0, 0)
	adv.AssertEq(a.HasBug(0, 0), true, t)

	adv.AssertEq(a.HasBug(2, 2), false, t)
}

func TestAdjecentBugs(t *testing.T) {
	var a Area

	adv.AssertEq(a.AdjecentBugs(0, 0), 0, t)
	adv.AssertEq(a.AdjecentBugs(1, 1), 0, t)

	a = a.Bug(0, 0)
	adv.AssertEq(a.AdjecentBugs(0, 0), 0, t)
	adv.AssertEq(a.AdjecentBugs(1, 0), 1, t)
	adv.AssertEq(a.AdjecentBugs(1, 1), 0, t)
	adv.AssertEq(a.AdjecentBugs(1, 1), 0, t)
}

func TestAreaRating(t *testing.T) {
	var a Area

	adv.AssertEq(a, Area(0), t)

	a = a.Bug(0, 3)
	adv.AssertEq(a, Area(32768), t)
	a = a.Bug(1, 4)
	adv.AssertEq(a, Area(2129920), t)
}

func TestLoadArea(t *testing.T) {
	a := LoadArea("advent24_test1.txt")

	if a.HasBug(0, 0) {
		t.Errorf("0,0 not expected to have a bug")
	}
	if !a.HasBug(4, 0) {
		t.Errorf("1,0 expected to have a bug")
	}

	if !a.HasBug(0, 1) {
		t.Errorf("0,1 expected to have a bug")
	}
	if !a.HasBug(3, 1) {
		t.Errorf("3,1 expected to have a bug")
	}
	if a.HasBug(4, 1) {
		t.Errorf("4,1 expected to have a bug")
	}

	if a.AdjecentBugs(1, 0) != 0 {
		t.Errorf("1,0 not expected to have any adjacent bugs, actual: %d", a.AdjecentBugs(1, 0))
	}
}

func TestFirstSeen(t *testing.T) {
	a := LoadArea("advent24_test1.txt")
	a = FirstSeen(a)
	if !a.HasBug(0, 3) {
		t.Errorf("0,3 expected to have a bug")
	}
	if !a.HasBug(1, 4) {
		t.Errorf("1,4 expected to have a bug")
	}
	adv.AssertEq(a, Area(2129920), t)
}

func TestCountAfter(t *testing.T) {
	a := LoadArea("advent24_test1.txt")
	n := CountAfter(10, a)
	if n != 99 {
		t.Errorf("Expected 99 bugs, actual: %d", n)
	}
}

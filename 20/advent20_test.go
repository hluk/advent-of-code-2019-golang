package main

import (
	"testing"

	"github.com/hluk/advent-of-code-2019-golang"
)

func TestLoadArea_1(t *testing.T) {
	a, p := LoadArea("advent20_test1.txt")

	adv.AssertEq(p[Tile{'A', 'A'}][0], Pos{9, 1}, t)
	adv.AssertEq(p[Tile{'A', 'A'}][1], Pos{}, t)

	adv.AssertEq(p[Tile{'Z', 'Z'}][0], Pos{13, 17}, t)
	adv.AssertEq(p[Tile{'Z', 'Z'}][1], Pos{}, t)

	adv.AssertEq(a[Pos{9, 1}], Tile{'A', 'A'}, t)
	adv.AssertEq(a[Pos{9, 2}], Floor(), t)
	adv.AssertEq(a[Pos{9, 3}], Floor(), t)

	fg := Tile{'F', 'G'}
	adv.AssertEq(p[fg][0], Pos{11, 11}, t)
	adv.AssertEq(p[fg][1], Pos{1, 15}, t)

	adv.AssertEq(a[Pos{11, 11}], fg, t)
	adv.AssertEq(a[Pos{1, 15}], fg, t)
}

func TestLoadArea_2(t *testing.T) {
	a, p := LoadArea("advent20_test2.txt")

	qg := Tile{'Q', 'G'}
	adv.AssertEq(p[qg][0], Pos{25, 17}, t)
	adv.AssertEq(p[qg][1], Pos{33, 23}, t)

	adv.AssertEq(a[Pos{25, 17}], qg, t)
	adv.AssertEq(a[Pos{33, 23}], qg, t)
}

func TestExplore_1(t *testing.T) {
	a, p := LoadArea("advent20_test1.txt")
	steps := Explore(a, p, false)
	adv.AssertEq(steps, 23, t)
}

func TestExplore_2(t *testing.T) {
	a, p := LoadArea("advent20_test2.txt")
	steps := Explore(a, p, false)
	adv.AssertEq(steps, 58, t)
}

func TestExplore2_1(t *testing.T) {
	a, p := LoadArea("advent20_test1.txt")
	steps := Explore(a, p, true)
	adv.AssertEq(steps, 26, t)
}

func TestExplore2_3(t *testing.T) {
	a, p := LoadArea("advent20_test3.txt")
	steps := Explore(a, p, true)
	adv.AssertEq(steps, 396, t)
}

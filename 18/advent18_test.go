package main

import (
	"testing"

	"github.com/hluk/advent-of-code-2019-golang"
)

func TestKeys(t *testing.T) {
	k := Keys{}
	adv.AssertEq(k.ToString(), "", t)

	adv.AssertEq(k.Unlock('A'), false, t)
	k.Add('a')
	adv.AssertEq(k.ToString(), "a", t)
	adv.AssertEq(k.Unlock('A'), true, t)

	adv.AssertEq(k.Unlock('Z'), false, t)
	k.Add('z')
	adv.AssertEq(k.ToString(), "az", t)
	adv.AssertEq(k.Unlock('Z'), true, t)

	k.Add('x')
	adv.AssertEq(k.ToString(), "axz", t)
	adv.AssertEq(k.Unlock('A'), true, t)
	adv.AssertEq(k.Unlock('X'), true, t)
	adv.AssertEq(k.Unlock('Z'), true, t)
	adv.AssertEq(k.Unlock('B'), false, t)
}

func TestExplore4_1(t *testing.T) {
	a, keys, starts := LoadArea("advent18_test1.txt")
	starts = Split(a, starts[0])
	steps := Explore4(a, keys, starts)
	adv.AssertEq(steps, 8, t)
}

func TestExplore4_2(t *testing.T) {
	a, keys, starts := LoadArea("advent18_test2.txt")
	steps := Explore4(a, keys, starts)
	adv.AssertEq(steps, 24, t)
}

func TestExplore4_3(t *testing.T) {
	a, keys, starts := LoadArea("advent18_test3.txt")
	steps := Explore4(a, keys, starts)
	adv.AssertEq(steps, 32, t)
}

func TestExplore4_4(t *testing.T) {
	a, keys, starts := LoadArea("advent18_test4.txt")
	steps := Explore4(a, keys, starts)
	adv.AssertEq(steps, 72, t)
}

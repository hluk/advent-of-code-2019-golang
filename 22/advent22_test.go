package main

import (
	"testing"

	"github.com/hluk/advent-of-code-2019-golang"
)

func TestCreateDeck(t *testing.T) {
	d := CreateDeck(10)
	adv.AssertEq(d.Cards, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, t)
}

func TestIndexOf(t *testing.T) {
	d := Deck{[]int{0, 7, 4, 1, 8, 5, 2, 9, 6, 3}}
	adv.AssertEq(d.IndexOf(4), 2, t)
	adv.AssertEq(d.IndexOf(3), 9, t)
}

func TestDeal(t *testing.T) {
	d := CreateDeck(10)
	d = d.Deal()
	adv.AssertEq(d.Cards, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, t)
}

func TestCutN(t *testing.T) {
	d := CreateDeck(10)
	d = d.CutN(3)
	adv.AssertEq(d.Cards, []int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2}, t)
}

func TestCutNNegative(t *testing.T) {
	d := CreateDeck(10)
	d = d.CutN(-4)
	adv.AssertEq(d.Cards, []int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5}, t)
}

func TestDealN(t *testing.T) {
	d := CreateDeck(10)
	d = d.DealN(3)
	adv.AssertEq(d.Cards, []int{0, 7, 4, 1, 8, 5, 2, 9, 6, 3}, t)
}

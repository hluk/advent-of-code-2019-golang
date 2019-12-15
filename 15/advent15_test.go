package main

import (
	"testing"

	"github.com/hluk/advent-of-code-2019-golang"
)

func fillTime0(area Area) int {
	return fillTime(area, Pos{0, 0})
}

func TestNext(t *testing.T) {
	d := Pos{0, -1}

	d = next(d)
	adv.AssertEq(d, Pos{1, 0}, t)

	d = next(d)
	adv.AssertEq(d, Pos{0, 1}, t)

	d = next(d)
	adv.AssertEq(d, Pos{-1, 0}, t)

	d = next(d)
	adv.AssertEq(d, Pos{0, -1}, t)
}

func TestFill0(t *testing.T) {
	area := Area{
		{0, 0}: '.',
	}

	adv.AssertEq(fillTime0(area), 0, t)
}

func TestFill1(t *testing.T) {
	area := Area{
		{0, 0}:  '.',
		{1, 0}:  '.',
		{-1, 0}: '.',
	}

	adv.AssertEq(fillTime0(area), 1, t)
}

func TestFillT(t *testing.T) {
	/*
			###
		   #...#
			#.#
			#O#
			 #
	*/
	area := Area{
		{0, 0}: '.',
		{0, 1}: '.',
		{0, 2}: '.',

		{-1, 2}: '.',
		{1, 2}:  '.',
	}

	adv.AssertEq(fillTime0(area), 3, t)
}

func TestFillO(t *testing.T) {
	/*
		#####
		#...#
		#.#.#
		#O..#
		#####
	*/
	area := Area{
		{0, 0}: '.',
		{0, 1}: '.',
		{0, 2}: '.',

		{2, 0}: '.',
		{2, 1}: '.',
		{2, 2}: '.',

		{1, 0}: '.',
		{1, 2}: '.',
	}

	adv.AssertEq(fillTime0(area), 4, t)
}

func TestFill(t *testing.T) {
	/*
		 ##
		#..##
		#.#..#
		#.O.#
		 ###
	*/
	area := Area{
		{0, 0}: '.',
		{1, 0}: '.',
		{1, 1}: '.',
		{2, 1}: '.',

		{-1, 0}: '.',
		{-1, 1}: '.',
		{-1, 2}: '.',
		{0, 2}:  '.',
	}

	//pos := Pos{0, 0}
	//d := Pos{1, 0}
	//printArea(area, pos, d, pos)
	adv.AssertEq(fillTime0(area), 4, t)
	//printArea(area, pos, d, pos)
}

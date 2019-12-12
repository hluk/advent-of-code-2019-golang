package main

import "testing"

func TestLoadMoons(t *testing.T) {
	moons := LoadMoons("advent12_test.txt")

	actual := moons.a.P
	expected := Pos{-1, 0, 2}
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

	actual = moons.a.V
	expected = Pos{0, 0, 0}
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

	actual = moons.d.P
	expected = Pos{3, 5, -1}
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestSimulate1(t *testing.T) {
	moons := LoadMoons("advent12_test.txt")

	moons.Simulate()
	/*
		pos=<x= 2, y=-1, z= 1>, vel=<x= 3, y=-1, z=-1>
		pos=<x= 3, y=-7, z=-4>, vel=<x= 1, y= 3, z= 3>
		pos=<x= 1, y=-7, z= 5>, vel=<x=-3, y= 1, z=-3>
		pos=<x= 2, y= 2, z= 0>, vel=<x=-1, y=-3, z= 1>
	*/

	{
		actual := moons.a.V
		expected := Pos{3, -1, -1}
		if actual != expected {
			t.Errorf("Expected: %v, Actual: %v", expected, actual)
		}
	}

	{
		actual := moons.a.P
		expected := Pos{2, -1, 1}
		if actual != expected {
			t.Errorf("Expected: %v, Actual: %v", expected, actual)
		}
	}
}

func TestSimulate10(t *testing.T) {
	moons := LoadMoons("advent12_test.txt")

	for i := 0; i < 10; i++ {
		moons.Simulate()
	}
	/*
		pos=<x= 2, y= 1, z=-3>, vel=<x=-3, y=-2, z= 1>
		pos=<x= 1, y=-8, z= 0>, vel=<x=-1, y= 1, z= 3>
		pos=<x= 3, y=-6, z= 1>, vel=<x= 3, y= 2, z=-3>
		pos=<x= 2, y= 0, z= 4>, vel=<x= 1, y=-1, z=-1>
	*/

	{
		actual := moons.a.V
		expected := Pos{-3, -2, 1}
		if actual != expected {
			t.Errorf("Expected: %v, Actual: %v", expected, actual)
		}
	}

	{
		actual := moons.a.P
		expected := Pos{2, 1, -3}
		if actual != expected {
			t.Errorf("Expected: %v, Actual: %v", expected, actual)
		}
	}
}

func TestTotalEnergy(t *testing.T) {
	moons := LoadMoons("advent12_test.txt")

	for i := 0; i < 10; i++ {
		moons.Simulate()
	}

	{
		actual := moons.c.PotentialEnergy()
		expected := 10
		if actual != expected {
			t.Errorf("Expected: %v, Actual: %v", expected, actual)
		}
	}

	{
		actual := moons.c.KineticEnergy()
		expected := 8
		if actual != expected {
			t.Errorf("Expected: %v, Actual: %v", expected, actual)
		}
	}

	{
		actual := moons.TotalEnergy()
		expected := 179
		if actual != expected {
			t.Errorf("Expected: %v, Actual: %v", expected, actual)
		}
	}
}

func TestFindLoop(t *testing.T) {
	{
		moons := LoadMoons("advent12_test.txt")
		expected := 2772
		actual := moons.FindLoop()
		if actual != expected {
			t.Errorf("Expected: %v, Actual: %v", expected, actual)
		}
	}

	{
		moons := LoadMoons("advent12_test2.txt")
		expected := 4686774924
		actual := moons.FindLoop()
		if actual != expected {
			t.Errorf("Expected: %v, Actual: %v", expected, actual)
		}
	}
}

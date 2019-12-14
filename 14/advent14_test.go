package main

import "testing"

func TestProduce1(t *testing.T) {
	r := LoadReactions("advent14_test1.txt")

	actual := MinOre(Component{2, "A"}, r)
	expected := 9
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

	actual = MinOre(Component{1, "A"}, r)
	expected = 9
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

	actual = MinOre(Component{10, "A"}, r)
	expected = 45
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestProduce2(t *testing.T) {
	/*
		9 ORE => 2 A
		8 ORE => 3 B
		7 ORE => 5 C
	*/
	r := LoadReactions("advent14_test1.txt")

	// 3 A, 4 B
	actual := MinOre(Component{2, "AB"}, r)
	expected := 3*9 + 3*8
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

	// 15 B, 21 C
	actual = MinOre(Component{3, "BC"}, r)
	expected = 5*8 + 5*7
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

	// 16 C, 4 A
	actual = MinOre(Component{4, "CA"}, r)
	expected = 4*7 + 2*9
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestProduce3(t *testing.T) {
	r := LoadReactions("advent14_test1.txt")

	actual := MinOre(Component{1, "FUEL"}, r)
	expected := 165
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestProduce4(t *testing.T) {
	r := LoadReactions("advent14_test2.txt")

	actual := MinOre(Component{1, "FUEL"}, r)
	expected := 13312
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestProduce5(t *testing.T) {
	r := LoadReactions("advent14_test3.txt")

	actual := MinOre(Component{1, "FUEL"}, r)
	expected := 180697
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestMaxFuel1(t *testing.T) {
	r := LoadReactions("advent14_test2.txt")

	actual := MaxFuel(Component{1000000000000, "FUEL"}, r)
	expected := 82892753
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

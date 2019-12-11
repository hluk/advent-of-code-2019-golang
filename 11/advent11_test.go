package main

import "testing"

func TestTurnLeft(t *testing.T) {
	actual := TurnLeft(Pos{0, -1})
	expected := Pos{-1, 0}
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

	actual = TurnLeft(Pos{-1, 0})
	expected = Pos{0, 1}
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

	actual = TurnLeft(Pos{0, 1})
	expected = Pos{1, 0}
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

	actual = TurnLeft(Pos{1, 0})
	expected = Pos{0, -1}
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestTurnRight(t *testing.T) {
	actual := TurnRight(Pos{0, -1})
	expected := Pos{1, 0}
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

	actual = TurnRight(Pos{1, 0})
	expected = Pos{0, 1}
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

	actual = TurnRight(Pos{0, 1})
	expected = Pos{-1, 0}
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

	actual = TurnRight(Pos{-1, 0})
	expected = Pos{0, -1}
	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

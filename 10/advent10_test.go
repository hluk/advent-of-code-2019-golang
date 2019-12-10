package main

import "testing"

func TestAngle(t *testing.T) {
	a := Angle(Pos{0, -1})
	expected := 0
	if int(a) != 0 {
		t.Errorf("Expected: %v, Actual: %v", expected, a)
	}

	a = Angle(Pos{1, -1})
	expected = 45
	if int(a) != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, a)
	}

	a = Angle(Pos{1, 0})
	expected = 90
	if int(a) != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, a)
	}

	a = Angle(Pos{0, 1})
	expected = 45
	if int(a) != 180 {
		t.Errorf("Expected: %v, Actual: %v", expected, a)
	}

	a = Angle(Pos{-1, 0})
	expected = 270
	if int(a) != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, a)
	}

	a = Angle(Pos{-1, -1})
	expected = 270 + 45
	if int(a) != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, a)
	}
}

func TestLoadMap(t *testing.T) {
	m := LoadMap("advent10_test1.txt")
	if m.W != 5 {
		t.Errorf("FAILED")
	}
	if m.H != 5 {
		t.Errorf("FAILED")
	}
	if m.Get(Pos{0, 0}) {
		t.Errorf("FAILED")
	}
	if !m.Get(Pos{1, 0}) {
		t.Errorf("FAILED")
	}
	if !m.Get(Pos{3, 4}) {
		t.Errorf("FAILED")
	}
}

func TestVisibleAsteroidInLOS(t *testing.T) {
	m := LoadMap("advent10_test1.txt")
	p := m.VisibleAsteroidInLOS(Pos{1, 0}, Pos{1, 0})
	expected := Pos{4, 0}
	if p != expected {
		t.Errorf("FAILED")
	}
}

func TestVisibleAsteroids(t *testing.T) {
	m := LoadMap("advent10_test1.txt")
	v := m.VisibleAsteroids(Pos{1, 0})
	if v != 7 {
		t.Errorf("FAILED")
	}
	v = m.VisibleAsteroids(Pos{3, 4})
	if v != 8 {
		t.Errorf("FAILED")
	}
}

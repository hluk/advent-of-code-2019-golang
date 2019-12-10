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

func TestHits(t *testing.T) {
	m := LoadMap("advent10_test1.txt")

	v := m.Hits(Pos{1, 0})
	if len(v) != 7 {
		t.Errorf("FAILED")
	}

	v = m.Hits(Pos{3, 4})
	if len(v) != 8 {
		t.Errorf("FAILED")
	}
}

func TestHitAt(t *testing.T) {
	m := LoadMap("advent10_test1.txt")
	p0 := Pos{1, 0}
	hits := m.Hits(p0)

	p := HitAt(hits, p0, 1)
	expected := Pos{4, 0}
	if p != expected {
		t.Errorf("FAILED")
	}

	p = HitAt(hits, p0, 2)
	expected = Pos{4, 2}
	if p != expected {
		t.Errorf("FAILED")
	}

	p = HitAt(hits, p0, 3)
	expected = Pos{3, 2}
	if p != expected {
		t.Errorf("FAILED")
	}
}

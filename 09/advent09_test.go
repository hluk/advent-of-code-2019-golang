package main

import "testing"

func TestQuine(t *testing.T) {
	r := []Value{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	chOut := make(chan Value)
	go Run(r, nil, chOut)
	for i := 0; i < len(r); i++ {
		if r[i] != <-chOut {
			t.Errorf("Unexpeted output value")
		}
	}
	_, ok := <-chOut
	if ok {
		t.Errorf("Expected end of output")
	}
}

func TestOutput16DigitNumber(t *testing.T) {
	r := []Value{1102, 34915192, 34915192, 7, 4, 7, 99, 0}
	chOut := make(chan Value, len(r))
	Run(r, nil, chOut)
	v := <-chOut
	if v < 1e15 || v >= 1e16 {
		t.Errorf("Failed")
	}
}

func TestOutputLarge(t *testing.T) {
	r := []Value{104, 1125899906842624, 99}
	chOut := make(chan Value, len(r))
	Run(r, nil, chOut)
	if 1125899906842624 != <-chOut {
		t.Errorf("Failed")
	}
}

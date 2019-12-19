package intcode

import "testing"

func RunInOut(r []Value, in Value) []Value {
	chIn := make(chan Value)
	chOut := make(chan Value)
	go Run(r, chIn, chOut)
	chIn <- in

	outputs := []Value{}
	for v, ok := <-chOut; ok; v, ok = <-chOut {
		outputs = append(outputs, v)
	}
	return outputs
}

func TestLoadOp(t *testing.T) {
	fixtures := map[Value]Op{
		1002: Op{2, 0, 1, 0},
	}
	for value := range fixtures {
		expected := fixtures[value]
		op := LoadOp(value)
		if op != expected {
			t.Errorf("Actual %v; Expected: %v", op, expected)
		}
	}
}

func TestInputOutput(t *testing.T) {
	r := []Value{3, 0, 4, 0, 99}
	outputs := RunInOut(r, 666)
	if len(outputs) != 1 || outputs[0] != 666 {
		t.Errorf("Expected out == 666, but out = %v", outputs)
	}
}

func TestEq8(t *testing.T) {
	r := []Value{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}

	outputs := RunInOut(r, 666)
	if len(outputs) != 1 || outputs[0] != 0 {
		t.Errorf("Expected outputs == [0], but outputs = %v", outputs)
	}

	outputs = RunInOut(r, 8)
	if len(outputs) != 1 || outputs[0] != 1 {
		t.Errorf("Expected outputs == [1], but outputs = %v", outputs)
	}
}

func TestLt8(t *testing.T) {
	r := []Value{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}

	outputs := RunInOut(r, 8)
	if len(outputs) != 1 || outputs[0] != 0 {
		t.Errorf("Expected outputs == [0], but outputs = %v", outputs)
	}

	outputs = RunInOut(r, 7)
	if len(outputs) != 1 || outputs[0] != 1 {
		t.Errorf("Expected outputs == [1], but outputs = %v", outputs)
	}
}

func TestEq8Im(t *testing.T) {
	r := []Value{3, 3, 1108, -1, 8, 3, 4, 3, 99}

	outputs := RunInOut(r, 666)
	if len(outputs) != 1 || outputs[0] != 0 {
		t.Errorf("Expected outputs == [0], but outputs = %v", outputs)
	}

	outputs = RunInOut(r, 8)
	if len(outputs) != 1 || outputs[0] != 1 {
		t.Errorf("Expected outputs == [1], but outputs = %v", outputs)
	}
}

func TestLt8Im(t *testing.T) {
	r := []Value{3, 3, 1107, -1, 8, 3, 4, 3, 99}

	outputs := RunInOut(r, 8)
	if len(outputs) != 1 || outputs[0] != 0 {
		t.Errorf("Expected outputs == [0], but outputs = %v", outputs)
	}

	outputs = RunInOut(r, 7)
	if len(outputs) != 1 || outputs[0] != 1 {
		t.Errorf("Expected outputs == [1], but outputs = %v", outputs)
	}
}

func TestJump(t *testing.T) {
	r := []Value{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}

	outputs := RunInOut(r, 0)
	if len(outputs) != 1 || outputs[0] != 0 {
		t.Errorf("Expected outputs == [0], but outputs = %v", outputs)
	}

	outputs = RunInOut(r, -1)
	if len(outputs) != 1 || outputs[0] != 1 {
		t.Errorf("Expected outputs == [1], but outputs = %v", outputs)
	}
}

func TestJumpIm(t *testing.T) {
	r := []Value{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}

	outputs := RunInOut(r, 0)
	if len(outputs) != 1 || outputs[0] != 0 {
		t.Errorf("Expected outputs == [0], but outputs = %v", outputs)
	}

	outputs = RunInOut(r, -1)
	if len(outputs) != 1 || outputs[0] != 1 {
		t.Errorf("Expected outputs == [1], but outputs = %v", outputs)
	}
}

func TestBigger(t *testing.T) {
	r := []Value{
		3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}

	outputs := RunInOut(r, 7)
	if len(outputs) != 1 || outputs[0] != 999 {
		t.Errorf("Expected outputs == [999], but outputs = %v", outputs)
	}

	outputs = RunInOut(r, 8)
	if len(outputs) != 1 || outputs[0] != 1000 {
		t.Errorf("Expected outputs == [1000], but outputs = %v", outputs)
	}

	outputs = RunInOut(r, 9)
	if len(outputs) != 1 || outputs[0] != 1001 {
		t.Errorf("Expected outputs == [1001], but outputs = %v", outputs)
	}
}

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

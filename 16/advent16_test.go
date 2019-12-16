package main

import (
	"testing"

	"github.com/hluk/advent-of-code-2019-golang"
)

func TestBasePattern1(t *testing.T) {
	adv.AssertEq(BasePattern(0, 0), 1, t)
	adv.AssertEq(BasePattern(0, 1), 0, t)
	adv.AssertEq(BasePattern(0, 2), -1, t)
	adv.AssertEq(BasePattern(0, 3), 0, t)
}

func TestBasePattern2(t *testing.T) {
	adv.AssertEq(BasePattern(1, 0), 0, t)
	adv.AssertEq(BasePattern(1, 1), 1, t)
	adv.AssertEq(BasePattern(1, 2), 1, t)
	adv.AssertEq(BasePattern(1, 3), 0, t)
	adv.AssertEq(BasePattern(1, 4), 0, t)
	adv.AssertEq(BasePattern(1, 5), -1, t)
	adv.AssertEq(BasePattern(1, 6), -1, t)
	adv.AssertEq(BasePattern(1, 7), 0, t)
	adv.AssertEq(BasePattern(1, 8), 0, t)
	adv.AssertEq(BasePattern(1, 9), 1, t)
}

func TestBasePattern3(t *testing.T) {
	adv.AssertEq(BasePattern(2, 0), 0, t)
	adv.AssertEq(BasePattern(2, 1), 0, t)
	adv.AssertEq(BasePattern(2, 2), 1, t)
	adv.AssertEq(BasePattern(2, 3), 1, t)
	adv.AssertEq(BasePattern(2, 4), 1, t)
	adv.AssertEq(BasePattern(2, 5), 0, t)
	adv.AssertEq(BasePattern(2, 6), 0, t)
	adv.AssertEq(BasePattern(2, 7), 0, t)
	adv.AssertEq(BasePattern(2, 8), -1, t)
	adv.AssertEq(BasePattern(2, 9), -1, t)
}

func TestSkip(t *testing.T) {
	adv.AssertEq(Step(0, 0), 2, t)
	adv.AssertEq(Step(0, 2), 2, t)
	adv.AssertEq(Step(0, 4), 2, t)

	adv.AssertEq(Step(1, 1), 1, t)
	adv.AssertEq(Step(1, 2), 3, t)
	adv.AssertEq(Step(1, 5), 1, t)
	adv.AssertEq(Step(1, 6), 3, t)

	adv.AssertEq(Step(3, 3), 1, t)
	adv.AssertEq(Step(3, 4), 1, t)
	adv.AssertEq(Step(3, 5), 1, t)
	adv.AssertEq(Step(3, 6), 5, t)
}

func TestFFT(t *testing.T) {
	inp := ParseInput("12345678")
	out := ParseInput("48226158")
	adv.AssertEq(FFT(inp, 1), out, t)

	inp = out
	out = ParseInput("34040438")
	adv.AssertEq(FFT(inp, 1), out, t)

	inp = ParseInput("12345678")
	out = ParseInput("34040438")
	adv.AssertEq(FFT(inp, 2), out, t)
}

func TestFFT_1(t *testing.T) {
	inp := ParseInput("80871224585914546619083218645595")
	out := ParseInput("24176176")
	adv.AssertEq(FFT(inp, 100)[:8], out, t)
}

func TestFFT_2(t *testing.T) {
	inp := ParseInput("19617804207202209144916044189917")
	out := ParseInput("73745418")
	adv.AssertEq(FFT(inp, 100)[:8], out, t)
}

func TestFFT_3(t *testing.T) {
	inp := ParseInput("69317163492948606335995924319873")
	out := ParseInput("52432133")
	adv.AssertEq(FFT(inp, 100)[:8], out, t)
}

func TestToNumber(t *testing.T) {
	inp := ParseInput("12345678")
	adv.AssertEq(ToNumber(1, inp), 1, t)
	adv.AssertEq(ToNumber(2, inp), 12, t)
	adv.AssertEq(ToNumber(3, inp), 123, t)
	adv.AssertEq(ToNumber(4, inp), 1234, t)
	adv.AssertEq(ToNumber(5, inp), 12345, t)
	adv.AssertEq(ToNumber(6, inp), 123456, t)
}

func TestPart2(t *testing.T) {
	inp := Repeated(ParseInput("03036732577212944063491565474664"))

	offset := Offset(inp)
	adv.AssertEq(offset, 303673, t)

	expected := ParseInput("84462026")
	actual := FFT2(inp, 100, offset, offset+8)
	adv.AssertEq(actual, expected, t)
	adv.AssertEq(ToNumber(8, actual), 84462026, t)
}

func TestPart2_1(t *testing.T) {
	inp := Repeated(ParseInput("02935109699940807407585447034323"))

	offset := Offset(inp)
	adv.AssertEq(offset, 293510, t)

	expected := ParseInput("78725270")
	actual := FFT2(inp, 100, offset, offset+8)
	adv.AssertEq(actual, expected, t)
}

func TestPart2_2(t *testing.T) {
	inp := Repeated(ParseInput("03081770884921959731165446850517"))

	offset := Offset(inp)
	adv.AssertEq(offset, 308177, t)

	expected := ParseInput("53553731")
	actual := FFT2(inp, 100, offset, offset+8)
	adv.AssertEq(actual, expected, t)
}

func TestRepeated(t *testing.T) {
	a := []byte{1, 2, 3}
	a = Repeated(a)
	adv.AssertEq(len(a), 3*10000, t)
	adv.AssertEq(a[:3], []byte{1, 2, 3}, t)
	adv.AssertEq(a[3:6], []byte{1, 2, 3}, t)
	adv.AssertEq(a[10000*3-3:], []byte{1, 2, 3}, t)
}

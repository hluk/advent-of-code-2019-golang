package main

import (
	"fmt"
	"github.com/hluk/advent-of-code-2019-golang"
)

func BasePattern(i int, j int) int {
	k := ((j + 1) / (i + 1)) & 3
	switch k {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 0
	case 3:
		return -1
	}
	panic("X")
}

func Step(i int, j int) int {
	k := ((j + 2) / (i + 1)) & 3
	if k == 0 || k == 2 {
		return i + 2
	}
	return 1
}

func FFT(inp []byte, n int) []byte {
	l := len(inp)
	out := make([]byte, l)
	copy(out, inp)

	for k := 0; k < n; k++ {
		for i := 0; i < l; i++ {

			x := int(out[i])
			for j := i + 1; j < l; j++ {
				k := ((j + 1) / (i + 1)) & 3
				if k == 0 || k == 2 {
					j += i
				} else {
					x += int(out[j]) * (-k + 2)
				}
			}

			out[i] = byte(adv.Abs(x) % 10)
		}
	}
	return out
}

func FFT2(inp []byte, n int, a int, b int) []byte {
	l := len(inp)
	lo := l - a
	out := make([]byte, lo)
	for i := 0; i < lo; i++ {
		out[i] = inp[a+i]
	}

	for k := 0; k < n; k++ {
		v := 0
		for i := lo - 1; i >= 0; i-- {
			v = (v + int(out[i])) % 10
			out[i] = byte(v)
		}
	}

	return out[0 : b-a]
}

func ParseInput(input string) []byte {
	out := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		out[i] = byte(input[i] - '0')
	}
	return out
}

func ToNumber(n int, inp []byte) int {
	x := 0
	for i := 0; i < n; i++ {
		x = x*10 + int(inp[i])
	}
	return x
}

func Repeated(inp []byte) []byte {
	l := len(inp)
	for i := 1; i < 10000; i++ {
		inp = append(inp, inp[:l]...)
	}
	return inp
}

func Offset(inp []byte) int {
	return ToNumber(7, inp)
}

func main() {
	input := "59717238168580010599012527510943149347930742822899638247083005855483867484356055489419913512721095561655265107745972739464268846374728393507509840854109803718802780543298141398644955506149914796775885246602123746866223528356493012136152974218720542297275145465188153752865061822191530129420866198952553101979463026278788735726652297857883278524565751999458902550203666358043355816162788135488915722989560163456057551268306318085020948544474108340969874943659788076333934419729831896081431886621996610143785624166789772013707177940150230042563041915624525900826097730790562543352690091653041839771125119162154625459654861922989186784414455453132011498"

	{
		inp := ParseInput(input)
		inp = FFT(inp, 100)
		fmt.Println(ToNumber(8, inp))
	}

	{
		inp := ParseInput(input)
		inp = Repeated(inp)

		offset := Offset(inp)
		fmt.Println(offset)

		out := FFT2(inp, 100, offset, offset+8)

		x := ToNumber(8, out)
		fmt.Println(x)
		if x <= 59717238 {
			panic("Too low")
		}
	}
}

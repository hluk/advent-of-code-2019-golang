package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"
)

const CmdDeal string = "deal into new stack"
const CmdCutN string = "cut"
const CmdDealN string = "deal with increment"

type Deck struct {
	Cards []int
}

func CreateDeck(n int) Deck {
	d := Deck{make([]int, n)}
	for i := 0; i < n; i++ {
		d.Cards[i] = i
	}
	return d
}

func (d Deck) Deal() Deck {
	l := len(d.Cards)
	d2 := Deck{make([]int, l)}
	for i := 0; i < l; i++ {
		d2.Cards[i] = d.Cards[l-i-1]
	}
	return d2
}

func (d Deck) DealN(n int) Deck {
	l := len(d.Cards)
	d2 := Deck{make([]int, l)}
	for i := 0; i < l; i++ {
		j := (i * n) % l
		d2.Cards[j] = d.Cards[i]
	}
	return d2
}

func (d Deck) CutN(n int) Deck {
	l := len(d.Cards)

	if n < 0 {
		n = l + n
	}

	d2 := Deck{make([]int, l)}
	for i := 0; i < l-n; i++ {
		d2.Cards[i] = d.Cards[i+n]
	}
	for i := 0; i < n; i++ {
		d2.Cards[l-n+i] = d.Cards[i]
	}
	return d2
}

func MatchCmd(cmd, cmdToMatch string) (bool, int) {
	if len(cmd) < len(cmdToMatch) || cmd[:len(cmdToMatch)] != cmdToMatch {
		return false, 0
	}

	if len(cmd) == len(cmdToMatch) {
		return true, 0
	}

	i := strings.LastIndex(cmd, " ")
	x := cmd[i+1:]
	n, err := strconv.Atoi(x)
	if err != nil {
		panic(err)
	}
	return true, n
}

func (d Deck) Shuffle(cmds []string) Deck {
	for _, cmd := range cmds {
		if cmd == CmdDeal {
			d = d.Deal()
		} else if ok, n := MatchCmd(cmd, CmdCutN); ok {
			d = d.CutN(n)
		} else if ok, n := MatchCmd(cmd, CmdDealN); ok {
			d = d.DealN(n)
		}
	}
	return d
}

func (d Deck) IndexOf(card int) int {
	l := len(d.Cards)
	for i := 0; i < l; i++ {
		if d.Cards[i] == card {
			return i
		}
	}
	panic("No such card")
}

func ReadCommands(path string) []string {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	dat = dat[:len(dat)]

	txt := string(dat)
	txt = strings.TrimRight(txt, "\n")
	lines := strings.Split(txt, "\n")

	return lines
}

func IndexShuffle(index int64, l int64, times int64, cmds []string) int64 {
	v := big.NewInt(index)
	m := big.NewInt(l)
	t := big.NewInt(times)

	i := big.NewInt(0)
	d := big.NewInt(1)

	for _, cmd := range cmds {
		if cmd == CmdDeal {
			d.Neg(d)
			i.Add(i, d)
		} else if ok, n := MatchCmd(cmd, CmdCutN); ok {
			x := big.NewInt(int64(n))
			i.Add(i, x.Mul(x, d))
		} else if ok, n := MatchCmd(cmd, CmdDealN); ok {
			x := big.NewInt(int64(n))
			x.ModInverse(x, m)
			d.Mul(d, x)
		}
	}

	// (1-d)**(m-2) % m
	a := big.NewInt(1)
	a.Sub(a, d).ModInverse(a, m)

	// d = d**t % m
	d.Exp(d, t, m)
	// it = (1-d) * a * i
	it := big.NewInt(1)
	it.Sub(it, d).Mul(it, a).Mul(it, i)
	v.Mul(v, d).Add(v, it).Mod(v, m)

	return v.Int64()
}

func main() {
	cmds := ReadCommands("advent22.txt")

	d := CreateDeck(10007)
	d = d.Shuffle(cmds)
	fmt.Println(d.IndexOf(2019))

	l := int64(119315717514047)
	times := int64(101741582076661)
	index := IndexShuffle(2020, l, times, cmds)
	fmt.Println(index)
	if index != 96196710942473 {
		panic("Wrong!")
	}
}

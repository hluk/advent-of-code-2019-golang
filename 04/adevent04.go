package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 256310
	b := 732736

	count1 := 0
	count2 := 0
	for i := a; i <= b; i++ {
		s := strconv.Itoa(i)
		if s[0] <= s[1] && s[1] <= s[2] && s[2] <= s[3] && s[3] <= s[4] && s[4] <= s[5] {
			if s[0] == s[1] || s[1] == s[2] || s[2] == s[3] || s[3] == s[4] || s[4] == s[5] {
				count1++

				m := make(map[rune]byte, 6)
				for _, v := range s {
					m[v]++
				}
				for k := range m {
					if m[k] == 2 {
						count2++
						break
					}
				}
			}
		}
	}
	fmt.Println(count1, count2)
}

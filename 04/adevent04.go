package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	start, stop := 256310, 732736
	count1, count2 := 0, 0
	for i := start; i <= stop; i++ {
		s := strconv.Itoa(i)
		sorted := sort.SliceIsSorted(s, func(i, j int) bool { return s[i] < s[j] })
		if sorted {
			inc1 := 0
			for i := 0; i < len(s)-1; {
				j := strings.LastIndexByte(s, s[i])
				if i != j {
					inc1 = 1
					if i+1 == j {
						count2++
						break
					}
				}
				i = j + 1
			}
			count1 += inc1
			continue
		}
	}
	fmt.Println(count1, count2)
}

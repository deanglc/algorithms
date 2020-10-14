package main

import (
	"fmt"
	"math"
)

func commonChars(a []string) (ans []string) {
	minFreq := [26]uint8{}
	for i := range minFreq {
		minFreq[i] = math.MaxUint8
	}
	for _, word := range a {
		freq := [26]uint8{}
		for _, b := range word {
			freq[b-'a']++
		}
		for i, f := range freq[:] {
			if f < minFreq[i] {
				minFreq[i] = f
			}
		}
	}
	for i := byte(0); i < 26; i++ {
		for j := uint8(0); j < minFreq[i]; j++ {
			ans = append(ans, string('a'+i))
		}
	}
	return
}

func main() {
	s := []string{"aaa", "bbb", "cc"}
	// s := []string{"bella", "label", "roller"}
	fmt.Println(commonChars(s))
}

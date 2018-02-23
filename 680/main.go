package main

import (
	"fmt"
)

func main() {
	test := "abxbbaszxz"
	fmt.Println(validPalindrome(test))
}

func validPalindrome(s string) bool {
	for lIndex, rIndex := 0, len(s)-1; lIndex < rIndex; lIndex, rIndex = lIndex+1, rIndex-1 {
		if !((s)[lIndex] == (s)[rIndex]) {
			return paliHelper(&s, lIndex+1, rIndex) || paliHelper(&s, lIndex, rIndex-1)
		}
	}
	return true
}

func paliHelper(s *string, lIndex, rIndex int) bool {
	for ; lIndex < rIndex; lIndex, rIndex = lIndex+1, rIndex-1 {
		if (*s)[lIndex] != (*s)[rIndex] {
			return false
		}
	}
	return true
}

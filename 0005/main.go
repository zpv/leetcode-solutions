package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)

	max := 0
	cur := 0
	curLength := 0

	bestLeft := 0
	bestRight := 0

	leftRunner := 0
	rightRunner := len(r) - 1
	// abba
	for cur < len(r)-1 {
		if leftRunner == rightRunner {
			cur++
			curLength = 0
			leftRunner = cur
			rightRunner = len(r) - 1
		} else if r[leftRunner] == r[rightRunner] {
			leftRunner++
			rightRunner--
			curLength++
			if leftRunner >= rightRunner {
				if curLength > max {
					bestLeft = leftRunner - curLength
					bestRight = rightRunner + curLength
					max = curLength
					curLength = 0
					cur++
					leftRunner = cur
					rightRunner = len(r) - 1
				} else {
					curLength = 0
					cur++
					leftRunner = cur
					rightRunner = len(r) - 1
				}
			}
		} else {
			rightRunner--
			leftRunner = cur

			curLength = 0
			if rightRunner < leftRunner {
				cur++
				curLength = 0
				leftRunner = cur
				rightRunner = len(r) - 1
			}
		}
	}

	return string(r[bestLeft : bestRight+1])
}

func main() {
	fmt.Println(longestPalindrome("aaabaaaa"))
}

package main

import (
	"fmt"
)

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	depth := make(map[byte]int)
	maxLength := 0

	left := 0

	for i := 0; i < len(s); i++ {
		depth[s[i]]++

		for len(depth) == k+1 {
			maxLength = max(maxLength, i-left)
			depth[s[left]]--

			if s[left] == 0 {
				fmt.Println(len(depth))
				delete(depth, s[left])
				fmt.Println(len(depth))
			}

			left++
		}
	}

	return max(maxLength, len(s)-1-left)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lengthOfLongestSubstringKDistinct("bacc", 2))
}

// eceba
// aa
// yuina

// Possible approach:
// Store a pointer to "last repeat of the first distinct character" of substring

// aabccdeefg
// k = 2

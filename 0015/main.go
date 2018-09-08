package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}

	sort.Ints(nums)

	if len(nums) < 3 {
		return result
	}

	if nums[0] == 0 && nums[len(nums)-1] == 0 {
		return [][]int{{0, 0, 0}}
	}

	// a + b + c = 0

	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]

		if a > 0 {
			return result
		}

		if i > 0 && nums[i-1] == a {
			continue
		}

		for j, k := i+1, len(nums)-1; j < k; {
			b := nums[j]
			c := nums[k]

			if c < 0 {
				break
			}

			sum := a + b + c

			if sum > 0 {
				for j < k && nums[k] == c {
					k--
				}
			} else if sum == 0 {
				result = append(result, []int{a, b, c})
				for j < k && nums[k] == c && nums[k] >= 0 {
					k--
				}
				for j < k && nums[j] == b && nums[k] >= 0 {
					j++
				}
			} else if nums[k] >= 0 {
				for j < k && nums[j] == b {
					j++
				}
			}
		}

	}

	return result
}

func main() {
	threeSum([]int{-1, 0, 1, 2, -1, -4})
}

/**
  Trivial Case:  [0, 0, 0]
  Solution: [0, 0 ,0]

  Trivial Case: []
  Solution: []

  Trivial Case: [1, -1]
  Solution: []

  Trivial Case: []
*/

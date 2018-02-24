package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if val, ok := m[target-v]; ok {
			if k == val {
				continue
			} else {
				if k > val {
					return []int{val, k}
				}
				return []int{k, val}
			}
		}
		m[v] = k
	}

	return []int{0, 0}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

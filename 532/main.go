package main

func main() {

}

func findPairs(nums []int, k int) int {
	if len(nums) == 0 || k < 0 {
		return 0
	}

	m := make(map[int]int)
	diff := 1

	if k == 0 {
		diff = 2
	}

	for _, v := range nums {
		m[v]++
	}

	sum := 0

	for _, v := range nums {
		if m[v-k] > diff-1 {
			sum++
			m[v-k] = 0
		}
	}

	return sum
}

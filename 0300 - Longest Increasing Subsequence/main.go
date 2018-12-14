package main

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))

	gmax := 0

	for i := len(nums) - 1; i >= 0; i-- {
		m := 0

		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				m = max(m, dp[j])
			}
		}
		dp[i] = m + 1
		gmax = max(gmax, dp[i])
	}

	return gmax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}

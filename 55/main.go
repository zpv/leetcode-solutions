package main

func canJump(nums []int) bool {
	cur := 0
	i := nums[0]

	if cur == len(nums)-1 {
		return true
	}

	for i > 0 {
		cur++
		i--
		if nums[cur] > i {
			i = nums[cur]
		}

		if cur == len(nums)-1 {
			return true
		}
	}

	return false
}

func main() {

}

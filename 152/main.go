package main

func maxProduct(nums []int) int {
	max := nums[0]
	length := len(nums)

	num := 0

	for i := 0; i < length; i++ {
		if num == 0 {
			num = nums[i]
		} else {
			num *= nums[i]
		}
		max = maxInt(num, max)
	}

	num = 0

	for i := length - 1; i >= 0; i-- {
		if num == 0 {
			num = nums[i]
		} else {
			num *= nums[i]
		}
		max = maxInt(num, max)
	}
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}

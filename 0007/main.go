package main

import "math"

func reverse(x int) int {
	digits, num, bit := 0, abs(x), signbit(x)

	temp := num

	for temp > 0 {
		digits++
		temp /= 10
	}

	sum := 0

	for i := 0; i < digits; i++ {
		digit := num % 10
		num /= 10

		sum += pow(10, (digits-1-i)) * digit
	}

	if bit {
		sum = -sum
	}

	if sum > math.MaxInt32 {
		sum = 0
	} else if sum < math.MinInt32 {
		sum = 0
	}

	return sum
}

func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}

func signbit(n int) bool {
	return uint(n)&(1<<63) != 0
}

func pow(n, p int) int {
	if p == 0 {
		return 1
	}

	if p == 1 {
		return n
	}

	if p%2 == 1 {
		return n * pow(n*n, (p-1)/2)
	}

	return pow(n*n, p/2)
}

func main() {
	reverse(123)
}

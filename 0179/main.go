package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
)

/*

Given a list of non negative integers, arrange them such that they form the largest number.

For example, given [3, 30, 34, 5, 9], the largest formed number is 9534330.

9, 5, 34, 3, 30

33, 30, 34, 55, 99

[421,452,641,365,423, 45]

641
45
452
423
421
365

[824,938,1399,5607,6973,5703,9609,4398,8247]

"9609938 824 8247 69735703560743981399"
"9609938 8247 824  69735703560743981399"

9609
938(8)
8247
824(4)
6973
5703
5607
4398
1399



Note: The result may be very large, so you need to return a string instead of an integer.

Credits:
Special thanks to @ts for adding this problem and creating all test cases.

*/

func largestNumber(nums []int) string {
	quickSort(nums)

	result := bytes.NewBufferString("")

	for _, v := range nums {
		result.WriteString(strconv.Itoa(v))
	}

	if a, _ := strconv.Atoi(result.String()); a == 0 {
		return "0"
	}

	return result.String()
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1
	pivotIndex := rand.Int() % len(arr)
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range arr {
		lr, _ := strconv.Atoi(strconv.Itoa(arr[i]) + strconv.Itoa(arr[right]))
		rl, _ := strconv.Atoi(strconv.Itoa(arr[right]) + strconv.Itoa(arr[i]))
		if lr > rl {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	arr[left], arr[right] = arr[right], arr[left]

	// Go down the rabbit hole
	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}

func main() {
	test := []int{55, 00, 23}
	fmt.Println(largestNumber(test))
}

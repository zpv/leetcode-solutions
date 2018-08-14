package main

func lengthOfLongestSubstring(s string) int {
	indices := make(map[rune]int)
	max, left, cur := 0, 0, 0

	runes := []rune(s)

	for cur < len(s) {
		character := runes[cur]

		if val, ok := indices[character]; ok {
			if val >= left {
				left = val + 1
			}
		}
		if (cur - left + 1) > max {
			max = cur - left + 1
		}
		indices[character] = cur
		cur++
	}

	return max
}

func main() {

}

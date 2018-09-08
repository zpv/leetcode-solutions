package main

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		if len(s)-i < (end-start)/2 {
			break
		}
		l1, r1 := expandAroundCentre(s, i, i)
		l2, r2 := expandAroundCentre(s, i, i+1)

		len1 := r1 - l1
		len2 := r2 - l2

		if max(len1, len2) > end-start {
			if len1 > len2 {
				start, end = l1, r1
			} else {
				start, end = l2, r2
			}
		}
	}

	return s[start : end+1]
}

func expandAroundCentre(s string, left int, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}

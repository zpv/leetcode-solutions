package main

/**
  DP decision points:
  - If star is present in the pattern, either:
      - Ignore part of this pattern
      - Delete matching character in text
*/

type input struct {
	si,
	pi int
}

func isMatch(s string, p string) bool {
	memo := make(map[input]bool)
	return isMatchHelper(0, 0, &s, &p, memo)
}

func isMatchHelper(si, pi int, s, p *string, memo map[input]bool) bool {
	if val, ok := memo[input{si, pi}]; ok {
		return val
	}

	if pi >= len(*p) {
		memo[input{si, pi}] = si >= len(*s)
		return memo[input{si, pi}]
	}

	if pi+1 < len(*p) && (*p)[pi+1] == '*' {
		memo[input{si, pi}] = isMatchHelper(si, pi+2, s, p, memo) ||
			(matches(si, pi, s, p) && isMatchHelper(si+1, pi, s, p, memo))
		return memo[input{si, pi}]
	} else {
		memo[input{si, pi}] = matches(si, pi, s, p) && isMatchHelper(si+1, pi+1, s, p, memo)
		return memo[input{si, pi}]
	}
}

func matches(si, pi int, s, p *string) bool {
	if si >= len(*s) || pi >= len(*p) {
		return false
	}

	if (*s)[si] == (*p)[pi] {
		return true
	}

	var dot uint8
	dot = '.'

	if (*p)[pi] == dot {
		return true
	}

	return false
}

func main() {

}

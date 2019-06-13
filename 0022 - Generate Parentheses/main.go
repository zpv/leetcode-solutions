package main

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	helper(0, 0, n, "", &res)

	return res
}

func helper(open, total, max int, cur string, result *[]string) {
	if len(cur) == max*2 {
		*result = append(*result, cur)
		return
	}

	if open > 0 {
		helper(open-1, total, max, cur+")", result)
	}
	if total < max {
		helper(open+1, total+1, max, cur+"(", result)
	}
}

func main() {

}

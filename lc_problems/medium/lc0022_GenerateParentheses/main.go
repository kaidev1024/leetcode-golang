func generateParenthesis(n int) []string {
	arr := make([]string, 0)

	var helper func(cur string, l, r int)
	helper = func(cur string, l, r int) {
		if l == 0 && r == 0 {
			arr = append(arr, cur)
			return
		}
		if l > 0 {
			helper(cur+"(", l-1, r)
		}
		if r > l {
			helper(cur+")", l, r-1)
		}
	}
	helper("", n, n)
	return arr
}
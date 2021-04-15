func optimalDivision(nums []int) string {
	numStrs := make([]string, len(nums))
	for i, n := range nums {
		numStrs[i] = strconv.Itoa(n)
	}

	if len(nums) <= 2 {
		return strings.Join(numStrs, `/`)
	}
	return fmt.Sprintf("%s/(%s)", numStrs[0], strings.Join(numStrs[1:], `/`))
}
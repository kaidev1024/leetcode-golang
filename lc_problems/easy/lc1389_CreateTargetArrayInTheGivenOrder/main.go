func createTargetArray(nums []int, index []int) []int {
	test := map[int]bool{}
	output := []int{}
	for k, v := range index {
		if _, ok := test[v]; ok {
			test[k] = true
			output = append(output, nums[index[k]])
		} else {
			output = append(output, 0)
			copy(output[v+1:], output[v:])
			output[index[k]] = nums[k]
		}
	}
	return output
}
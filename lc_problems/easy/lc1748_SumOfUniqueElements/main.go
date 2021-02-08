func sumOfUnique(nums []int) int {
	arr := make([]int, 101)
	for _, v := range nums {
		arr[v]++
	}
	result := 0
	for i, v := range arr {
		if v == 1 {
			result += i
		}
	}
	return result
}
func largestSubmatrix(matrix [][]int) int {
	var res int
	for i := 1; i < len(matrix); i++ {
		for j := range matrix[0] {
			if matrix[i][j] == 1 {
				matrix[i][j] += matrix[i-1][j]
			}
		}
	}
	l := len(matrix[0])
	for _, nums := range matrix {
		sort.Ints(nums)
		for i := 1; i <= l; i++ {
			if nums[l-i]*i > res {
				res = nums[l-i] * i
			}
		}
	}
	return res
}
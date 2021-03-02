func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	res := make([]bool, 0)

	for i := 0; i < len(l); i++ {
		start, end := l[i], r[i]+1
		n := nums[start:end]

		copyN := make([]int, end-start)
		copy(copyN, n)

		a := isArithmetic(copyN)
		res = append(res, a)
	}

	return res
}

func isArithmetic(nums []int) bool {
	sort.Ints(nums)

	i, j := 0, 1

	start := false
	diff := 0
	for j < len(nums) {

		if start && diff != nums[j]-nums[i] {
			return false
		}

		diff = nums[j] - nums[i]

		if !start {
			start = true
		}

		i++
		j++
	}
	return true
}
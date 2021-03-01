func maxResult(nums []int, k int) int {
	n := len(nums)
	queue := make([]int, 0, n)
	x := nums[n-1]
	queue = append(queue, x)

	for i := n - 2; i >= 0; i-- {
		x = queue[0] + nums[i]
		queue = append(queue, x)

		if len(queue) > k {
			queue = queue[1:]
		}

		for x > queue[0] {
			queue = queue[1:]
		}
	}
	return x
}
func rangeSum(nums []int, n int, left int, right int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	arr := make([]int, sum+1)
	for i := 0; i < n; i++ {
		amount := 0
		for j := i; j < n; j++ {
			amount += nums[j]
			arr[amount]++
		}
	}
	ret, idx := 0, 0
	for i := 1; i < sum+1; {
		if arr[i] > 0 {
			arr[i]--
			idx++
			if idx >= left && idx <= right {
				ret += i
			} else if idx > right {
				break
			}
		} else {
			i++
		}
	}
	return ret % int(math.Pow10(9)+7)
}
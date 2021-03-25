func isPossibleDivide(nums []int, k int) bool {
	n := len(nums)
	arr := make([]int, 0, n)
	counts := make(map[int]int)

	for _, num := range nums {
		_, ok := counts[num]
		if !ok {
			arr = append(arr, num)
		}
		counts[num]++
	}
	sort.Ints(arr)
	na := len(arr)
	for i := na - 1; i >= 0; i-- {
		curNum := arr[i]
		count, _ := counts[curNum]
		if count == 0 {
			continue
		}
		for j := 0; j < k; j++ {
			if counts[curNum-j] < count {
				return false
			}
			counts[curNum-j] -= count
		}
	}
	return true
}
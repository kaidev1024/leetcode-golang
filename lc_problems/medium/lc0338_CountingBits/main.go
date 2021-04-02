func countBits(num int) []int {
	arr := make([]int, num+1)
	for i := 1; i <= num; i++ {
		if i&1 == 0 {
			arr[i] = arr[i>>1]
		} else {
			arr[i] = arr[i>>1] + 1
		}
	}
	return arr
}
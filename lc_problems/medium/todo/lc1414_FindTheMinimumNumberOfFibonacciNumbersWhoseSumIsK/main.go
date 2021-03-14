func findMinFibonacciNumbers(k int) int {
	ans := 0
	f1, f2 := 1, 1
	arr := []int{f1, f2}
	for {
		if f1+f2 > k {
			break
		}
		f1, f2 = f2, f1+f2
		arr = append(arr, f2)
	}
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] <= k {
			ans++
			k = k - arr[i]
		}
		if k == 0 {
			return ans
		}
	}
	return ans
}

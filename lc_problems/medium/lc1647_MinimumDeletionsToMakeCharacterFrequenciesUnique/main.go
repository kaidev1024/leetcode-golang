func minDeletions(s string) int {
	freqs := make([]int, 26)
	for _, c := range s {
		freqs[c-'a']++
	}

	arr := make([]int, 0, 26)
	for _, freq := range freqs {
		if freq > 0 {
			arr = append(arr, freq)
		}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	fmt.Println(arr)
	n := len(arr)
	result := 0
	for i := 1; i < n; i++ {
		if arr[i] < arr[i-1] {
			continue
		}
		if arr[i-1] == 0 {
			result += arr[i]
			arr[i] = 0
			continue
		}
		diff := arr[i] - arr[i-1]
		arr[i] -= diff + 1
		result += diff + 1
	}
	return result
}
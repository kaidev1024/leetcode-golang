func frequencySort(s string) string {
	arr := []byte(s)
	freq := make(map[byte]int)
	for _, char := range arr {
		freq[char]++
	}
	sort.Slice(arr, func(i, j int) bool {
		if freq[arr[i]] == freq[arr[j]] {
			return arr[i] < arr[j]
		}
		return freq[arr[i]] > freq[arr[j]]
	})

	return string(arr)
}
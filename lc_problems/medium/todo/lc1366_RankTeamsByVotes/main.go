func rankTeams(votes []string) string {
	lvote := len(votes[0])
	counts := make(map[byte][]int)
	for _, vote := range votes {
		for i := 0; i < lvote; i++ {
			if _, ok := counts[vote[i]]; !ok {
				counts[vote[i]] = make([]int, lvote)
			}
			counts[vote[i]][i]++
		}
	}

	arr := []byte(votes[0])
	sort.Slice(arr, func(i, j int) bool {
		for k := 0; k < lvote; k++ {
			if counts[arr[i]][k] != counts[arr[j]][k] {
				return counts[arr[i]][k] > counts[arr[j]][k]
			}
		}
		return arr[i] < arr[j]
	})
	return string(arr)
}
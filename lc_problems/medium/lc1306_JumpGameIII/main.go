func canReach(arr []int, start int) bool {
	queue, l := []int{start}, len(arr)
	seen := make([]bool, l)
	seen[start] = true
	for len(queue) > 0 {
		index := queue[0]
		queue = queue[1:]
		if arr[index] == 0 {
			return true
		}
		for _, nextIndex := range []int{index + arr[index], index - arr[index]} {
			if 0 <= nextIndex && nextIndex < l && seen[nextIndex] == false {
				seen[nextIndex] = true
				queue = append(queue, nextIndex)
			}
		}
	}
	return false
}
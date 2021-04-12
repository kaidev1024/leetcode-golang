func minMutation(start string, end string, bank []string) int {
	set := make(map[string]struct{})
	for _, s := range bank {
		set[s] = struct{}{}
	}

	queue := []string{start}
	lq := 1
	result := 0
	for lq > 0 {
		result++
		for _, str := range queue {
			for key, _ := range set {

				if isNear(key, str) {
					if key == end {
						return result
					}
					delete(set, key)
					queue = append(queue, key)
				}
			}
		}
		queue = queue[lq:]
		lq = len(queue)
	}
	return -1
}

func isNear(s1, s2 string) bool {
	dif := false
	for i, char := range s1 {
		if char != rune(s2[i]) {
			if dif {
				return false
			}
			dif = true
		}
	}
	return dif
}
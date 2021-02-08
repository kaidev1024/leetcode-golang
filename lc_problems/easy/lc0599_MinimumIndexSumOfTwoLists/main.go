func findRestaurant(list1 []string, list2 []string) []string {
	m := make(map[string]int)
	l1 := len(list1)
	for i := 0; i < l1; i++ {
		m[list1[i]] = i
	}

	l2 := len(list2)
	minVal := l1 + l2
	result := make([]string, 0, l1)

	for i := 0; i < l2; i++ {
		if ind, ok := m[list2[i]]; ok {
			sum := ind + i
			if sum == minVal {
				result = append(result, list2[i])
			} else if sum < minVal {
				result = result[0:0]
				result = append(result, list2[i])
				minVal = sum
			}
		}
	}

	return result
}
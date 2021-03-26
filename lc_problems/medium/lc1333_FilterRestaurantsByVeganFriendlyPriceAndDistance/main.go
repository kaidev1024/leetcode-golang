func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	filtered := make([][]int, 0)
	for _, rest := range restaurants {
		if rest[3] > maxPrice || rest[4] > maxDistance || (veganFriendly == 1 && rest[2] == 0) {
			continue
		}
		filtered = append(filtered, rest)
	}
	sort.Slice(filtered, func(i, j int) bool {
		if filtered[i][1] == filtered[j][1] {
			return filtered[i][0] > filtered[j][0]
		}
		return filtered[i][1] > filtered[j][1]
	})
	result := make([]int, len(filtered))
	for i := range filtered {
		result[i] = filtered[i][0]
	}

	return result
}
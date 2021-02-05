func distributeCandies(candies int, num_people int) []int {
	result := make([]int, num_people)
	k := 1

	for i := 0; ; i++ {
		for j := 0; j < num_people; j++ {
			if candies <= k {
				result[j] += candies
				return result
			} else {
				result[j] += k
				candies -= k
				k++
			}
		}
	}
	return result
}
func findNumbers(nums []int) int {

	count, even := 0, 0
	for _, num := range nums {
		for num > 0 {
			count++
			num = num / 10
		}
		if count%2 == 0 {
			even++
		}
		count = 0
	}

	return even
}
func maxProduct(nums []int) int {
	n1, n2 := 1, 1
	for _, num := range nums {
		if num > n1 {
			n2 = n1
			n1 = num
			continue
		}
		if num > n2 {
			n2 = num
		}
	}
	return (n1 - 1) * (n2 - 1)
}
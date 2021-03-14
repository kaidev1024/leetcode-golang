func kthFactor(n int, k int) int {
	firstHalf := make([]int, 0)
	secondHalf := make([]int, 0)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			firstHalf = append(firstHalf, i)
			temp := n / i
			if temp != i {
				secondHalf = append([]int{temp}, secondHalf...)
			}
		}
	}
	firstHalf = append(firstHalf, secondHalf...)
	if k > len(firstHalf) {
		return -1
	}
	return firstHalf[k-1]
}
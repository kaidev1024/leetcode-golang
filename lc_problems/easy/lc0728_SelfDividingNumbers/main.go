func selfDividingNumbers(left int, right int) []int {

	var array []int
	for i := left; i <= right; i++ {
		k := i
		isVaild := true
		for k != 0 {
			if num := k % 10; num == 0 || i%num != 0 {
				isVaild = false
				break
			}
			k = k / 10
		}
		if isVaild {
			array = append(array, i)
		}
	}
	return array

}
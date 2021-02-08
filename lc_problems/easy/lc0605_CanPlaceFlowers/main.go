func canPlaceFlowers(flowerbed []int, n int) bool {
	flowerbed = append([]int{0}, flowerbed...)
	flowerbed = append(flowerbed, 0)
	l := len(flowerbed)
	result := 0

	for i := 1; i < l-1; i++ {
		if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			result++
			flowerbed[i] = 1
		}
	}

	return result >= n
}
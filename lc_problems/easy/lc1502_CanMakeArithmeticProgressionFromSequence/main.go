func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	dif := arr[1] - arr[0]
	for i, l := 2, len(arr); i < l; i++ {
		if dif != arr[i]-arr[i-1] {
			return false
		}
	}
	return true
}
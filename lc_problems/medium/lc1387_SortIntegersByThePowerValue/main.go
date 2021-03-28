var m map[int]int

func getPower(x int) int {
	if v, ok := m[x]; ok {
		return v
	}

	var y int
	if x%2 == 0 {
		y = x / 2
	} else {
		y = 3*x + 1
	}

	p := getPower(y)
	m[x] = p + 1
	return p + 1
}

func getKth(lo int, hi int, k int) int {
	m = make(map[int]int)
	cur := 1
	count := 0
	for cur <= hi {
		m[cur] = count
		cur *= 2
		count++
	}

	arr := [][]int{}

	for i := lo; i <= hi; i++ {
		m[i] = getPower(i)
		arr = append(arr, []int{m[i], i})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0] || arr[i][0] == arr[j][0] && arr[i][1] < arr[j][1]
	})

	return arr[k-1][1]
}
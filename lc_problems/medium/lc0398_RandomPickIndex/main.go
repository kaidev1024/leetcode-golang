type Solution struct {
	m map[int][]int
}

func Constructor(nums []int) Solution {
	m := make(map[int][]int)
	for i, v := range nums {
		m[v] = append(m[v], i)
	}
	return Solution{m}
}

func (this *Solution) Pick(target int) int {
	arr := this.m[target]
	ind := rand.Intn(len(arr))
	return arr[ind]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
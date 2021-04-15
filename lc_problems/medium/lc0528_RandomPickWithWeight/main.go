]type Solution struct {
    weights []int
}


func Constructor(w []int) Solution {
    n := len(w)
    weights := make([]int, n + 1)
    for i, v := range w {
        weights[i + 1] = v + weights[i]
    }
    return Solution{weights}
}


func (this *Solution) PickIndex() int {
    weights := this.weights
    n := len(weights)
    total := weights[n - 1]
    randNum := rand.Intn(total) + 1
    index := sort.SearchInts(weights, randNum)
    return index - 1
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
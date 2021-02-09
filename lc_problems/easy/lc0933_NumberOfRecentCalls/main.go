type RecentCounter struct {
	pings []int
}

func Constructor() RecentCounter {
	return RecentCounter{make([]int, 0)}
}

func (this *RecentCounter) Ping(t int) int {
	this.pings = append(this.pings, t)
	p := t - 3000

	n := len(this.pings)
	l := 0
	r := n - 1
	for l < r {
		mid := (r + l) >> 1
		if this.pings[mid] >= p {
			r = mid
		} else {
			l = mid + 1
		}
	}

	this.pings = this.pings[l:]
	return len(this.pings)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
import "math/rand"

type RandomizedSet struct {
	arr []int
	m   map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		arr: []int{},
		m:   make(map[int]int),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	n := len(this.arr)
	this.m[val] = n
	this.arr = append(this.arr, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {

	if index, ok := this.m[val]; ok {
		n := len(this.arr)
		if n-1 != index {
			this.arr[index], this.arr[n-1] = this.arr[n-1], this.arr[index]
			this.m[this.arr[index]] = index
		}
		this.arr = this.arr[:n-1]
		delete(this.m, val)
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.arr))
	return this.arr[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
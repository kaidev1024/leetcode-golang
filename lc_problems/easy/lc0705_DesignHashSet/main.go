type MyHashSet struct {
	arr []uint64
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{
		arr: make([]uint64, 1000000/64+1),
	}
}

func (this *MyHashSet) Add(key int) {
	pos := key / 64
	res := key % 64
	this.arr[pos] |= 1 << res
}

func (this *MyHashSet) Remove(key int) {
	pos := key / 64
	res := key % 64
	setBit := uint64(1 << res)
	if this.arr[pos]&setBit > 0 {
		this.arr[pos] -= setBit
	}

}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	pos := key / 64
	res := key % 64
	return this.arr[pos]&(1<<res) > 0
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
type MyHashMap struct {
	arr [][][2]int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{arr: make([][][2]int, 1000)}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	pos := key / 1000
	arr := this.arr[pos]
	flag := false
	for i, v := range arr {
		if v[0] == key {
			this.arr[pos][i][1] = value
			flag = true
			break
		}
	}
	if !flag {
		this.arr[pos] = append(this.arr[pos], [2]int{key, value})
	}
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	pos := key / 1000
	arr := this.arr[pos]
	for _, v := range arr {
		if v[0] == key {
			return v[1]
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	pos := key / 1000
	arr := this.arr[pos]
	index := -1
	for i, v := range arr {
		if v[0] == key {
			index = i
			break
		}
	}
	if index != -1 {
		this.arr[pos] = append(this.arr[pos][:index], this.arr[pos][index+1:]...)
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
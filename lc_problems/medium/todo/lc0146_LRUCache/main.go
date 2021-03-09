import "container/list"

type entry struct {
	key, value int
}

type LRUCache struct {
	capacity int
	ll       *list.List
	cache    map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		ll:       list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	ele, ok := c.cache[key]
	if !ok {
		return -1
	}
	c.ll.MoveToFront(ele)
	return ele.Value.(*entry).value
}

func (c *LRUCache) Put(key int, value int) {
	if ele, ok := c.cache[key]; ok {
		ele.Value.(*entry).value = value
		c.ll.MoveToFront(ele)
		return
	}
	newEle := c.ll.PushFront(&entry{key, value})
	c.cache[key] = newEle
	if c.ll.Len() > c.capacity {
		eleToRemove := c.ll.Back()
		c.ll.Remove(eleToRemove)
		keyToRemove := eleToRemove.Value.(*entry).key
		delete(c.cache, keyToRemove)
	}
}

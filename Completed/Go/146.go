import (
    "container/list"
)

type LRUCache struct {
	cap int
	ll  *list.List
	m   map[int]*list.Element
}

type node struct {
	key, value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity,
		list.New(),
		make(map[int]*list.Element),
	}
}

func (cache *LRUCache) Get(key int) int {
	if ele, found := cache.m[key]; found {
		defer cache.ll.MoveToFront(ele)
		return ele.Value.(node).value
	} else {
		return -1
	}
}

func (cache *LRUCache) Put(key, value int) {
	if ele, found := cache.m[key]; found {
		ele.Value = node{key, value}
		cache.ll.MoveToFront(ele)
		return
	} else if len(cache.m) == cache.cap {
		ele := cache.ll.Back()
		n := ele.Value.(node)

		cache.ll.Remove(ele)   // remove last
		delete(cache.m, n.key) // delete key from map
	}
	cache.m[key] = cache.ll.PushFront(node{key, value})
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
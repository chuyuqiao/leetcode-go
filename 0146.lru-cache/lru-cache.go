package problem0146

import "container/list"

type LRUCache struct {
	cap int
	l   *list.List
	m   map[int]*list.Element
}

type Pair struct {
	key   int
	value int
}

func New(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
		l:   new(list.List),
		m:   make(map[int]*list.Element, capacity),
	}
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.m[key]; ok {
		val := node.Value.(*list.Element).Value.(Pair).value
		c.l.MoveToFront(node)
		return val
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.m[key]; ok {
		c.l.MoveToFront(node)
		node.Value.(*list.Element).Value = Pair{key: key, value: value}
	} else {
		if c.l.Len() == c.cap {
			idx := c.l.Back().Value.(*list.Element).Value.(Pair).key
			delete(c.m, idx)
			c.l.Remove(c.l.Back())
		}
		node := &list.Element{
			Value: Pair{
				key:   key,
				value: value,
			},
		}
		ptr := c.l.PushFront(node)
		c.m[key] = ptr
	}
}

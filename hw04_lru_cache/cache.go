package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if element, exists := c.items[key]; exists {
		c.queue.MoveToFront(element)
		element.Value.(*cacheItem).value = value
		return true
	}

	item := &cacheItem{
		key:   key,
		value: value,
	}

	element := c.queue.PushFront(item)
	c.items[item.key] = element
	if c.queue.Len() > c.capacity {
		if last := c.queue.Back(); last != nil {
			c.queue.Remove(last)
			delete(c.items, last.Value.(*cacheItem).key)
		}
	}

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if element, exists := c.items[key]; exists {
		c.queue.MoveToFront(element)
		return element.Value.(*cacheItem).value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem, c.capacity)
}

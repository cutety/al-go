package lru

// LRUCache 缓存
type LRUCache struct {
	size, capacity int
	cache          map[int]*DoubleLinkedList
	head, tail     *DoubleLinkedList
}

// DoubleLinkedList 双向链表
type DoubleLinkedList struct {
	key, value int
	prev, next *DoubleLinkedList
}

func initDoubleLinkedList(key, value int) *DoubleLinkedList {
	return &DoubleLinkedList{
		key:   key,
		value: value,
	}
}

// NewLRUCache 创建一个LRUCache
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		size:     0,
		capacity: capacity,
		cache:    map[int]*DoubleLinkedList{},
		head:     initDoubleLinkedList(0, 0),
		tail:     initDoubleLinkedList(0, 0),
	}
}

// Put 存
func (l *LRUCache) Put(key int, value int) {
	if _, ok := l.cache[key]; !ok {
		node := initDoubleLinkedList(key, value)
		l.cache[key] = node
		l.addToHead(node)
		l.size++
		if l.size > l.capacity {
			removedNode := l.removeTail()
			delete(l.cache, removedNode.key)
			l.size--
		}
	} else {
		node := l.cache[key]
		node.value = value
		l.moveToHead(node)
	}
}

// Get 取
func (l *LRUCache) Get(key int) int {
	if len(l.cache) == 0 {
		return -1
	}
	node := l.cache[key]
	l.moveToHead(node)
	return node.value
}

func (l *LRUCache) addToHead(node *DoubleLinkedList) {
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
	node.prev = l.head
}

func (l *LRUCache) removeNode(node *DoubleLinkedList) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache) moveToHead(node *DoubleLinkedList) {
	l.removeNode(node)
	l.addToHead(node)
}

func (l *LRUCache) removeTail() *DoubleLinkedList {
	node := l.tail.prev
	l.removeNode(node)
	return node
}

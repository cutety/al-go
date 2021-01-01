package lru

import (
	"fmt"
	"testing"
)

func TestLRU(test *testing.T) {
	lru := NewLRUCache(3)
	lru.Put(3, 2)
	lru.Put(1, 3)
	lru.Put(2, 0)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.cache)
	lru.Put(3,3)
	fmt.Println(lru.Get(3))
}

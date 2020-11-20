package eviction

import "testing"

func TestEvict(t *testing.T) {
	fifo := &fifo{}
	cache := initCache(fifo)

	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")

	lru := &lru{}
	cache.setEvictionAlgo(lru)
	cache.add("d", "4")

	lfu := &lfu{}
	cache.setEvictionAlgo(lfu)
	cache.add("e", "5")

	for key, value := range cache.storage {
		t.Logf("the key is: %s, and the value is: %s\n", key, value)
	}
}

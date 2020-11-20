package eviction

import "fmt"

type cache struct {
	storage      map[string]string
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

type evictionAlgo interface {
	evict(c *cache)
}

// implementation fifo eviction algorithm
type fifo struct{}

func (fifo) evict(c *cache) {
	fmt.Println("Evicting by fifo strtegy")
}

// implementation lru eviction algorithm
type lru struct{}

func (lru) evict(c *cache) {
	fmt.Println("Evicting by lru strtegy")
}

// implementation lfu eviction algorithm
type lfu struct{}

func (lfu) evict(c *cache) {
	fmt.Println("Evicting by lfu strtegy")
}

func initCache(e evictionAlgo) *cache {
	storage := make(map[string]string)
	return &cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *cache) setEvictionAlgo(e evictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *cache) get(key string) {
	delete(c.storage, key)
}

func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

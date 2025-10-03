package main

import (
	"fmt"
	"sync"
	"time"
)

type cacheItem struct {
	value      interface{}
	expiration time.Time   
}

type Cache struct {
	items map[string]cacheItem
	mutex sync.RWMutex
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	expirationTime := time.Now().Add(ttl)
	c.items[key] = cacheItem{
		value:      value,
		expiration: expirationTime,
	}
	fmt.Println("Элемент был добавлен:")
	fmt.Println("Значение:", value)
	fmt.Println("Время жизни:", ttl)
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	item, found := c.items[key]
	if !found {
		return nil, false
	}
	if time.Now().After(item.expiration) {
		return nil, false
	}
	return item.value, true
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	fmt.Println("Элемент:")
	fmt.Println("Значение:", c.items[key].value)
	fmt.Println("Время жизни:", c.items[key].expiration)
	fmt.Println("Был удалён")
	delete(c.items, key)
}

func (c *Cache) startCleanup() {
	for {
		time.Sleep(1 * time.Second)
		c.mutex.Lock()
		for key, item := range c.items {
			if time.Now().After(item.expiration) {
				fmt.Println("Срок жизни элемента истёк:")
				fmt.Println("Значение:", c.items[key])
				delete(c.items, key)
			}
		}
		c.mutex.Unlock()
	}
}

func main() {
	cache := Cache{
		items: make(map[string]cacheItem),
	}
	go cache.startCleanup()

	cache.Set("Test", "Work", 3*time.Second)
	val, ok := cache.Get("Test")
	if ok {
		fmt.Println("Значение:", val)
	} else {
		fmt.Println("Значение не найдено или истекло")
	}
	time.Sleep(4 * time.Second)

}

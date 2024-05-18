package main

import (
	"fmt"
	"sync"
	"time"
)

type Map struct {
	mx sync.Mutex
	m  map[string]interface{}
}

func NewMap() *Map {
	return &Map{
		m: make(map[string]interface{}),
	}
}

func (c *Map) Get(key string) (interface{}, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	val, ok := c.m[key]
	return val, ok
}

func (c *Map) Put(key string, value interface{}) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key] = value
}

func (c *Map) Size() int {
	return len(c.m)
}

func (c *Map) Delete(key string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.m, key)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	m := NewMap()
	var quit bool
	go func() {
		defer wg.Done()
		for i := 0; !quit; i++ {
			m.Put("a", i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; !quit; i++ {
			m.Put("b", i)
		}
	}()

	time.Sleep(2 * time.Second)
	quit = true
	fmt.Println(m.Get("a"))
	fmt.Println(m.Get("b"))
	wg.Wait()
}

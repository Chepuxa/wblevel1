package main

import (
	"fmt"
	"sync"
	"time"
)

//Реализовать конкурентную запись данных в map.
type Map struct {
	mx sync.RWMutex
	m  map[string]interface{}
}

func NewMap() *Map {
	return &Map{
		m: make(map[string]interface{}),
	}
}

func (c *Map) Get(key string) (interface{}, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, ok := c.m[key]
	return val, ok
}

func (c *Map) Put(key string, value interface{}) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key] = value
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

	time.Sleep(time.Second)
	quit = true
	fmt.Println(m.Get("a"))
	fmt.Println(m.Get("b"))
	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
)

/*
Вопрос: Что выведет данная программа и почему?

	func main() {
	  wg := sync.WaitGroup{}
	  for i := 0; i < 5; i++ {
	     wg.Add(1)
	     go func(wg sync.WaitGroup, i int) {
	        fmt.Println(i)
	        wg.Done()
	     }(wg, i)
	  }
	  wg.Wait()
	  fmt.Println("exit")
	}

Ответ:
deadlock, т.к. в анонимной функции, в которой выполняется горутина, будет использоваться копия WaitGroup, и ранее добавленная wg.Add(1) никогда не сократится
Это можно исправить, передав в анонимную функцию поинтер на WaitGroup
*/
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(&wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}

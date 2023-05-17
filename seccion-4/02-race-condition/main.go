package main

import (
	"fmt"
	"sync"
)

var counter int = 0
var mutex sync.Mutex

func main()  {
	var wg sync.WaitGroup
	wg.Add(2)
	go incremetCounter(&wg, &mutex)
	go incremetCounter(&wg, &mutex)
	wg.Wait()

	fmt.Println(counter)
}

func incremetCounter(wg *sync.WaitGroup, m *sync.Mutex)  {
	for i := 0; i < 100; i++ {
		m.Lock()
		counter++
		m.Unlock()
	}
	wg.Done()
}

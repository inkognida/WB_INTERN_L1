package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	Value int
}

func (c *Counter) Inc() {
	c.Value++
}

func main() {
	// структура счетчика
	counter := &Counter{Value: 0}

	// создаем wg для контроля выполнения всех горутин
	wg := sync.WaitGroup{}

	// создаем mu для синхронизации доступа к counter
	mu := sync.Mutex{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(counter *Counter) {
			defer wg.Done()
			mu.Lock()
			for j := 0; j < 10; j++{
				counter.Inc()
			}
			mu.Unlock()
		}(counter)
	}

	wg.Wait()

	fmt.Println(counter.Value)
}
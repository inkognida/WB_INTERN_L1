package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	// мапа с данными
	data := make(map[int]int)

	// создаем mu для синхронизации доступа к data
	mu := sync.Mutex{}

	for i := 0; i < 10; i++ {
		// увеличиваем счетчик wg для новой горутины
		wg.Add(1)
		key := i

		go func(key int) {
			// по завершению выполнения горутины уменьшаем счетчик wg
			defer wg.Done()

			// блокируем mu для синхронизации доступа к data (другие горутины не имеют доступа до mu.Unlock)
			mu.Lock()
			// создаем новую пару в data
			data[key] = rand.Int()
			// разблокируем mu
			mu.Unlock()
		}(key)
	}
	// ждем выполнения всех горутин
	wg.Wait()

	// выводим данные
	for k, v := range data {
		fmt.Println("Key:", k, "Value:", v)
	}
}
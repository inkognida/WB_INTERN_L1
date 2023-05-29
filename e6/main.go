package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func StopCh() {
	// кан
	stopCh := make(chan struct{})

	go func() {
		for {
			select {
			case value, ok := <-stopCh:
				if !ok {
					fmt.Println("Channel closed")
					return
				}
				fmt.Println("Value:", value)
			}
		}
	}()

	stopCh <- struct{}{}
	close(stopCh)
}

func StopContext() {
	// создаем контекст с функцией cancel
	ctx, cancel := context.WithCancel(context.Background())
	// после завершения функии и вызываем cancel() для передачи сигнала в ctx
	defer cancel()

	go func() {
		for {
			select {
			// сигнал о завершении работы
			case <-ctx.Done():
				fmt.Println("Ctx cancel")
				return
			default:
				// выполняем горутину в случае отсутствия сигнала о завершении
				fmt.Println("Working...")
				// время для выполнения
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(time.Second)
}

func StopWg() {
	wg := sync.WaitGroup{}

	// канал для записи данных
	dataCh := make(chan string)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				// читаем из канала
				for v := range dataCh {
					fmt.Println(v)
				}

				fmt.Println("Goroutine closed")
				return
			}
		}()

		// записываем данные в канал
		dataCh <- "hello"
		dataCh <- "world"
	}

	// закрываем канал для завершения горутин
	close(dataCh)

	// ждем выполнения всех горутин
	wg.Wait()
}

func main() {
	// завершение горутины с использованием канала
	StopCh()
	// ожидаем выполнения
	time.Sleep(time.Second)

	// завершение горутины с использованием контекста
	StopContext()
	// ожидаем выполнения
	time.Sleep(time.Second*2)

	// заверешение горутины с использованием sync.Waitgroup и канала с данными
	StopWg()
	// ожидаем выполнения
	time.Sleep(time.Second)
}
package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	Workers = 10
)

func ReadWorkers(wg *sync.WaitGroup, dataCh chan int, shutdownCh chan struct{}) {
	for i := 0; i < Workers; i++ {
		// увеличиваем счетчик wg для новой горутины
		wg.Add(1)
		go func() {
			// по завершению выполнения горутины уменьшаем счетчик wg
			defer wg.Done()

			// в цикле обрабатываем данные из каналов
			for {
				// считываем данные или завершаем работу
				select {
					case data := <-dataCh:
						fmt.Println(data)
					case <-shutdownCh:
						return
				}
			}
		}()
	}
}

func GenerateData(dataCh chan int) {
	for {
		// генерируем случайно число
		data := rand.Int()

		// data записываем в канал
		dataCh <- data
	}
}

func main() {
	// канал для записи данных
	dataCh := make(chan int)

	// канал для контроля завершения
	shutdownCh := make(chan struct{})

	// создаем wg для контроля выполнения всех горутин
	wg := sync.WaitGroup{}

	// Канал для сигнала ctrl+c
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGINT)


	// запускаем генератор данных
	go GenerateData(dataCh)

	// запускаем чтение данных
	go ReadWorkers(&wg, dataCh, shutdownCh)

	// Горутина для обработки stopCh
	<-stopCh

	fmt.Println("Handle ctrl+c signal")

	// закрываем канал для завершения ReadWorkers
	close(shutdownCh)

	// ждем выполнения всех горутин
	wg.Wait()

	// закрываем канал для записи данных
	close(dataCh)
}
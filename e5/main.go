package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	TimeSeconds = 3
)

func main() {
	// канал для записи данных
	dataCh := make(chan int, 1)

	// запись в канал
	go func() {
		for {
			// генерируем случайно число
			data := rand.Int()
			fmt.Println("New value:", data)

			time.Sleep(time.Millisecond*100)

			// data записываем в канал
			dataCh <- data
		}
	}()

	// чтение из канала
	go func() {
		for {
			fmt.Println("Receive value:", <-dataCh)

			time.Sleep(time.Millisecond*100)
		}
	}()

	// Ждем N секунд
	time.Sleep(time.Second*TimeSeconds)

	// закрываем dataCh
	close(dataCh)
}
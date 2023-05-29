package main

import "fmt"

func main() {
	arr := []int{1,2,3,4,5,6}

	// канал с данными arr
	dataCh := make(chan int)

	// канал со значениями из dataCh * 2
	dataOperationCh := make(chan int)

	// записываем значения arr в канал dataCh
	go func() {
		for _, v := range arr {
			dataCh <- v
		}
		// закрываем канал для дальнейшего чтения из него
		close(dataCh)
	}()

	go func() {
		// читаем данные из канала и записываем v * 2 в dataOperationCh
		for v := range dataCh {
			dataOperationCh <- v*2
		}
		close(dataOperationCh)
	}()

	// читаем данные из канала dataOperationCh
	for v := range dataOperationCh {
		fmt.Println("Value of data operation channel:", v )
	}
}
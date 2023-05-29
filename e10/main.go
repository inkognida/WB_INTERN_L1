package main

import "fmt"

func main() {
	// данные
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// мапа для хранения
	groups := make(map[int][]float64)
	for _, v := range arr {
		// создаем ключ для хранения значений
		key := int(v) / 10
		groups[key*10] = append(groups[key*10], v)
	}

	// выводим результат
	fmt.Println(groups)
}
package main

import "fmt"

// Функция для обмена элементов массива по индексам
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// Функция для разделения массива и возврата индекса опорного элемента
func partition(arr []int, low, high int) int {
	pivot := arr[high] // Опорный элемент
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			swap(arr, i, j)
		}
	}

	swap(arr, i+1, high)
	return i + 1
}

// QuickSort функция сортировки
func QuickSort(arr []int, low, high int) {
	if low < high {
		// определяем опорный элемент
		pivot := partition(arr, low, high)

		// сортируем левую часть от pivot
		QuickSort(arr, low, pivot-1)
		// сортируем правую часть от pivot
		QuickSort(arr, pivot+1, high)
	}
}

func main() {
	arr := []int{9, 3, 1, 7, 5, 6, 2, 8, 4}

	fmt.Println("Исходный массив:", arr)

	QuickSort(arr, 0, len(arr)-1)

	fmt.Println("Отсортированный массив:", arr)
}

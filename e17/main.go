package main

import "fmt"

// QuickSort
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

// BinarySearch функция поиска
func BinarySearch(arr []int, elem int) int {
	l := 0
	r := len(arr)-1
	// итерируемся по arr
	for l < r {

		// находим середину (arr отсортирован)
		mid := (l + r) / 2
		// меняем границу для l если mid < elem (elem справа)
		if arr[mid] < elem {
			l = mid + 1
			// меняем границу для r если mid > elem (elem слева)
		} else if arr[mid] > elem {
			r = mid - 1
		} else {
			return mid
		}
	}

	// elem не найден
	return -1
}

func main() {
	arr := []int{4,5,1,3,6,10}
	// сортируем arr
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println("Массив:", arr, "Поиск:", 6)

	elem := BinarySearch(arr, 6)
	fmt.Println("Индекс элемента:", elem)

}
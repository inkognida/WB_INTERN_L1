package main

import "fmt"

func main() {
	// создаем мапу для хранения пересечений в виде ключей
	intersection := make(map[int]struct{})

	set1 := []int{1,2,4,5,6}
	set2 := []int{7,8,4,1,5}

	for _, v := range set1 {
		// добавляем значения set1 в качестве ключей
		intersection[v] = struct{}{}
	}

	result := make([]int, 0)

	for _, v := range set2 {
		// при пересечении - добавляем в result
		if _, ok := intersection[v]; ok {
			result = append(result, v)
		}
	}

	fmt.Println(result)
}
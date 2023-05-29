package main

import "fmt"

func main() {
	arr := []int{0,1,2,3,4,5}
	i := 5

	// если i за пределами или arr пустой
	if len(arr) == 0 || len(arr)-1 < i {
		fmt.Println("Wrong")
	}

	arr = append(arr[:i], arr[i + 1:]...)
	fmt.Println(arr)
}
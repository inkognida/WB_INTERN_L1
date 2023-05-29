package main

import "fmt"

func Reverse(s string) string {
	// создаем r для перестановки значений
	r := []rune(s)

	// итериуремся по r и меняем значения местами
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

func main() {
	s := "главрыба"
	fmt.Println(s, Reverse(s))
}
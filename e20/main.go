package main

import (
	"fmt"
	"strings"
)

func Reverse(s string) string {
	// создаем r для перестановки слов
	r := strings.Split(s, " ")

	// итериуремся по r и меняем значения местами
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return strings.Join(r, " ")
}

func main() {
	s := "snow dog sun"

	fmt.Println(s, Reverse(s))
}
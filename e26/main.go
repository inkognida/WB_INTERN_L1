package main

import (
	"fmt"
	"strings"
)

func UniqueSymbols(s string) bool {
	s = strings.ToLower(s)

	// seen уникальные символы строки
	seen := make(map[rune]struct{})

	for _, v := range s {
		// добавляем новый элемент
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}

		// если уже встречался - есть дубликаты
		} else {
			return false
		}
	}
	return true
}


func main() {
	fmt.Println(UniqueSymbols("abcd"))
	fmt.Println(UniqueSymbols("abCdefAaf"))
	fmt.Println(UniqueSymbols("aabcd"))
}
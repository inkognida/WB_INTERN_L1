package main

import "fmt"

// Set структура множества
type Set map[string]bool

// NewSet создает новый Set
func NewSet() Set {
	return make(Set)
}

// Add добавляет element в Set
func (s Set) Add(element string) {
	s[element] = true
}

// Remove удаляет element из Set
func (s Set) Remove(element string) {
	delete(s, element)
}

// Contains проверяет наличие element
func (s Set) Contains(element string) bool {
	_, exists := s[element]
	return exists
}

// Size возвращает размер Set
func (s Set) Size() int {
	return len(s)
}

// Elements возвращает элементы множества
func (s Set) Elements() []string {
	elements := make([]string, 0, len(s))
	for element := range s {
		elements = append(elements, element)
	}
	return elements
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	// создаем новое множество
	set := NewSet()
	for _, v := range arr {
		// добавляем элементы
		set.Add(v)
	}

	fmt.Println(set.Elements())
}


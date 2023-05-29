package main

import "fmt"

// Human структура с полями name (имя) и country (страна)
type Human struct {
	Name string
	Country string
}

// Set метод для изменения полей структуры Human
func (h *Human) Set(name string, country string) {
	h.Name = name
	h.Country = country
}

// Action структура, которая наследует Human (поля и методы)
type Action struct {
	Human
}

func main() {
	// создаем структуру Human, инициализруем поля и выводим их
	h := Human{}
	h.Set("Aboba", "aboba")
	fmt.Println(h.Name, h.Country)


	// создаем структуру Action, инициализруем поля и выводим их
	a := Action{Human: h}
	a.Set("action", "action")
	fmt.Println(a.Name, a.Country)
}
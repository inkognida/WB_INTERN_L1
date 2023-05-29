package main

import (
	"fmt"
	"reflect"
)

// Types типы данных
type Types struct {
	Int int
	String string
	Bool bool
	Channel chan string
}

// Type определяет тип val
func Type(val interface{}) reflect.Type{
	return reflect.TypeOf(val)

}

func main()  {
	types := Types{
		Int:     0,
		String:  "",
		Bool:    false,
		Channel: nil,
	}

	fmt.Println(Type(types.Channel))
	fmt.Println(Type(types.Int))
	fmt.Println(Type(types.String))
	fmt.Println(Type(types.Bool))
}
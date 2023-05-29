package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Задаем значения переменных a и b
	a := big.NewInt(0)
	b := big.NewInt(0)
	a.SetString("12381293819221338", 10)
	b.SetString("12831283721821983129371892371823783", 10)


	// Умножение
	mul := new(big.Int).Mul(a, b)
	fmt.Println("Умножение:", mul)

	// Деление
	div := new(big.Int).Div(b, a)
	fmt.Println("Деление:", div)

	// Сложение
	add := new(big.Int).Add(a, b)
	fmt.Println("Сложение:", add)

	// Вычитание
	sub := new(big.Int).Sub(b, a)
	fmt.Println("Вычитание:", sub)
}

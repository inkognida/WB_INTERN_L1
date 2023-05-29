package main

import "fmt"

func setBit(value int64, pos uint, bitValue int64) int64 {
	// Создаем маску с установленным битом в позиции pos
	mask := int64(1) << pos

	// Инвертируем маску
	mask = ^mask

	// Сбрасываем бит в позиции pos
	value &= mask

	// Устанавливаем новое значение бита в позиции pos
	value |= bitValue << pos

	return value
}
func main() {
	n := int64(10)
	bitValue := int64(1)

	n1 := setBit(n, 7, bitValue)
	n2 := setBit(n1, 7, int64(0))
	fmt.Println(n, n1, n2)
}
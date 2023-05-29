package main

//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
//var justString string
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//func main() {
//	someFunc()
//}

/*
	Утечка памяти, поскольку justString ссылается на новый срез, т.к
	оператор среза создает новый срез, который ссылается на оригинальный массив,
	и сохраняет ссылку на v, что может привести к утечке памяти.
*/

var justString string

func createHugeString(length int) string {
	s := make([]rune, length)
	return string(s)
}

func someFunc() {
	v := []rune(createHugeString(1 << 10))
	justString = string(v[:100])
}

func main() {
	someFunc()
}
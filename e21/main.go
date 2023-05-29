package main

import "fmt"

// интерфейс транспорта
type transport interface {
	GoTo()
}

// клиент с запросом о транспорте
type client struct {

}

func (c *client) start(t transport) {
	fmt.Println("client starts GoTo process")
	t.GoTo()
}

type boat struct {

}

func (b *boat) 	GoTo() {
	fmt.Println("boat is going...")
}

type car struct {

}

func (c *car) GoTo() {
	fmt.Println("car is going...")
}

// адаптер для машины для продолжения goto client
type carAdapter struct {
	carTransport *car
}

func (c *carAdapter) GoTo() {
	fmt.Println("Modify car for going to...")
	c.carTransport.GoTo()
}

func main() {
	// клиент
	c := &client{}
	// лодка
	b := &boat{}

	// начало движения
	c.start(b)

	// машина
	newCar := &car{}

	// адаптер newCar для продолжения goto
	carAdapt := &carAdapter{carTransport: newCar}

	// goto с новым транспортом
	c.start(carAdapt)
}
package main

import "fmt"

type TransportAdapter interface {
	Start()
}

//Класс машина
type Car struct{}

func (c *Car) CarStart() {
	fmt.Println("Car is start")
}

//Адаптер для машины
type CarAdapter struct {
	*Car
}

//Заводим машину через адаптер
func (adapter *CarAdapter) Start() {
	adapter.CarStart()
}

//Класс вертолёт
type Helicopter struct{}

func (h *Helicopter) HelicopterStart() {
	fmt.Println("Helicopter start")
}

//Адаптер для вертолёта
type HelicopterAdapter struct {
	*Helicopter
}

//Заводим вертолёт через адаптер
func (adapter *HelicopterAdapter) Start() {
	adapter.HelicopterStart()
}

func startTransport(transport TransportAdapter) {
	transport.Start()
}

func main() {
	//Создаём транспорт
	car := &CarAdapter{}
	helicopter := &HelicopterAdapter{}

	//Запускаем транспорт через функцию
	startTransport(car)
	startTransport(helicopter)
}

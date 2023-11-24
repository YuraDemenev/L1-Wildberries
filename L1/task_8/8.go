package main

import "fmt"

func SetBit(num int64, index int) {
	fmt.Println("До изменения")
	fmt.Println(num)
	fmt.Println("После изменения")

	//Создаём маску, у которой только index bit = 1
	mask := 1 << index
	//Приеняем маску устанавливая index bit в 1
	num = num | int64(mask)
	fmt.Println(num)
}

func ClearBit(num int64, index int) {
	fmt.Println("До изменения")
	fmt.Println(num)
	fmt.Println("После изменения")

	//Создаем маско где только index bit = 1
	mask := 1 << index
	//Инвертируем теперь index bit = 0
	mask = ^mask
	//Пременяем маску и index bit = 0
	num = num & int64(mask)
	fmt.Println(num)
}

func main() {
	var num int64 = 256
	index := 3

	fmt.Println("Set bit")
	SetBit(num, index)
	fmt.Println()
	fmt.Println("Clear bit")
	ClearBit(num, index)

}

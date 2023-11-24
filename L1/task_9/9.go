package main

import (
	"fmt"
	"sync"
)

func Writer(val int, wg *sync.WaitGroup, chInput chan int, chOutput chan int) {
	//Чтобы Wait group уменьшил счётчик на 1
	defer wg.Done()

	//Пишем в канал 1
	chInput <- val

	//Возводим в квадрат и пишем во 2 канал
	val *= val
	chOutput <- val
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//Создаём 2 канала
	chInput := make(chan int, len(nums))
	chOutput := make(chan int, len(nums))

	//Wait group чтобы все горутины отработали
	var wg sync.WaitGroup

	//В цикле передаём числа Writer
	for _, v := range nums {
		wg.Add(1)
		go Writer(v, &wg, chInput, chOutput)
	}
	wg.Wait()
	close(chOutput)

	//Получаем все чисал
	for val := range chOutput {
		fmt.Println(val)
	}

}

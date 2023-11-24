package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	sum := 0

	//Создаём waitgroup чтобы всё горутины отработали
	var wg sync.WaitGroup
	ch := make(chan int)

	for _, val := range nums {
		//Увеличиваем счётчик
		wg.Add(1)
		go func(squere int) {
			squere *= squere

			//Пишем в канал
			ch <- squere

		}(val)

		go func() {
			//Достём данные из канала
			sum += <-ch

			//Уменьшаем счетчик
			wg.Done()
		}()

	}

	wg.Wait()

	fmt.Fprintln(os.Stdout, sum)
}

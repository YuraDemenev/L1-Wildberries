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
	ch := make(chan int, len(nums))

	for _, val := range nums {
		//Увеличиваем счётчик
		wg.Add(1)
		go func(squere int) {
			squere *= squere

			//Пишем в канал
			ch <- squere
			wg.Done()
		}(val)
	}
	wg.Wait()
	close(ch)

	for v := range ch {
		sum += v
	}

	fmt.Fprintln(os.Stdout, sum)
}

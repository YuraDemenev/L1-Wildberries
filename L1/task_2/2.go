package main

import (
	"fmt"
	"sync"
)

func WithWaitGroup(nums []int) {
	//Создание Wait group чтобы все горутины отработали
	var wg sync.WaitGroup

	for _, val := range nums {
		wg.Add(1)

		go func(v int) {
			v *= v
			fmt.Println(v)
			wg.Done()
		}(val)

	}
	wg.Wait()
}

func WithChanel(nums []int) {
	ch := make(chan int)
	for _, v := range nums {
		go func(v int) {
			v *= v
			ch <- v
		}(v)
	}

	for i := 0; i < len(nums); i++ {
		fmt.Println(<-ch)
	}
	close(ch)
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	fmt.Println("Use Wait Group")
	WithWaitGroup(nums)

	fmt.Println("Use chanels")
	WithChanel(nums)
}

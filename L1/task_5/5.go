package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Reader(ctx context.Context, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for {
		select {
		case val := <-ch:
			fmt.Printf("Reader read %d\n", val)
		case <-ctx.Done():
			return
		}
	}

}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	fmt.Println("Write a time")
	//Получаем время
	var N int
	fmt.Scan(&N)

	//Создаём контекст чтобы завершить программу через N сек
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(N))
	defer cancel()

	wg.Add(1)
	go Reader(ctx, &wg, ch)

	check := true
	for check {
		select {
		case ch <- rand.Intn(1000):
		case <-ctx.Done():
			fmt.Println("End")
			check = false
		}
	}

	wg.Wait()
	close(ch)

}

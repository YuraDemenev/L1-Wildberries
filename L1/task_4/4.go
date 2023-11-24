package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
)

func worker(ctx context.Context, ch <-chan int, workerId int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d end\n", workerId)
			return
		case val := <-ch:
			fmt.Printf("Worker %d read number: %d\n", workerId, val)
		}
	}
}

func main() {
	var N int
	fmt.Println("Write a number of worker")
	fmt.Scan(&N)

	//Создаём канал для отправки чисел
	ch := make(chan int, N)

	//Создаём контекст для прерывания работы программы
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)

	//Следим за сигналом чтобы отключить программу
	signal.Notify(c, os.Interrupt)

	//Функция отмены горутин
	defer func() {
		signal.Stop(c)
		cancel()
	}()

	//Создаём воркеров
	for i := 0; i < N; i++ {
		go worker(ctx, ch, i)
	}

	//Пишем в канал
	check := true
	for check {
		select {
		case ch <- rand.Intn(1000):

		case <-c:
			fmt.Println("Stop program")
			check = false
		}
	}

}

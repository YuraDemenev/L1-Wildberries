package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Остановка с использованием канала
func StopWithChanell(stopCh chan struct{}) {

	for {
		select {
		case <-stopCh:
			fmt.Println("Stop with chanel is stopped")
			return
		default:
			val := rand.Intn(1000)
			fmt.Printf("number : %d\n", val)
		}
	}
}

// Остановка с использованием контекста
func StopWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stop with context is stopped")
			return
		default:
			val := rand.Intn(1000)
			fmt.Printf("number : %d\n", val)
		}
	}
}

func main() {
	//Для остановки через закрытие канала
	stopCh := make(chan struct{})

	go StopWithChanell(stopCh)
	time.Sleep(time.Second * 2)

	close(stopCh)
	time.Sleep(time.Second * 1)

	//Для остановки через контекст
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	go StopWithContext(ctx)
	time.Sleep(time.Second * 3)

}

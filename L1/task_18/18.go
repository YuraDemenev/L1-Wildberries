package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	//Создаём Mutex чтобы небыло data race
	mu sync.Mutex
}

func (c *Counter) Increment() {
	//Блокируем остальным горутинам доступ к
	//Увеличению счетчика
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func main() {
	//Создаём счетчик
	count := Counter{}

	//Создаём wait group чтобы все горутины отработали
	var wg sync.WaitGroup

	//В цикле создаём 1000 горутин
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count.Increment()
		}()
	}

	wg.Wait()
	fmt.Println(count.count)
}

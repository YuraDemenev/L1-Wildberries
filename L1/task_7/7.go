package main

import (
	"fmt"
	"sync"
)

func UseNormalMap(val int, myMap map[int]int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	//Блокируем для других горутин доступ к записи в map
	mu.Lock()
	myMap[val] = val
	mu.Unlock()
}

func UseSyncMap(val int, myMap *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	myMap.Store(val, val)
}

func main() {
	//Use a normal map
	myMap := make(map[int]int, 10)
	//Wait group чтобы все горутины отработали
	var wg sync.WaitGroup
	//Mutex чтобы небыло data race
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go UseNormalMap(i, myMap, &wg, &mu)
	}
	wg.Wait()

	fmt.Println("Add data in normal map")
	for _, v := range myMap {
		fmt.Printf("%d ", v)
	}
	fmt.Print(" \n")
	fmt.Print(" \n")

	//Use a sync map
	syncMap := sync.Map{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go UseSyncMap(i, &syncMap, &wg)
	}
	wg.Wait()

	fmt.Println("Add data in sync map")
	syncMap.Range(func(key, value any) bool {
		fmt.Printf("%d ", value)
		return true
	})
}

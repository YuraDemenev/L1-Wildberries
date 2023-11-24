package main

import (
	"fmt"
	"math/rand"
)

func main() {
	count := rand.Intn(6)
	workers := []string{"Product Manager", "Program Manager", "Architect", "Developer", "Tester", "User Experience", "Release & Operations"}

	dictionar := make(map[int]int)
	for len(dictionar) != count {
		proc := rand.Intn(11)
		index := rand.Intn(7)

		_, ok := dictionar[index]
		if !ok {
			dictionar[index] = proc
		}
	}

	for key, val := range dictionar {
		fmt.Printf("Worker %s, have proc %d", workers[key], val)
		fmt.Println()
	}
}

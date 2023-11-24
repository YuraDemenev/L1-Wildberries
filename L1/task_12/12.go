package main

import "fmt"

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree"}
	dictionary := make(map[string]int)

	for _, str := range slice {
		dictionary[str]++
	}

	for key := range dictionary {
		fmt.Print(key + " ")
	}
}

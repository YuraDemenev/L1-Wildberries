package main

import "fmt"

func getIntersection(numsFirst []int, numsSecond []int) {
	dictionary := make(map[int]struct{})
	result := make([]int, 0)

	for _, v := range numsFirst {
		dictionary[v] = struct{}{}
	}

	for _, v := range numsSecond {
		_, ok := dictionary[v]
		if ok {
			result = append(result, v)
		}
	}

	fmt.Println("Пересекающиеся значения :")
	fmt.Println(result)

}

func main() {
	numsFirst := []int{3, 1, 6, 55, 55, 101, 256, 1024}
	numsSecond := []int{100, 99, 55, 1024, 4, 5, 5, 5, 4}
	getIntersection(numsFirst, numsSecond)
}

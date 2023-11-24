package main

import "fmt"

func quickSort(slice []int, start int, end int) {
	//Base case
	if end <= start {
		return
	}
	//Находим локальный pivot который раздели
	//Массив на 2 других массива
	pivot := partiton(slice, start, end)
	quickSort(slice, start, pivot-1)
	quickSort(slice, pivot+1, end)
}

func partiton(slice []int, start int, end int) int {
	//Перем число находящиеся в конце
	//и ищем место где оно должно быть
	pivot := slice[end]
	i := start - 1

	//Передвигаем числа которые меньше pivot
	//В лево от него а число большие в право
	for j := start; j <= end-1; j++ {
		if slice[j] < pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
	//Когда цикл закончился сдвигаем i на 1 и
	// Перемещаем туда pivot
	i++
	slice[i], slice[end] = slice[end], slice[i]

	return i
}

func main() {
	fmt.Println("До сортировки")
	slice := []int{8, 2, 4, 1, 6, 7, 3, 5, 9}
	fmt.Println(slice)

	fmt.Println()
	fmt.Println("После сортировки")
	quickSort(slice, 0, len(slice)-1)
	fmt.Println(slice)
}

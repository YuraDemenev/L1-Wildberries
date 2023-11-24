package main

import (
	"fmt"
	"sort"
)

func binarySearch(slice []int, target int) {
	//Сортируем слайс.Бинарный поиск работает
	//Только в отсортированном массиве
	sort.Ints(slice)

	//Создаём два поинта указываеющие на края массива
	l, r := 0, len(slice)-1
	for l <= r {
		//Получаем середину
		mid := (l + r) / 2
		if slice[mid] == target {
			fmt.Println("Yes this number in slice")
			return
		}
		//Если значение больше или меньше значит можно отбросить всё что
		//за индексом mid и его тоже
		if slice[mid] > target {
			//Отбрасываем всё справа
			r = mid - 1
		} else {
			//Отбрасываем всё слева
			l = mid + 1
		}

	}
	fmt.Println("No this target not in slice")
	return
}

func main() {
	var target int
	slice := []int{8, 2, 4, 1, 6, 7, 3, 5, 9}
	fmt.Println("Write a target")
	fmt.Scan(&target)

	binarySearch(slice, target)
}

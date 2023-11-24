package main

import "fmt"

func deleteElementFromSlice(nums []int, index int) []int {
	//Присваиваем локальнома массиву все элементы до индекас
	locNums := nums[:index]

	//Добваляем все элементы после индекса
	locNums = append(locNums, nums[index+1:]...)
	return locNums
}

func main() {
	var index int
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("Write a index")
	fmt.Scan(&index)

	//Проверяем находится ли индекс в диапозоне
	if index < 0 || index >= len(nums) {
		fmt.Println("Index out of range")
		return
	}

	fmt.Println("Before delete:")
	fmt.Println(nums)

	nums = deleteElementFromSlice(nums, index)
	fmt.Println("After delete:")
	fmt.Println(nums)
}

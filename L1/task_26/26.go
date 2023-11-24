package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abcd"
	//Создаём словарь
	dictionary := make(map[rune]int)

	strings.ToLower(str)

	check := false
	for _, val := range str {
		//Проверяем есть ли такое значение в словаре
		_, ok := dictionary[val]
		if ok {
			check = true
			break
		}
		dictionary[val]++
	}

	if check {
		fmt.Println("not unique")
	} else {
		fmt.Println("unique")
	}

}

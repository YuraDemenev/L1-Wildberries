package main

import "fmt"

func reverseString(str string) string {
	runes := []rune(str)
	//Делим длину на 2 чтобы перевернуть строку инача строка 2 раза
	//Перевернётся и останется прежней
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}
	return string(runes)
}

func main() {
	str := "главрыба Test"
	fmt.Println(reverseString(str))
}

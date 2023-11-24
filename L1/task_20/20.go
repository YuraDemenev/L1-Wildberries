package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"

	//Разбиваем строки по пробелу
	sliceStrs := strings.Split(str, " ")

	//string builder для эффективной работы со строками
	sb := strings.Builder{}

	//Цикл до 0 элемента чтобы не добавлять лишний пробел
	for i := len(sliceStrs) - 1; i > 0; i-- {
		sb.WriteString(sliceStrs[i] + " ")
	}
	sb.WriteString(sliceStrs[0])

	fmt.Println(sb.String())
}

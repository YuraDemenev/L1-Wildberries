// var justString string

// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }

// func main() {
// 	someFunc()
// }

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// 1) V[:100] может вызвать панику если в строке нет 100 байт
// 2)Если в строке символы не из ASCII table то они занимают
// больше чем 1 байт => вернёт не 100 символов а меньше
// 3) глобальная переменная не нужна
func createHugeString(length int) string {
	//Создаём string builder для эффективного добавления данных в строку
	sb := strings.Builder{}
	for i := 0; i < length; i++ {
		str := strconv.Itoa(rand.Intn(10))
		sb.WriteString(str)
	}
	return sb.String()
}

func someFunc() {
	var justString string
	justString = createHugeString(1 << 10)
	fmt.Println(justString)
}

func main() {
	someFunc()
}

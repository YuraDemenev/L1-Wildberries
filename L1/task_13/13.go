package main

import "fmt"

func SwapTwoNumbers(a, b int) (int, int) {
	return b, a
}

func main() {
	//Метод a,b=b,a
	fmt.Println("method a,b=b,a")
	a, b := 20, 30
	fmt.Printf("a:%d, b:%d\n", a, b)

	a, b = b, a
	fmt.Printf("a:%d, b:%d\n\n", a, b)

	//Использование return in fucntion"
	fmt.Println("Use return in fucntion")
	a, b = 20, 30
	fmt.Printf("a:%d, b:%d\n", a, b)

	a, b = SwapTwoNumbers(a, b)
	fmt.Printf("a:%d, b:%d\n\n", a, b)

	//Использование + -
	fmt.Println("Use + -")
	a, b = 20, 30
	fmt.Printf("a:%d, b:%d\n", a, b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("a:%d, b:%d\n\n", a, b)

	//Использование XOR
	fmt.Println("Use XOR")
	a, b = 20, 30
	fmt.Printf("a:%d, b:%d\n", a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("a:%d, b:%d\n", a, b)

}

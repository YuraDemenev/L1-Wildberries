package main

import (
	"fmt"
	"math/big"
)

func main() {
	first, second, result := new(big.Int), new(big.Int), new(big.Int)
	first.SetString("24000000000000000000000000000000000000000000000000000000000000", 10)
	second.SetString("30000000000000000000000000000000000000000000000000000000000000", 10)

	//Умножение
	result.Mul(first, second)
	fmt.Printf("Mul :%d\n", result)

	//Деление
	result.Div(first, second)
	fmt.Printf("Div :%d\n", result)

	//Сумма
	result.Add(first, second)
	fmt.Printf("Add :%d\n", result)

	//Разность
	result.Sub(first, second)
	fmt.Printf("Sub :%d\n", result)

}

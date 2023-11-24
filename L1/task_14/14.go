package main

import (
	"fmt"
	"reflect"
)

func RecognizeType(val interface{}) {
	switch type_ := val.(type) {
	case int:
		fmt.Printf("%d is int\n", type_)
	case string:
		fmt.Printf("%s is string\n", type_)
	case bool:
		fmt.Printf("%v is bool\n", type_)
	case chan int:
		fmt.Printf("%v is int chan\n", type_)
	case float64:
		fmt.Printf("%v is float", type_)
	default:
		fmt.Printf("%v unexpected type\n", type_)
	}
}

func main() {
	types := []interface{}{make(chan int), "Hello world!", 99, true, 2.3}

	fmt.Println("Use switch case")
	//Use switch case
	for _, val := range types {
		RecognizeType(val)
	}
	fmt.Print("\n\n")

	fmt.Println("Use typeOf")
	//Use typeOf
	for _, val := range types {
		fmt.Printf("%v is %v\n", val, reflect.TypeOf(val))
	}

}

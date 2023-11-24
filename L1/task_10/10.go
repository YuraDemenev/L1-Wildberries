package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	tempeture := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	//Создаем словарь для сохранения данных
	dictionary := make(map[int][]float64)

	for _, val := range tempeture {
		//Округляем число
		roundValue := math.Trunc(val/10) * 10

		//Проверяем есть ли такое в словаре
		mySlice, ok := dictionary[int(roundValue)]
		if !ok {
			//Создаём слайс и добавляем в словарь
			mySlice := []float64{val}
			dictionary[int(roundValue)] = mySlice

		} else {
			//Добавляем в слайс данные
			mySlice = append(mySlice, val)
			dictionary[int(roundValue)] = mySlice
		}
	}

	//Выводим данные
	for key, mySlice := range dictionary {
		fmt.Printf("%d:{", key)
		for _, v := range mySlice {
			//Приводим к более читаемому формату
			str := strconv.FormatFloat(v, 'f', -1, 64)
			fmt.Print(str + " ")
		}
		fmt.Printf("}, ")
	}
}

package main

import (
	"fmt"
)

func main() {
	exampleNames := []string{
		"#1 Get professional development",
		"#2 Get bachelor degree",
		"#3 Get masters degree",
		"#4 Get all degree",
		"#5 Get Salary",
		"#6 Go to vacation",
		"#0 Back"}

	checkVal := 1

	//Создания объекта через реализованный конструктор
	human := NewHuman(50000, 50000)

	for checkVal != 0 {
		fmt.Println()

		//Вывод в консоль все возможные действия
		for _, val := range exampleNames {
			fmt.Println(val)
		}

		//Получение данных из консли
		var checkVal int
		fmt.Scanln(&checkVal)

		switch checkVal {
		case 0:
			fmt.Println()
			return
		case 1:
			human.GetProfessionalDevelopment()
		case 2:
			human.GetBachelorDegree()
		case 3:
			human.GetMasterDegree()
		case 4:
			human.GetAllDegree()
		case 5:
			human.GetSalary()
		case 6:
			human.GoToVacation()
		default:
			fmt.Println("Wrong answer")
		}
	}
	return
}

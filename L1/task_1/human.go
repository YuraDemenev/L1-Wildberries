package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Human struct {
	salary              int
	educations          []string
	countBachelorDegree int
	countMasterDegree   int
	countMoney          int
}

func NewHuman(mySalary int, myCountMoney int) *Human {
	slice := make([]string, 0)
	human := &Human{
		salary:              mySalary,
		educations:          slice,
		countBachelorDegree: 0,
		countMasterDegree:   0,
		countMoney:          myCountMoney,
	}

	return human
}

func (h *Human) GetProfessionalDevelopment() {
	h.salary += 10000
	locStr := fmt.Sprintf("Now salary: %d", h.salary)
	fmt.Println(locStr)

	fmt.Println("Now your salary : " + locStr)
}

func (h *Human) GetBachelorDegree() {
	locStr := fmt.Sprintf("bachelor degree №%d", h.countBachelorDegree+1)
	h.educations = append(h.educations, locStr)
	h.countBachelorDegree++
	countString := strconv.Itoa(h.countBachelorDegree)

	fmt.Println("Now your have got " + countString + " bachelor degree")
}

func (h *Human) GetMasterDegree() {
	locStr := fmt.Sprintf("master degree №%d ", h.countMasterDegree+1)
	h.educations = append(h.educations, locStr)
	h.countMasterDegree++
	countString := strconv.Itoa(h.countMasterDegree)

	fmt.Println("Now your have got " + countString + " master degree")
}

func (h *Human) GetAllDegree() {
	fmt.Println("Your degrees")
	for _, val := range h.educations {
		fmt.Println(val)
	}
}

func (h *Human) GetSalary() {
	h.countMoney += h.salary
	fmt.Println("You get a salary")
	fmt.Printf("Your money %d \n", h.countMoney)
}

func (h *Human) GoToVacation() {
	count := int(rand.Intn(100001))
	if count > h.countMoney {
		fmt.Println("You have not got money")
		return
	}
	h.countMoney -= count
	fmt.Println("You go to vacation")
	fmt.Printf("Your money %d \n", h.countMoney)
}

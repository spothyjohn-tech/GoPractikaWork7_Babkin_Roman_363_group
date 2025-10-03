package main

import (
	"fmt"
)

type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float64
}

type Department struct {
	ID        int
	Name      string
	employees []Employee
}

func (d *Department) AddEmployee(name string, position string, salary float64) {
	newEmployee := Employee{
		ID:       len(d.employees) + 1, 
		Name:     name,
		Position: position,
		Salary:   salary,
	}
	d.employees = append(d.employees, newEmployee)
	fmt.Println("Добавлен сотрудник:", name)
}

func (d *Department) RemoveEmployee() {
	for _, emp := range d.employees {
		fmt.Println("ID:", emp.ID, "Имя:", emp.Name, "Должность:", emp.Position)
	}
	fmt.Println("Введите ID сотрудника для удаления:")
	var choice int
	fmt.Scanln(&choice)
	for i, emp := range d.employees {
		if emp.ID == choice {
			fmt.Println("Сотрудник", emp.Name, "удалён из отдела.")
			d.employees = append(d.employees[:i], d.employees[i+1:]...)
			return
		}
	}
	fmt.Println("Сотрудник с таким ID не найден.")
}

func (d *Department) CalculateSalaryFund() float64 {
	var total float64
	for _, emp := range d.employees {
		total += emp.Salary
	}
	return total
}

func (d *Department) GetEmployeesByPosition() {
	fmt.Println("Введите должность для поиска:")
	var position string
	fmt.Scanln(&position)
	found := false
	for _, emp := range d.employees {
		if emp.Position == position {
			fmt.Println("--------------------------------------------------")
			fmt.Println("ID:", emp.ID, "Имя:", emp.Name, "Должность:", emp.Position, "Оклад:", emp.Salary)
			fmt.Println("--------------------------------------------------")
			found = true
		}
	}

	if !found {
		fmt.Println("Сотрудники с такой должностью не найдены.")
	}
}

func (d *Department) PrintAllEmployees() {
	if len(d.employees) == 0 {
		fmt.Println("В отделе нет сотрудников.")
		return
	}
	for _, emp := range d.employees {
		fmt.Println("--------------------------------------------------")
		fmt.Println("ID:", emp.ID, "Имя:", emp.Name, "Должность:", emp.Position, "Оклад:", emp.Salary)
		fmt.Println("--------------------------------------------------")
	}
}

func main() {
	dept := Department{
		ID:   1,
		Name: "Отдел разработки",
	}
	dept.AddEmployee("Иван Иванов", "Разработчик", 100000)
	dept.AddEmployee("Пётр Петров", "Тестировщик", 80000)
	dept.AddEmployee("Светлана Смирнова", "Разработчик", 110000)
	dept.PrintAllEmployees()
	dept.RemoveEmployee()
	dept.PrintAllEmployees()
	dept.GetEmployeesByPosition()
	total := dept.CalculateSalaryFund()
	fmt.Println("Общий зарплатный фонд отдела:", total)
}

package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	// Example usage of the Employee and Company structs
	company := Company{}
	employee1 := Employee{Name: "Alice", Age: 30, Position: "Developer", Salary: 70000}
	employee2 := Employee{Name: "Bob", Age: 25, Position: "Designer", Salary: 60000}
	employee3 := Employee{Name: "Charlie", Age: 35, Position: "Developer", Salary: 80000}
	employee4 := Employee{Name: "David", Age: 51, Position: "Manager", Salary: 90000}
	employee5 := Employee{Name: "Eve", Age: 22, Position: "Intern", Salary: 30000}
	employee6 := Employee{Name: "Frank", Age: 42, Position: "Manager", Salary: 95000}

	company.AddEmployee(employee1)
	company.AddEmployee(employee2)
	company.AddEmployee(employee3)
	company.AddEmployee(employee4)
	company.AddEmployee(employee5)
	company.AddEmployee(employee6)

	fmt.Println("Single employee info:")
	Print(employee1)

	fmt.Println("Company employee info")
	Print(&company)

	fmt.Println("Average age of employees in the company:", strconv.FormatFloat(company.CalculateAverageAge(), 'f', 2, 64))

	employee, err := company.SearchByName("Alice")
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Found employee:", employee.Info())
	}

	employee2, err2 := company.SearchByName("FooBar")
	if err2 != nil {
		fmt.Println("Error", err2)
	} else {
		fmt.Println("Found employee:", employee2.Info())
	}

	report := company.GeneratePositionSalaryReport()
	fmt.Println("Average salary by position:")
	for position, avgSalary := range report {
		fmt.Printf("Position: %s, Average Salary: $%.2f\n", position, avgSalary)
	}
}

type Printable interface {
	Print()
}

func Print(p Printable) {
	// Dummy method, just to show how we can use the Printable interface
	p.Print()
}

type Employee struct {
	Name     string
	Age      int
	Position string
	Salary   float64
}

func (e Employee) Info() string {
	// A bit pointless but according to the task, we need to implement this method as well, basically duplicating the Print method
	return "Employee: " + e.Name + ", Age: " + strconv.Itoa(e.Age) + ", Position: " + e.Position + ", Salary: $" + strconv.FormatFloat(e.Salary, 'f', 2, 64)
}

func (e Employee) Print() {
	fmt.Println(e.Info())
}

type Company struct {
	Employees []Employee
}

func (c *Company) Print() {
	for _, emp := range c.Employees {
		Print(emp)
	}
}

func (c *Company) AddEmployee(e Employee) {
	c.Employees = append(c.Employees, e)
}

func (c *Company) SearchByName(name string) (Employee, error) {
	// This method returns the FIRST found employee with the given name or an error if not found
	for _, emp := range c.Employees {
		if emp.Name == name {
			return emp, nil
		}
	}
	return Employee{}, errors.New("employee " + name + "not found in the company")
}

func (c *Company) CalculateAverageAge() float64 {
	if len(c.Employees) == 0 {
		return 0
	}
	totalAge := float64(0)
	for _, emp := range c.Employees {
		totalAge += float64(emp.Age)
	}
	return totalAge / float64(len(c.Employees))
}

func (c *Company) GeneratePositionSalaryReport() map[string]float64 {
	// Count salary and number of employees for each position separately
	salarySum := make(map[string]float64)
	count := make(map[string]int)
	for _, emp := range c.Employees {
		if emp.Position == "" {
			// Skip employees without a position
			continue
		}
		salarySum[emp.Position] += emp.Salary
		count[emp.Position]++
	}

	// Make a report with average salary for each position
	report := make(map[string]float64)
	for position, totalSalary := range salarySum {
		averageSalary := totalSalary / float64(count[position])
		report[position] = averageSalary
	}
	return report
}

package main

import (
	"errors"
	"fmt"
)

func main() {
	printSeparator()

	// TASK 1
	a := 42
	b := 69
	fmt.Printf("Result of sum of %d and %d is %d\n", a, b, calculateSum(a, b))
	printSeparator()

	a = 35
	b = 4
	result, remainder, err := divideWithRemainder(a, b)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Result of division of %d by %d is %d with remainder %d\n", a, b, result, remainder)
	}
	printSeparator()

	a = 60
	b = 0
	result, remainder, err = divideWithRemainder(a, b)
	if err != nil {
		fmt.Printf("Result of division of %d by %d is error: %s\n", a, b, err)
	} else {
		fmt.Printf("Result of division of %d by %d is %d with remainder %d\n", a, b, result, remainder)
	}
	printSeparator()

	s := convertFromFunction(convertIntToString, 5)
	fmt.Printf("Converted 5 to string: %s\n", s)
	printSeparator()

	// TASK 2
	arr := []int{4, 8, 15, 16, 23}
	fmt.Println("Original array:")
	printArray(arr)
	printSeparator()

	fmt.Println("Array after doubling all values:")
	printArray(doubleAllValues(arr))
	printSeparator()

	fmt.Println("Original array should not change:")
	printArray(arr)
	printSeparator()

	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	fmt.Println("Concatenated slices:", append(slice1, slice2...))
	printSeparator()

	// TASK 3
	ageMap := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
	}
	printMap(ageMap)
	printSeparator()

	items := []string{"hello", "world", "hello", "golang", "world", "golang", "i", "love", "golang"}
	counts := countStrings(items)
	printMap(counts)
	printSeparator()

	delete(ageMap, "Bob")
	fmt.Println("Map after deleting Bob from it:")
	printMap(ageMap)
	printSeparator()

	// TASK 4
	employee := Person{Name: "John", Salary: 1_000}
	fmt.Printf("Employee: %s, Salary: %d\n", employee.Name, employee.Salary)
	printSeparator()
}

func printSeparator() {
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println()
}

func calculateSum(a int, b int) int {
	return a + b
}

func divideWithRemainder(a int, b int) (result int, remainder int, err error) {
	if b == 0 {
		return 0, 0, errors.New("division by zero")
	}
	return a / b, a % b, nil
}

func convertIntToString(n int) string {
	switch {
	case n < 1:
		return "small or zero"
	case n <= 10:
		return "average"
	default:
		return "big"
	}
}

func convertFromFunction(fn func(int) string, n int) string {
	fmt.Printf("Converting %d to string using passed function\n", n)
	r := fn(n)
	return r
}

func printArray(arr []int) {
	fmt.Println("Array contents:")
	for i, v := range arr {
		fmt.Printf("Index %d: %d\n", i, v)
	}
}

func doubleAllValues(arr []int) []int {
	newArray := make([]int, len(arr))
	for i, v := range arr {
		newArray[i] = v * 2
	}
	return newArray
}

func printMap(m map[string]int) {
	fmt.Println("Map contents:")
	for k, v := range m {
		fmt.Printf("%s: %d\n", k, v)
	}
}

func countStrings(items []string) map[string]int {
	counts := make(map[string]int)

	for _, item := range items {
		counts[item]++
	}
	return counts
}

type Person struct {
	Name   string
	Salary int
}

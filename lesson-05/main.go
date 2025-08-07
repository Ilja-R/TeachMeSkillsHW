package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for {
		fmt.Println("Let us divide two numbers!")

		result, err := processDivision()
		if err != nil {
			fmt.Printf("Error: %s. Please retry!\n", err)
			continue
		}

		printResult(result)

		fmt.Println("Type 'exit' and hit enter to stop or hit Enter to continue:")
		input, err := getStringInput()
		if err != nil {
			fmt.Printf("Error: %s. Starting new cycle!\n", err)
			continue
		}
		if strings.EqualFold(input, "exit") {
			fmt.Println("Goodbye!")
			break
		}
	}
}

func processDivision() (int, error) {
	dividend, err := getNumber("dividend")
	if err != nil {
		return 0, err
	}

	divisor, err := getNumber("divisor")
	if err != nil {
		return 0, err
	}

	return divide(dividend, divisor)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func printResult(result int) {
	fmt.Printf("Result of whole number division is : %d\n", result)
	switch {
	case result > 10:
		fmt.Println("The result is big")
	case result >= 1:
		fmt.Println("The result is average")
	default:
		fmt.Println("The result is small or zero")
	}
}

func getStringInput() (string, error) {
	if !scanner.Scan() {
		return "", errors.New("failed to read input")
	}
	return strings.TrimSpace(scanner.Text()), nil
}

func getNumber(label string) (int, error) {
	for {
		fmt.Printf("Enter %s: ", label)
		text, err := getStringInput()
		if err != nil {
			return 0, err
		}
		num, err := strconv.Atoi(text)
		if err == nil {
			return num, nil
		}
		fmt.Printf("Invalid number '%s'. Please try again.\n", text)
	}
}

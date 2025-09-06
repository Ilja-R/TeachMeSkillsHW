package main

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func loadUsers(filename string) ([]User, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func writeToExcel(filename string, users []User) error {
	f := excelize.NewFile()
	sheet := "Sheet1"

	f.SetCellValue(sheet, "A1", "Name")
	f.SetCellValue(sheet, "B1", "Age")

	for i, user := range users {
		row := i + 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), user.Name)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), user.Age)
	}

	if err := f.SaveAs(filename); err != nil {
		return err
	}

	return nil
}

func main() {
	const inputFile = "users.json"
	const outputFile = "report.xlsx"

	users, err := loadUsers(inputFile)
	if err != nil {
		fmt.Println("Error loading users:", err)
		os.Exit(1)
	}

	if err := writeToExcel(outputFile, users); err != nil {
		fmt.Println("Error writing Excel report:", err)
		os.Exit(1)
	}

	fmt.Println("Report successfully generated:", outputFile)
}

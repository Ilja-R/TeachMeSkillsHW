package main

import (
	"encoding/json"
	"encoding/xml"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Book struct {
	Title string `xml:"title"`
	Year  int    `xml:"year"`
}

type Library struct {
	Books []Book `xml:"book"`
}

func main() {
	users, err := loadUsers("users.json")
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		println("Name:", user.Name, "Age:", user.Age+1)
	}

	library, err := loadLibrary("books.xml")
	if err != nil {
		panic(err)
	}
	for _, book := range library.Books {
		println("Title:", book.Title, "Year:", book.Year+1)
	}
}

func loadLibrary(path string) (*Library, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var library Library
	err = xml.Unmarshal(data, &library)
	if err != nil {
		return nil, err
	}
	return &library, nil
}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

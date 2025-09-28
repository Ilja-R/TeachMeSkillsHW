package models

import "fmt"

type Employee struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
	Age   int    `json:"age" db:"age"`
}

func (u Employee) CacheKey() string {
	return fmt.Sprintf("employee_%d", u.ID)
}

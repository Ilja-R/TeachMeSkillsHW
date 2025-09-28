package dto

// For now keep it simple, but may need different logic for Create request, Update request and so on

type EmployeeCreateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

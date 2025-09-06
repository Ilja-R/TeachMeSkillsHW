package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Greeting struct {
	Name string `json:"name"`
}

type GreetingResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed!", http.StatusMethodNotAllowed)
		return
	}
	var g Greeting
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	w.Header().Set("Content-Type", "application/json")
	if err := dec.Decode(&g); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ErrorResponse{err.Error()})
		return
	}
	if strings.TrimSpace(g.Name) == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ErrorResponse{Error: "name is required"})
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(GreetingResponse{fmt.Sprintf("Hello, %s!", g.Name)}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ErrorResponse{err.Error()})
		return
	}
}

func main() {
	http.HandleFunc("/greet", greetHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

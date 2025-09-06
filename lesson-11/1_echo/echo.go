package main

import (
	"fmt"
	"io"
	"net/http"
)

/*
POST /echo
Content-Type: text/plain
Привет, сервер!
Ответ:
Привет, сервер!
*/
func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed!", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading body!", http.StatusBadRequest)
		return
	}
	_, err = w.Write(body)
	if err != nil {
		http.Error(w, "Error writing answer!", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/echo", echoHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

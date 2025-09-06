package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	const logFile = "app.log"

	file, err := openFile(logFile)
	if err != nil {
		fmt.Println("Error opening/creating log file:", err)
	}
	defer file.Close()

	configureLogger(file)
	startLogging()
}

func openFile(filename string) (*os.File, error) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return os.Create(filename)
	}
	return os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
}

func configureLogger(file *os.File) {
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func startLogging() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Logging started...")
	fmt.Println("Enter text to log or type 'exit' to quit:")
	for scanner.Scan() {
		fmt.Println("Enter text to log or type 'exit' to quit:")
		line := scanner.Text()
		if line == "exit" {
			fmt.Println("Exiting program.")
			break
		}
		log.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

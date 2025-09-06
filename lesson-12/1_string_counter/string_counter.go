package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {

	words, err := getWordsFromFile("input.txt")
	if err != nil {
		fmt.Println("Error getting words from file:", err)
		return
	}

	freq := make(map[string]int)
	for _, word := range words {
		freq[word]++
	}

	err = mapToCSV("output.csv", freq, "word", "frequency")
	if err != nil {
		fmt.Println("Error saving to CSV:", err)
	} else {
		fmt.Println("Word frequencies saved to output.csv")
	}
}

func getWordsFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return nil, err
	}
	defer file.Close()

	var builder strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
		builder.WriteString(" ")
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return nil, err
	}

	text := strings.ToLower(strings.TrimSpace(builder.String()))
	words := strings.Fields(text)
	return words, nil
}

func mapToCSV[T any](filename string, frequency map[string]T, keyName string, valueName string) error {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{keyName, valueName}); err != nil {
		fmt.Println("Error writing header:", err)
		return err
	}

	for word, count := range frequency {
		record := []string{word, fmt.Sprintf("%d", count)}
		if err := writer.Write(record); err != nil {
			fmt.Println("Error writing record:", err)
			return err
		}
	}

	return nil
}

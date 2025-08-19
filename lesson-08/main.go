package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {

	var vg sync.WaitGroup
	var oddNumbers []int
	var evenNumbers []int

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			evenNumbers = append(evenNumbers, i)
		} else {
			oddNumbers = append(oddNumbers, i)
		}
	}

	startTime := time.Now()
	fmt.Println("Starting to print numbers...")
	vg.Add(2)
	go printNumbers(evenNumbers, "Even:", &vg)
	go printNumbers(oddNumbers, "Odd:", &vg)
	vg.Wait()

	fmt.Println("All numbers printed!")
	elapsedTimeTask1 := time.Since(startTime).Milliseconds()
	fmt.Printf("Elapsed time %d milliseconds\n", elapsedTimeTask1)
	fmt.Println("-----------------------------------")

	var vgFiles sync.WaitGroup
	startTime = time.Now()
	fmt.Println("Starting to load files...")
	vgFiles.Add(3)
	go loadFile("file1.txt", time.Second, &vgFiles)
	go loadFile("file2.txt", 2*time.Second, &vgFiles)
	go loadFile("file3.txt", 500*time.Millisecond, &vgFiles)
	vgFiles.Wait()

	fmt.Println("All files loaded!")
	elapsedTimeTask2 := strconv.FormatFloat(time.Since(startTime).Seconds(), 'f', 3, 64)
	fmt.Printf("Total elapsed time to load files: %s seconds\n", elapsedTimeTask2)
}

func printNumbers(numbers []int, prefix string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, number := range numbers {
		fmt.Println(prefix + strconv.Itoa(number))
		time.Sleep(time.Millisecond) // Simulate some work, otherwise one goroutine ends before the other
	}
}

// Not an actual file loading, just an "imitation"
func loadFile(fileName string, waitTime time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	startTime := time.Now()
	fmt.Println("Loading file:", fileName)
	time.Sleep(waitTime)
	elapsedTime := strconv.FormatFloat(time.Since(startTime).Seconds(), 'f', 3, 64)
	fmt.Printf("File %s loaded in %s seconds\n", fileName, elapsedTime)
}

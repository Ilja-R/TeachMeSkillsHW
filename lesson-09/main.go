package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// Task 1: Channels
	channels()
	channelsBuff()
	printSeparator()

	// Task 2: Select
	selectTest()
	printSeparator()

	// Task 3: Timer and Timout
	timerAndTimeout()
	printSeparator()

	// Task 4: Async gathering of data
	asyncGathering()
	printSeparator()

	// Task 5: Context
	contextExample()
	printSeparator()

	// Task 6: Worker Pool
	workerPool()
	printSeparator()

	// Task 7: WaitGroup
	waitGroup()
	printSeparator()

	// Task 8: Race Condition
	raceCondition()
	printSeparator()
}

func raceCondition() {
	i := 0
	const numGoroutines = 1000
	var wg = sync.WaitGroup{}
	//var mu = sync.Mutex{}
	wg.Add(numGoroutines)

	for j := 0; j < numGoroutines; j++ {
		go func() {
			defer wg.Done()
			//mu.Lock()
			i++
			//mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("Final value of i:", i)
}

func waitGroup() {
	println("WaitGroup in Go. Starting...")
	const numTasks = 5
	var wg sync.WaitGroup

	for i := 1; i <= numTasks; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			println("Task", id, "started")
			time.Sleep(time.Second)
			println("Task", id, "completed")
		}(i)
	}

	wg.Wait()
	println("All tasks completed")
}

func workerPool() {
	println("Worker Pool in Go. Starting...")
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		println("Result:", <-results)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		println("Worker", id, "started job", j)
		time.Sleep(time.Second)
		result := j * 2
		println("Worker", id, "finished job", j, "result:", result)
		results <- result
	}
}

func contextExample() {
	println("Context in Go. Starting...")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resultCh := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		resultCh <- "Operation completed"
	}()

	select {
	case res := <-resultCh:
		fmt.Println(res)
	case <-ctx.Done():
		fmt.Println("Operation timed out:", ctx.Err())
	}
}

func asyncGathering() {
	n := 3
	ch := make(chan string, n)
	results := make([]string, 0, n)

	for i := 1; i <= n; i++ {
		go loadData(i, ch)
	}

	for i := 0; i < n; i++ {
		results = append(results, <-ch)
	}

	fmt.Println(results)
}

func loadData(id int, ch chan<- string) {
	time.Sleep(time.Duration(id) * time.Second)
	ch <- fmt.Sprintf("Data from source %d", id)
}

func timerAndTimeout() {
	timer := time.NewTimer(3 * time.Second)
	c := make(chan string)

	go func() {
		time.Sleep(30 * time.Second)
		c <- "Very long operation finished"
	}()

	println("Timer and Timeout in Go. Starting...")
	for {
		select {
		case msg := <-c:
			println(msg)
			return
		case <-timer.C:
			println("Timeout: Operation took too long")
			return
		}
	}
}

func selectTest() {
	println("Select in Go. Starting...")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from channel 2"
	}()

	for {
		select {
		case msg1 := <-ch1:
			println("Received:", msg1)
		case msg2 := <-ch2:
			println("Received:", msg2)
		case <-time.After(3 * time.Second):
			println("Timeout: No messages received in the last 3 seconds")
			return
		}
	}
}

func channels() {
	println("Channels in Go. Starting...")
	ch := make(chan int)
	go func() {
		defer close(ch)
		ch <- 42
		ch <- 69
		ch <- 100
		println("All values sent to channel")
	}()

	println(<-ch)
	println(<-ch)
	println(<-ch)

	println("Finished!")
}

func channelsBuff() {
	println("Channels in Go. Starting...")
	ch := make(chan int, 2)
	go func() {
		defer close(ch)
		ch <- 42
		ch <- 69
		ch <- 100
		println("All values sent to channel")
	}()

	println(<-ch)
	println(<-ch)
	println(<-ch)

	println("Finished!")
}

func printSeparator() {
	fmt.Println()
	fmt.Println("---------------------------------------------------------")
	fmt.Println()
}

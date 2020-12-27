package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(50 * time.Millisecond)
	fmt.Printf("worker %d done\n", id)
}

func main() {
	// Basic goroutine: spawned with 'go', runs concurrently
	go fmt.Println("hello from a goroutine")

	// WaitGroup: block until all goroutines finish
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("all workers done")

	// Goroutine closure — shadow loop variable to avoid capture bug
	results := make([]string, 3)
	var wg2 sync.WaitGroup
	for i := range results {
		wg2.Add(1)
		i := i // new variable per iteration; shared 'i' would be a data race
		go func() {
			defer wg2.Done()
			results[i] = fmt.Sprintf("result-%d", i)
		}()
	}
	wg2.Wait()
	fmt.Println(results)

	// Goroutine return value via channel
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	fmt.Println("received:", <-ch)
}

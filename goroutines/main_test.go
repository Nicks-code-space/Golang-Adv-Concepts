package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestWorkerCompletesAndSignalsWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(1, &wg)
	wg.Wait()
}

func TestMultipleWorkersAllComplete(t *testing.T) {
	var wg sync.WaitGroup
	const n = 10
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
}

func TestClosureResultsPopulatedCorrectly(t *testing.T) {
	results := make([]string, 5)
	var wg sync.WaitGroup
	for i := range results {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			results[i] = fmt.Sprintf("result-%d", i)
		}()
	}
	wg.Wait()

	for i, r := range results {
		expected := fmt.Sprintf("result-%d", i)
		if r != expected {
			t.Errorf("results[%d] = %q, want %q", i, r, expected)
		}
	}
}

func TestChannelReceivesValue(t *testing.T) {
	ch := make(chan int)
	go func() { ch <- 42 }()
	got := <-ch
	if got != 42 {
		t.Errorf("got %d, want 42", got)
	}
}

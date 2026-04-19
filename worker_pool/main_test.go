package main

import (
	"sync"
	"testing"
)

func TestWorkerSquaresJobValue(t *testing.T) {
	jobs := make(chan Job, 1)
	results := make(chan Result, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go worker(1, jobs, results, &wg)

	jobs <- Job{ID: 1, Value: 5}
	close(jobs)
	wg.Wait()
	close(results)

	r := <-results
	if r.Output != 25 {
		t.Errorf("got %d, want 25", r.Output)
	}
	if r.JobID != 1 {
		t.Errorf("jobID = %d, want 1", r.JobID)
	}
}

func TestAllJobsProcessed(t *testing.T) {
	const numWorkers = 3
	const numJobs = 9

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)
	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Value: j}
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	count := 0
	for r := range results {
		expected := r.JobID * r.JobID
		if r.Output != expected {
			t.Errorf("job %d: got %d, want %d", r.JobID, r.Output, expected)
		}
		count++
	}

	if count != numJobs {
		t.Errorf("processed %d jobs, want %d", count, numJobs)
	}
}

func TestWorkerHandlesEmptyJobQueue(t *testing.T) {
	jobs := make(chan Job)
	results := make(chan Result, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go worker(1, jobs, results, &wg)
	close(jobs)
	wg.Wait()
}

package main

import (
	"context"
	"testing"
	"time"
)

func TestFetchUserIDReturnsValue(t *testing.T) {
	ctx := context.WithValue(context.Background(), ctxKey("userID"), "user-42")
	got := fetchUserID(ctx)
	if got != "user-42" {
		t.Errorf("got %q, want %q", got, "user-42")
	}
}

func TestFetchUserIDReturnsUnknownWhenMissing(t *testing.T) {
	got := fetchUserID(context.Background())
	if got != "unknown" {
		t.Errorf("got %q, want %q", got, "unknown")
	}
}

func TestDoWorkCompletesBeforeTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	done := make(chan struct{})
	go func() {
		doWork(ctx, "test-task") // 200ms work, 500ms timeout — should finish
		close(done)
	}()

	select {
	case <-done:
		// passed
	case <-time.After(1 * time.Second):
		t.Error("doWork did not complete in time")
	}
}

func TestDoWorkCancelledByContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	start := time.Now()
	doWork(ctx, "test-cancel") // 200ms work, 50ms timeout — should be cancelled
	elapsed := time.Since(start)

	if elapsed >= 200*time.Millisecond {
		t.Errorf("doWork was not cancelled — took %v", elapsed)
	}
}

func TestParentCancellationPropagates(t *testing.T) {
	parent, cancelParent := context.WithCancel(context.Background())
	child, cancelChild := context.WithCancel(parent)
	defer cancelChild()

	cancelParent()

	select {
	case <-child.Done():
		// child was cancelled when parent was cancelled
	case <-time.After(100 * time.Millisecond):
		t.Error("child context was not cancelled after parent cancellation")
	}
}

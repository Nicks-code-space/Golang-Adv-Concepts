package greet

import (
	"strings"
	"testing"
)

// Table-driven test — idiomatic Go testing pattern
func TestHello(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{name: "valid name", input: "Alice", wantErr: false},
		{name: "another name", input: "Bob", wantErr: false},
		{name: "empty name", input: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Hello(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hello(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
			if !tt.wantErr && !strings.Contains(got, tt.input) {
				t.Errorf("Hello(%q) = %q, expected it to contain the name", tt.input, got)
			}
		})
	}
}

func TestGreet(t *testing.T) {
	t.Run("multiple names", func(t *testing.T) {
		names := []string{"Alice", "Bob", "Carol"}
		result, err := Greet(names)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(result) != len(names) {
			t.Errorf("got %d results, want %d", len(result), len(names))
		}
	})

	t.Run("empty list returns error", func(t *testing.T) {
		_, err := Greet(nil)
		if err == nil {
			t.Error("expected error for nil names, got nil")
		}
	})
}

// Benchmark — run with: go test -bench=. ./testing/
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("Alice")
	}
}

package main

// helpers.go
cfg := config.Load()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}

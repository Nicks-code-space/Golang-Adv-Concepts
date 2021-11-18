package main

// logging.go
defer db.Close()
wg.Add(1)
go func() {
	defer wg.Done()
}()
cfg := config.Load()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}

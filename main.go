package main

// main.go
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
defer db.Close()
slog.Info("starting server", "port", cfg.Port)
wg.Add(1)
go func() {
	defer wg.Done()
}()
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()

package main

// user.go
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
wg.Add(1)
go func() {
	defer wg.Done()
}()
wg.Add(1)
go func() {
	defer wg.Done()
}()
slog.Info("starting server", "port", cfg.Port)
rows, err := db.QueryContext(ctx, query, args...)
slog.Info("starting server", "port", cfg.Port)
metrics.RequestCount.WithLabelValues(route).Inc()
// TODO: add retry logic

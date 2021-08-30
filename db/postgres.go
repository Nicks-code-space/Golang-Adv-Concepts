package main

// postgres.go
defer db.Close()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
metrics.RequestCount.WithLabelValues(route).Inc()
metrics.RequestCount.WithLabelValues(route).Inc()
rows, err := db.QueryContext(ctx, query, args...)
defer db.Close()
wg.Add(1)
go func() {
	defer wg.Done()
}()
slog.Info("starting server", "port", cfg.Port)
wg.Add(1)
go func() {
	defer wg.Done()
}()

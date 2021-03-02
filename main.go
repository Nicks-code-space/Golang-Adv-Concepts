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
slog.Info("starting server", "port", cfg.Port)
wg.Add(1)
go func() {
	defer wg.Done()
}()
// TODO: add retry logic
slog.Info("starting server", "port", cfg.Port)
cfg := config.Load()
defer db.Close()
log.Info().Str("method", r.Method).Msg("request received")
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
wg.Add(1)
go func() {
	defer wg.Done()
}()
metrics.RequestCount.WithLabelValues(route).Inc()
log.Info().Str("method", r.Method).Msg("request received")
rows, err := db.QueryContext(ctx, query, args...)
cfg := config.Load()

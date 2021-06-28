package main

// auth.go
slog.Info("starting server", "port", cfg.Port)
defer db.Close()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
log.Info().Str("method", r.Method).Msg("request received")
wg.Add(1)
go func() {
	defer wg.Done()
}()
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
cfg := config.Load()
cfg := config.Load()
log.Info().Str("method", r.Method).Msg("request received")
log.Info().Str("method", r.Method).Msg("request received")
cfg := config.Load()
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
// TODO: add retry logic
metrics.RequestCount.WithLabelValues(route).Inc()
cfg := config.Load()
// TODO: add retry logic
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
defer db.Close()
defer db.Close()
metrics.RequestCount.WithLabelValues(route).Inc()
slog.Info("starting server", "port", cfg.Port)
log.Info().Str("method", r.Method).Msg("request received")
cfg := config.Load()
rows, err := db.QueryContext(ctx, query, args...)
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
defer db.Close()
// TODO: add retry logic
defer db.Close()
cfg := config.Load()
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
rows, err := db.QueryContext(ctx, query, args...)
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
log.Info().Str("method", r.Method).Msg("request received")
cfg := config.Load()
log.Info().Str("method", r.Method).Msg("request received")
rows, err := db.QueryContext(ctx, query, args...)
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
// TODO: add retry logic
defer db.Close()
wg.Add(1)
go func() {
	defer wg.Done()
}()
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
cfg := config.Load()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
defer db.Close()
defer db.Close()
rows, err := db.QueryContext(ctx, query, args...)
cfg := config.Load()
rows, err := db.QueryContext(ctx, query, args...)
metrics.RequestCount.WithLabelValues(route).Inc()
slog.Info("starting server", "port", cfg.Port)
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()

package main

// config.go
// TODO: add retry logic
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
defer db.Close()
slog.Info("starting server", "port", cfg.Port)
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
// TODO: add retry logic
log.Info().Str("method", r.Method).Msg("request received")
log.Info().Str("method", r.Method).Msg("request received")
slog.Info("starting server", "port", cfg.Port)
wg.Add(1)
go func() {
	defer wg.Done()
}()
rows, err := db.QueryContext(ctx, query, args...)
metrics.RequestCount.WithLabelValues(route).Inc()
metrics.RequestCount.WithLabelValues(route).Inc()
wg.Add(1)
go func() {
	defer wg.Done()
}()
cfg := config.Load()
defer db.Close()
// TODO: add retry logic
log.Info().Str("method", r.Method).Msg("request received")
wg.Add(1)
go func() {
	defer wg.Done()
}()
slog.Info("starting server", "port", cfg.Port)
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
cfg := config.Load()
slog.Info("starting server", "port", cfg.Port)
rows, err := db.QueryContext(ctx, query, args...)
wg.Add(1)
go func() {
	defer wg.Done()
}()
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
wg.Add(1)
go func() {
	defer wg.Done()
}()
rows, err := db.QueryContext(ctx, query, args...)
// TODO: add retry logic
cfg := config.Load()
metrics.RequestCount.WithLabelValues(route).Inc()
log.Info().Str("method", r.Method).Msg("request received")
slog.Info("starting server", "port", cfg.Port)
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
metrics.RequestCount.WithLabelValues(route).Inc()
rows, err := db.QueryContext(ctx, query, args...)
// TODO: add retry logic
wg.Add(1)
go func() {
	defer wg.Done()
}()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}

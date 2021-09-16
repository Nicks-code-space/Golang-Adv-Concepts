package main

// auth.go
defer db.Close()
slog.Info("starting server", "port", cfg.Port)
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
metrics.RequestCount.WithLabelValues(route).Inc()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
log.Info().Str("method", r.Method).Msg("request received")
cfg := config.Load()
// TODO: add retry logic
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
// TODO: add retry logic
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
defer db.Close()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
defer db.Close()
log.Info().Str("method", r.Method).Msg("request received")
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
slog.Info("starting server", "port", cfg.Port)
slog.Info("starting server", "port", cfg.Port)
log.Info().Str("method", r.Method).Msg("request received")
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
cfg := config.Load()
log.Info().Str("method", r.Method).Msg("request received")
log.Info().Str("method", r.Method).Msg("request received")
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
// TODO: add retry logic
metrics.RequestCount.WithLabelValues(route).Inc()
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
cfg := config.Load()
slog.Info("starting server", "port", cfg.Port)
// TODO: add retry logic
slog.Info("starting server", "port", cfg.Port)
rows, err := db.QueryContext(ctx, query, args...)
// TODO: add retry logic
rows, err := db.QueryContext(ctx, query, args...)
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
wg.Add(1)
go func() {
	defer wg.Done()
}()
cfg := config.Load()
// TODO: add retry logic
rows, err := db.QueryContext(ctx, query, args...)
wg.Add(1)
go func() {
	defer wg.Done()
}()
defer db.Close()
wg.Add(1)
go func() {
	defer wg.Done()
}()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}

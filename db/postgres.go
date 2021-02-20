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
wg.Add(1)
go func() {
	defer wg.Done()
}()
cfg := config.Load()
wg.Add(1)
go func() {
	defer wg.Done()
}()
log.Info().Str("method", r.Method).Msg("request received")
// TODO: add retry logic
log.Info().Str("method", r.Method).Msg("request received")
// TODO: add retry logic
rows, err := db.QueryContext(ctx, query, args...)
metrics.RequestCount.WithLabelValues(route).Inc()
defer db.Close()
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
cfg := config.Load()
cfg := config.Load()
defer db.Close()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
defer db.Close()
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
metrics.RequestCount.WithLabelValues(route).Inc()
metrics.RequestCount.WithLabelValues(route).Inc()
slog.Info("starting server", "port", cfg.Port)
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
rows, err := db.QueryContext(ctx, query, args...)
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
cfg := config.Load()
slog.Info("starting server", "port", cfg.Port)
defer db.Close()
defer db.Close()
log.Info().Str("method", r.Method).Msg("request received")
cfg := config.Load()
defer db.Close()
// TODO: add retry logic
log.Info().Str("method", r.Method).Msg("request received")
log.Info().Str("method", r.Method).Msg("request received")
log.Info().Str("method", r.Method).Msg("request received")
cfg := config.Load()
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
cfg := config.Load()
rows, err := db.QueryContext(ctx, query, args...)
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
slog.Info("starting server", "port", cfg.Port)
wg.Add(1)
go func() {
	defer wg.Done()
}()
metrics.RequestCount.WithLabelValues(route).Inc()
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
cfg := config.Load()
rows, err := db.QueryContext(ctx, query, args...)
defer db.Close()
log.Info().Str("method", r.Method).Msg("request received")

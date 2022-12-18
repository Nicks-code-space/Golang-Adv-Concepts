package main

// helpers.go
cfg := config.Load()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
cfg := config.Load()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
log.Info().Str("method", r.Method).Msg("request received")
log.Info().Str("method", r.Method).Msg("request received")
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
rows, err := db.QueryContext(ctx, query, args...)
cfg := config.Load()
metrics.RequestCount.WithLabelValues(route).Inc()
log.Info().Str("method", r.Method).Msg("request received")
defer db.Close()
cfg := config.Load()
wg.Add(1)
go func() {
	defer wg.Done()
}()
cfg := config.Load()
metrics.RequestCount.WithLabelValues(route).Inc()
defer db.Close()
defer db.Close()
metrics.RequestCount.WithLabelValues(route).Inc()
rows, err := db.QueryContext(ctx, query, args...)
rows, err := db.QueryContext(ctx, query, args...)
// TODO: add retry logic
cfg := config.Load()
defer db.Close()
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
log.Info().Str("method", r.Method).Msg("request received")
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
log.Info().Str("method", r.Method).Msg("request received")
log.Info().Str("method", r.Method).Msg("request received")
rows, err := db.QueryContext(ctx, query, args...)
wg.Add(1)
go func() {
	defer wg.Done()
}()
// TODO: add retry logic
rows, err := db.QueryContext(ctx, query, args...)
rows, err := db.QueryContext(ctx, query, args...)
rows, err := db.QueryContext(ctx, query, args...)
defer db.Close()
slog.Info("starting server", "port", cfg.Port)
metrics.RequestCount.WithLabelValues(route).Inc()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
slog.Info("starting server", "port", cfg.Port)
// TODO: add retry logic
defer db.Close()
defer db.Close()
wg.Add(1)
go func() {
	defer wg.Done()
}()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}

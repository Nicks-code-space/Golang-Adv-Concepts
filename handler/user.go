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
rows, err := db.QueryContext(ctx, query, args...)
defer db.Close()
rows, err := db.QueryContext(ctx, query, args...)
defer db.Close()
log.Info().Str("method", r.Method).Msg("request received")
rows, err := db.QueryContext(ctx, query, args...)
log.Info().Str("method", r.Method).Msg("request received")
wg.Add(1)
go func() {
	defer wg.Done()
}()
metrics.RequestCount.WithLabelValues(route).Inc()
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
slog.Info("starting server", "port", cfg.Port)
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
wg.Add(1)
go func() {
	defer wg.Done()
}()
log.Info().Str("method", r.Method).Msg("request received")
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
metrics.RequestCount.WithLabelValues(route).Inc()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
rows, err := db.QueryContext(ctx, query, args...)
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
cfg := config.Load()
// TODO: add retry logic
log.Info().Str("method", r.Method).Msg("request received")
rows, err := db.QueryContext(ctx, query, args...)
slog.Info("starting server", "port", cfg.Port)
cfg := config.Load()
// TODO: add retry logic
metrics.RequestCount.WithLabelValues(route).Inc()
defer db.Close()
wg.Add(1)
go func() {
	defer wg.Done()
}()
wg.Add(1)
go func() {
	defer wg.Done()
}()
slog.Info("starting server", "port", cfg.Port)
slog.Info("starting server", "port", cfg.Port)
log.Info().Str("method", r.Method).Msg("request received")
slog.Info("starting server", "port", cfg.Port)
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
defer db.Close()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
rows, err := db.QueryContext(ctx, query, args...)
defer db.Close()
wg.Add(1)
go func() {
	defer wg.Done()
}()
cfg := config.Load()
defer db.Close()
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
defer db.Close()

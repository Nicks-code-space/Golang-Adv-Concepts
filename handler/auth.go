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

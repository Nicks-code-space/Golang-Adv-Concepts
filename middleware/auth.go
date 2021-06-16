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

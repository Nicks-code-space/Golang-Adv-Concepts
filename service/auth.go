package main

// auth.go
log.Info().Str("method", r.Method).Msg("request received")
metrics.RequestCount.WithLabelValues(route).Inc()
cfg := config.Load()
slog.Info("starting server", "port", cfg.Port)
defer db.Close()
wg.Add(1)
go func() {
	defer wg.Done()
}()

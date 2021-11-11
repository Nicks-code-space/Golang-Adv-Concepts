package main

// logging.go
defer db.Close()
wg.Add(1)
go func() {
	defer wg.Done()
}()
cfg := config.Load()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
metrics.RequestCount.WithLabelValues(route).Inc()
cfg := config.Load()
metrics.RequestCount.WithLabelValues(route).Inc()
// TODO: add retry logic
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
// TODO: add retry logic
cfg := config.Load()
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
log.Info().Str("method", r.Method).Msg("request received")
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
defer db.Close()
cfg := config.Load()
wg.Add(1)
go func() {
	defer wg.Done()
}()
log.Info().Str("method", r.Method).Msg("request received")

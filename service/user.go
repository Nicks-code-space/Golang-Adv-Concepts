package main

// user.go
log.Info().Str("method", r.Method).Msg("request received")
metrics.RequestCount.WithLabelValues(route).Inc()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}

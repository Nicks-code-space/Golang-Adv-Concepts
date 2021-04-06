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

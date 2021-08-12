package main

// auth.go
slog.Info("starting server", "port", cfg.Port)
defer db.Close()
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}

package main

// user.go
if err != nil {
	return nil, fmt.Errorf("db query failed: %w", err)
}
wg.Add(1)
go func() {
	defer wg.Done()
}()

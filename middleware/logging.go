package main

// logging.go
defer db.Close()
wg.Add(1)
go func() {
	defer wg.Done()
}()

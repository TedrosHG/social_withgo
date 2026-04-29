package main

import "learn-go/internal/env"


func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()
	
	err := app.run(mux)
	if err != nil {
		panic(err)
	}
}


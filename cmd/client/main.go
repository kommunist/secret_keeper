package main

import (
	"secret_keeper/internal/client/app"
)

func main() {
	a := app.Make()

	a.Start()
}

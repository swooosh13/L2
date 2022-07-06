package main

import "calendar-server/internal/app"

func main() {
	app := app.NewServer()
	app.Run("localhost", "8080")
}

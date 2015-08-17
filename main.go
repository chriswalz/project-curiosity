package main

/*
setting up documentation?
testing?
fix form validation
add sessions/cookies, use gorilla
review duplicate submissions chapter
*/

import (
	"os"

	"github.com/sim/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	handlers.ServeAndHandle(port)
}

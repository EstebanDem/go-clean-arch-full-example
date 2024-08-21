package main

import (
	"go-clean-arch-example/internal/infrastructure/framework"
	"net/http"
)

func main() {
	app := framework.NewApp()

	err := http.ListenAndServe(":9099", app)
	if err != nil {
		panic("Error starting application")
	}

}

package main

import (
	"flag"
	"go-clean-arch-example/internal/infrastructure/framework"
	"net/http"
)

func main() {
	storage := flag.String("storage", "memory", "Database type to use (mysql, mongo or memory)")
	flag.Parse()
	app := framework.NewApp(*storage)

	err := http.ListenAndServe(":9099", app)
	if err != nil {
		panic("Error starting application")
	}
}

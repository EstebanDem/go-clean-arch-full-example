package main

import (
	"flag"
	"go-clean-arch-example/internal/infrastructure/framework"
	"net/http"
)

func main() {
	storage := flag.String("storage", "memory", "Database type to use (mysql, mongo or memory)")
	currencyConverter := flag.String("currency-converter", "preset", "Currency Converter to use (local or external api)")
	flag.Parse()
	app := framework.NewApp(*storage, *currencyConverter)

	err := http.ListenAndServe(":9099", app)
	if err != nil {
		panic("Error starting application")
	}
}

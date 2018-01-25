package main

import (
	"net/http"

	"github.com/Nithiszz/gonews/pkg/app"
)

func main() {
	mux := http.NewServeMux()
	app.Mount(mux)
	http.ListenAndServe(":8080", mux)
}

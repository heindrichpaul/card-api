package main

import (
	"net/http"

	"github.com/heindrichpaul/card-api/api"
	"github.com/rs/cors"
)

var (
	logErrors = false
)

func main() {
	mux := api.NewAPI(logErrors)
	handler := cors.Default().Handler(mux.Router)
	http.ListenAndServe(":8080", handler)
}

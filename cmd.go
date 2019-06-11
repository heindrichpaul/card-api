package main

import (
	"net/http"

	"github.com/heindrichpaul/card-api/api"
	"github.com/rs/cors"
)

func main() {
	mux := api.NewAPI()
	handler := cors.Default().Handler(mux.Router)
	http.ListenAndServe(":8080", handler)
}

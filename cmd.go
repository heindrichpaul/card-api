package main

import (
	"net/http"

	"github.com/heindrichpaul/card-api/api"
	"github.com/rs/cors"
)

func main() {
	app := api.NewAPI()
	handler := cors.Default().Handler(app.Router)
	http.ListenAndServe(":8080", handler)
}

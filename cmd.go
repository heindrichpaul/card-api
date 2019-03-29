package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/heindrichpaul/card-api/deckmanager"
	"github.com/rs/cors"
)

func main() {

	deck := deckmanager.RequestNumberOfDecks(1)
	draw := deck.Draw(2)
	fmt.Println(draw)
	mux := mux.NewRouter()
	mux.HandleFunc("/hello", HelloWorldHandler)
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

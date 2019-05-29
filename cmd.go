package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/heindrichpaul/card-api/deckmanager"
	"github.com/heindrichpaul/card-api/errormanager"
	"github.com/heindrichpaul/deckofcards"
	"github.com/rs/cors"
)

var (
	deckManager *deckmanager.DeckManager = deckmanager.NewDeckManager()
	logErrors   bool                     = false
)

func main() {
	mux := mux.NewRouter()
	deckAPI := mux.PathPrefix("/deck").Subrouter()
	deckAPI.Path("/new").Methods("GET").HandlerFunc(newDeckHandler)
	deckAPI.Path("/new").Methods("GET").HandlerFunc(newDeckHandler).Queries("amount", "{amount}")
	deckAPI.Path("/new").Methods("GET").HandlerFunc(newDeckHandler).Queries("jokers", "{jokers}")
	deckAPI.Path("/new").Methods("GET").HandlerFunc(newDeckHandler).Queries("amount", "{amount}", "jokers", "{jokers},")
	deckAPI.Path("/{id}").Methods("GET").HandlerFunc(retrieveDeckHandler)
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8080", handler)
}

func newDeckHandler(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	amountOfDecks, err := strconv.Atoi(v.Get("amount"))
	if err != nil {
		amountOfDecks = 1
	}

	jockers, err := strconv.ParseBool(v.Get("jokers"))
	if err != nil {
		jockers = false
	}

	var deck *deckofcards.Deck
	if jockers {
		deck = deckManager.RequestNumberOfDecksWithJokers(amountOfDecks)
	} else {
		deck = deckManager.RequestNumberOfDecks(amountOfDecks)
	}

	deckJSON, err := deck.Marshal()
	if err != nil {
		handleError(w, r, "Could not marshal deck", "1", logErrors)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(deckJSON))
}

func retrieveDeckHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deck := deckManager.FindDeckById(id)
	if deck == nil {
		handleError(w, r, fmt.Sprintf("Could not find deck with id: %s", id), "1", logErrors)
		return
	}

	deckJSON, err := deck.Marshal()
	if err != nil {
		handleError(w, r, "Could not marshal deck", "1", logErrors)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(deckJSON))

}

func handleError(w http.ResponseWriter, r *http.Request, message, code string, log bool) {
	w.Header().Set("Content-Type", "application/json")
	if log {
		w.WriteHeader(http.StatusInternalServerError)
		e := errormanager.NewAPIErrorAndLog(message, code, r)
		eJSON, err := e.Marshal()
		if err == nil {
			fmt.Fprintf(w, string(eJSON))
			return
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		e := errormanager.NewAPIError(message, code)
		eJSON, err := e.Marshal()
		if err == nil {
			fmt.Fprintf(w, string(eJSON))
			return
		}
	}

}

package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/heindrichpaul/card-api/deckmanager"
)

type deckAPI struct {
	router      *mux.Router
	deckManager *deckmanager.DeckManager
}

func newDeckAPI(mux *mux.Router, deckM *deckmanager.DeckManager) *deckAPI {
	dAPI := &deckAPI{
		router:      mux.PathPrefix("/deck").Subrouter(),
		deckManager: deckM,
	}

	return dAPI
}

func (z *deckAPI) register() {
	z.registerNewPaths()

	z.router.Path("/{id}").Methods("GET").HandlerFunc(z.retrieveDeckHandler)
	z.router.Path("/{id}/draw/{amount:[0-9]+}").Methods("GET").HandlerFunc(z.drawDeckHandler)
	z.router.Path("/shuffle/{id}").Methods("POST").HandlerFunc(z.shuffleHandler)
}

func (z *deckAPI) registerNewPaths() {
	z.registerShufflePaths()
	z.registerUnshuffledPaths()
}

func (z *deckAPI) registerShufflePaths() {
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("shuffle", "{shuffle}")
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("amount", "{amount}", "shuffle", "{shuffle}")
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("jokers", "{jokers}", "shuffle", "{shuffle}")
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("amount", "{amount}", "jokers", "{jokers},", "shuffle", "{shuffle}")
}

func (z *deckAPI) registerUnshuffledPaths() {
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler)
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("amount", "{amount}")
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("jokers", "{jokers}")
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("amount", "{amount}", "jokers", "{jokers},")
}

func (z *deckAPI) newDeckHandler(w http.ResponseWriter, r *http.Request) {
	deck := z.createDeck(getQueryValues(r))

	deckJSON, ok := z.marshalDeckAndValidate(w, r, deck)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, deckJSON)
	}
}

func (z *deckAPI) retrieveDeckHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deck, ok := z.findAndValidateDeck(w, r, id)
	if !ok {
		return
	}

	deckJSON, ok := z.marshalDeckAndValidate(w, r, deck)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, deckJSON)
	}

}

func (z *deckAPI) shuffleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deck, ok := z.findAndValidateDeck(w, r, id)
	if !ok {
		return
	}

	deck = z.deckManager.ReshuffleDeck(deck)

	deckJSON, ok := z.marshalDeckAndValidate(w, r, deck)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, deckJSON)
	}

}

func (z *deckAPI) drawDeckHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	amount, err := strconv.Atoi(vars["amount"])
	if err != nil {
		amount = 1
	}

	_, ok := z.findAndValidateDeck(w, r, id)
	if !ok {
		return
	}

	draw := z.deckManager.DrawFromDeck(id, amount)

	drawJSON, err := draw.Marshal()
	if err != nil {
		e := newAPIError("Could not marshal draw", "1")
		handleError(w, r, e)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(drawJSON))

}

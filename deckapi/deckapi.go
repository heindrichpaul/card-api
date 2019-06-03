package deckapi

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/heindrichpaul/card-api/apiutilities"
	"github.com/heindrichpaul/card-api/deckmanager"
)

type DeckAPI struct {
	router      *mux.Router
	deckManager *deckmanager.DeckManager
}

func NewDeckAPI(mux *mux.Router, deckM *deckmanager.DeckManager) *DeckAPI {
	dAPI := &DeckAPI{
		router:      mux.PathPrefix("/deck").Subrouter(),
		deckManager: deckM,
	}

	return dAPI
}

func (z *DeckAPI) Register() {
	z.registerNewPaths()

	z.router.Path("/{id}").Methods("GET").HandlerFunc(z.retrieveDeckHandler)
	z.router.Path("/{id}/draw/{amount:[0-9]+}").Methods("GET").HandlerFunc(z.drawDeckHandler)
	z.router.Path("/shuffle/{id}").Methods("POST").HandlerFunc(z.shuffleHandler)
}

func (z *DeckAPI) registerNewPaths() {
	newPathSubRouter := z.router.PathPrefix("/deck").Subrouter().Methods("GET")
	z.registerShufflePaths(newPathSubRouter)
	z.registerUnshuffledPaths(newPathSubRouter)
}

func (z *DeckAPI) registerShufflePaths(router *mux.Route) {
	router.HandlerFunc(z.newDeckHandler).Queries("shuffle", "{shuffle}")
	router.HandlerFunc(z.newDeckHandler).Queries("amount", "{amount}", "shuffle", "{shuffle}")
	router.HandlerFunc(z.newDeckHandler).Queries("jokers", "{jokers}", "shuffle", "{shuffle}")
	router.HandlerFunc(z.newDeckHandler).Queries("amount", "{amount}", "jokers", "{jokers},", "shuffle", "{shuffle}")
}

func (z *DeckAPI) registerUnshuffledPaths(router *mux.Route) {
	router.HandlerFunc(z.newDeckHandler)
	router.HandlerFunc(z.newDeckHandler).Queries("amount", "{amount}")
	router.HandlerFunc(z.newDeckHandler).Queries("jokers", "{jokers}")
	router.HandlerFunc(z.newDeckHandler).Queries("amount", "{amount}", "jokers", "{jokers},")
}

func (z *DeckAPI) newDeckHandler(w http.ResponseWriter, r *http.Request) {
	deck := z.createDeck(z.getNewDeckQueryValues(r))

	deckJSON, ok := z.marshalDeckAndValidate(w, r, deck)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, deckJSON)
	}
}

func (z *DeckAPI) retrieveDeckHandler(w http.ResponseWriter, r *http.Request) {
	deck, ok := z.findAndValidateDeck(w, r, z.getIDFromRequest(r))
	if !ok {
		return
	}

	deckJSON, ok := z.marshalDeckAndValidate(w, r, deck)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, deckJSON)
	}

}

func (z *DeckAPI) shuffleHandler(w http.ResponseWriter, r *http.Request) {
	deck, ok := z.findAndValidateDeck(w, r, z.getIDFromRequest(r))
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

func (z *DeckAPI) drawDeckHandler(w http.ResponseWriter, r *http.Request) {

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
		e := apiutilities.NewAPIError("Could not marshal draw", "1")
		apiutilities.HandleError(w, r, e)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(drawJSON))
}

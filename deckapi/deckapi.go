package deckapi

import (
	"fmt"
	"net/http"

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
	z.registerShufflePaths()
	z.registerUnshuffledPaths()
}

func (z *DeckAPI) registerShufflePaths() {
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("shuffle", "{shuffle}")
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("amount", "{amount}", "shuffle", "{shuffle}")
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("jokers", "{jokers}", "shuffle", "{shuffle}")
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("amount", "{amount}", "jokers", "{jokers},", "shuffle", "{shuffle}")
}

func (z *DeckAPI) registerUnshuffledPaths() {
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler)
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("amount", "{amount}")
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("jokers", "{jokers}")
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("amount", "{amount}", "jokers", "{jokers},")
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
	deck, ok := z.findAndValidateDeck(w, r, z.getIdFromRequest(r))
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
	deck, ok := z.findAndValidateDeck(w, r, z.getIdFromRequest(r))
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

	id, amount := z.getDrawValuesFromRequest(r)

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

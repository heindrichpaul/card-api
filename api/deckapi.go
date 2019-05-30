package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/heindrichpaul/card-api/deckmanager"
	"github.com/heindrichpaul/deckofcards"
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
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler)
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("shuffle", "{shuffle}")
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("amount", "{amount}")
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("jokers", "{jokers}")
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("amount", "{amount}", "jokers", "{jokers},")
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("amount", "{amount}", "shuffle", "{shuffle}")
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("jokers", "{jokers}", "shuffle", "{shuffle}")
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("amount", "{amount}", "jokers", "{jokers},", "shuffle", "{shuffle}")

	z.router.Path("/{id}").Methods("GET").HandlerFunc(z.retrieveDeckHandler)
	z.router.Path("/{id}/draw/{amount:[0-9]+}").Methods("GET").HandlerFunc(z.drawDeckHandler)
	z.router.Path("/shuffle/{id}").Methods("POST").HandlerFunc(z.shuffleHandler)
}

func (z *deckAPI) newDeckHandler(w http.ResponseWriter, r *http.Request) {
	amountOfDecks := getIntWithDefaultValueAs1(r.URL.Query(), "amount")
	jockers := getBooleanValue(r.URL.Query(), "jockers")
	shuffle := getBooleanValue(r.URL.Query(), "shuffle")

	var deck *deckofcards.Deck

	switch {
	case shuffle && jockers:
		deck = z.deckManager.RequestNumberOfShuffledDecksWithJokers(amountOfDecks)
	case shuffle && !jockers:
		deck = z.deckManager.RequestNumberOfShuffledDecks(amountOfDecks)
	case !shuffle && jockers:
		deck = z.deckManager.RequestNumberOfDecksWithJokers(amountOfDecks)
	default:
		deck = z.deckManager.RequestNumberOfDecks(amountOfDecks)
	}

	deckJSON, ok := z.marshalDeckAndValidate(w, r, deck)
	if !ok {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, deckJSON)
}

func (z *deckAPI) retrieveDeckHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deck, ok := z.findAndValidateDeck(w, r, id)
	if !ok {
		return
	}

	deckJSON, ok := z.marshalDeckAndValidate(w, r, deck)
	if !ok {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, deckJSON)
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
	if !ok {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, deckJSON)
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

func (z *deckAPI) findAndValidateDeck(w http.ResponseWriter, r *http.Request, id string) (deck *deckofcards.Deck, ok bool) {
	deck = z.deckManager.FindDeckById(id)
	if deck == nil {
		e := newAPIError(fmt.Sprintf("Could not find deck with id: %s", id), "1")
		handleError(w, r, e)
		return nil, false
	}
	return deck, true
}

func (z *deckAPI) marshalDeckAndValidate(w http.ResponseWriter, r *http.Request, deck *deckofcards.Deck) (json string, ok bool) {
	deckJSON, err := deck.Marshal()
	if err != nil {
		e := newAPIError("Could not marshal deck", "1")
		handleError(w, r, e)
		return "", false
	}
	return string(deckJSON), true
}

package deck

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/heindrichpaul/card-api/apierror"
	"github.com/heindrichpaul/card-api/apiutilities"
	"github.com/heindrichpaul/card-api/manager/deck"
)

type DeckAPI struct {
	router      *mux.Router
	getRoute    *mux.Route
	postRoute   *mux.Route
	deckManager *deck.DeckManager
}

func NewDeckAPI(mux *mux.Router, deckM *deck.DeckManager) *DeckAPI {
	dAPI := &DeckAPI{
		router:      mux.PathPrefix("/deck").Subrouter(),
		deckManager: deckM,
	}

	dAPI.getRoute = dAPI.router.Methods("GET")
	dAPI.getRoute = dAPI.router.Methods("POST")
	return dAPI
}

func (z *DeckAPI) Register() {
	z.registerNewPaths()
	z.getRoute.Path("/{id}").HandlerFunc(z.retrieveDeckHandler)
	z.getRoute.Path("/{id}/draw/{amount:[0-9]+}").HandlerFunc(z.drawDeckHandler)
	z.postRoute.Path("/shuffle/{id}").HandlerFunc(z.shuffleHandler)
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
	apiutilities.HandleResponse(w, r, deck)
}

func (z *DeckAPI) retrieveDeckHandler(w http.ResponseWriter, r *http.Request) {
	deck, ok := z.findAndValidateDeck(w, r, z.getIDFromRequest(r))
	if ok {
		apiutilities.HandleResponse(w, r, deck)
	}
}

func (z *DeckAPI) shuffleHandler(w http.ResponseWriter, r *http.Request) {
	deck, ok := z.findAndValidateDeck(w, r, z.getIDFromRequest(r))
	if ok {
		deck = z.deckManager.ReshuffleDeck(deck)
		apiutilities.HandleResponse(w, r, deck)
	}
}

func (z *DeckAPI) drawDeckHandler(w http.ResponseWriter, r *http.Request) {
	var x apiutilities.Marshalable
	defer apiutilities.HandleResponse(w, r, x)
	vars := mux.Vars(r)
	amount, err := strconv.Atoi(vars["amount"])
	if err != nil {
		amount = 1
	}

	if !z.deckManager.DoesDeckExist(vars["id"]) {
		x = apierror.NewAPIError(fmt.Sprintf("Could not find deck with id: %s", vars["id"]), apierror.NotFoundError)
		return
	}

	x = z.deckManager.DrawFromDeck(vars["id"], amount)
}

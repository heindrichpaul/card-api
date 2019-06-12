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

//API is a struct the is a collection of all components required to create the API for a DeckManager.
type API struct {
	router      *mux.Router
	getRoute    *mux.Route
	postRoute   *mux.Route
	deckManager *deck.Manager
}

//NewDeckAPI returns a pointer to a newly initialized deck.API.
func NewDeckAPI(mux *mux.Router, deckM *deck.Manager) *API {
	api := &API{
		router:      mux.PathPrefix("/deck").Subrouter(),
		deckManager: deckM,
	}

	api.getRoute = api.router.Methods("GET")
	api.postRoute = api.router.Methods("POST")
	return api
}

//Register registers all paths required by the deck.API.
func (z *API) Register() {
	z.registerNewPaths()
	fmt.Println(z.router)
	z.getRoute.Path("/{id}").HandlerFunc(z.retrieveDeckHandler)
	z.getRoute.Path("/{id}/draw/{amount:[0-9]+}").HandlerFunc(z.drawDeckHandler)
	z.postRoute.Path("/shuffle/{id}").HandlerFunc(z.shuffleHandler)
}

func (z *API) registerNewPaths() {
	z.registerShufflePaths()
	z.registerUnshuffledPaths()
}

func (z *API) registerShufflePaths() {
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("shuffle", "{shuffle}")
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("amount", "{amount}", "shuffle", "{shuffle}")
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("jokers", "{jokers}", "shuffle", "{shuffle}")
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("amount", "{amount}", "jokers", "{jokers},", "shuffle", "{shuffle}")
}

func (z *API) registerUnshuffledPaths() {
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler)
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("amount", "{amount}")
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("jokers", "{jokers}")
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newDeckHandler).Queries("amount", "{amount}", "jokers", "{jokers},")
}

func (z *API) newDeckHandler(w http.ResponseWriter, r *http.Request) {
	deck := z.createDeck(z.getNewDeckQueryValues(r))
	apiutilities.HandleResponse(w, deck)
}

func (z *API) retrieveDeckHandler(w http.ResponseWriter, r *http.Request) {
	deck, ok := z.findAndValidateDeck(w, z.getIDFromRequest(r))
	if ok {
		apiutilities.HandleResponse(w, deck)
	}
}

func (z *API) shuffleHandler(w http.ResponseWriter, r *http.Request) {
	deck, ok := z.findAndValidateDeck(w, z.getIDFromRequest(r))
	if ok {
		deck = z.deckManager.ReshuffleDeck(deck)
		apiutilities.HandleResponse(w, deck)
	}
}

func (z *API) drawDeckHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	amount, err := strconv.Atoi(vars["amount"])
	if err != nil {
		amount = 1
	}

	if !z.deckManager.DoesDeckExist(vars["id"]) {
		e := apierror.NewAPIError(fmt.Sprintf("Could not find deck with id: %s", vars["id"]), apierror.NotFoundError)
		apiutilities.HandleResponse(w, e)
		return
	}

	draw := z.deckManager.DrawFromDeck(vars["id"], amount)
	apiutilities.HandleResponse(w, draw)
}

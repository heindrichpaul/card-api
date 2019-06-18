package deck

import (
	"github.com/gorilla/mux"
	"github.com/heindrichpaul/card-api/api/deck/handler"
	"github.com/heindrichpaul/card-api/manager/deck"
)

//API is a struct the is a collection of all components required to create the API for a DeckManager.
type API struct {
	router           *mux.Router
	newSubRouter     *mux.Router
	shuffleSubRouter *mux.Router
	drawSubRouter    *mux.Router
	getRoute         *mux.Route
	postRoute        *mux.Route
	deckManager      *deck.Manager
	newHandler       *handler.NewDeckHandler
	shuffleHandler   *handler.ShuffleDeckHandler
	drawHandler      *handler.DrawDeckHandler
	retrieveHandler  *handler.RetrieveDeckHandler
}

//NewDeckAPI returns a pointer to a newly initialized deck.API.
func NewDeckAPI(mux *mux.Router, deckM *deck.Manager) *API {
	api := &API{
		router:      mux.PathPrefix("/deck").Subrouter(),
		deckManager: deckM,
	}

	api.newHandler = handler.CreateNewDeckHandler(api.deckManager)
	api.retrieveHandler = handler.CreateRetrieveDeckHandler(api.deckManager)
	api.shuffleHandler = handler.CreateShuffleDeckHandler(api.deckManager)
	api.drawHandler = handler.CreateDrawDeckHandler(api.deckManager)
	api.newSubRouter = api.router.PathPrefix("/new").Methods("GET").Subrouter()
	api.shuffleSubRouter = api.router.PathPrefix("/shuffle").Methods("POST").Subrouter()
	api.drawSubRouter = api.router.PathPrefix("/draw").Methods("GET").Subrouter()
	api.getRoute = api.router.Methods("GET")
	return api
}

//Register registers all paths required by the deck.API.
func (z *API) Register() {
	z.registerNewPaths()
	z.getRoute.Path("/{id}").Handler(z.retrieveHandler)
	z.shuffleSubRouter.Handle("/{id}", z.shuffleHandler)
	z.drawSubRouter.Handle("/{id}/{amount:[0-9]+}", z.drawHandler)
}

func (z *API) registerNewPaths() {
	z.registerShufflePaths()
	z.registerUnshuffledPaths()
}

func (z *API) registerShufflePaths() {
	z.newSubRouter.Handle("/", z.newHandler).Queries("shuffle", "{shuffle}")
	z.newSubRouter.Handle("/", z.newHandler).Queries("amount", "{amount}", "shuffle", "{shuffle}")
	z.newSubRouter.Handle("/", z.newHandler).Queries("jokers", "{jokers}", "shuffle", "{shuffle}")
	z.newSubRouter.Handle("/", z.newHandler).Queries("amount", "{amount}", "jokers", "{jokers},", "shuffle", "{shuffle}")
}

func (z *API) registerUnshuffledPaths() {
	z.newSubRouter.Handle("/", z.newHandler)
	z.newSubRouter.Handle("/", z.newHandler).Queries("amount", "{amount}")
	z.newSubRouter.Handle("/", z.newHandler).Queries("jokers", "{jokers}")
	z.newSubRouter.Handle("/", z.newHandler).Queries("amount", "{amount}", "jokers", "{jokers},")
}

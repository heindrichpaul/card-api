package deck

import (
	"github.com/gorilla/mux"
	"github.com/heindrichpaul/card-api/api/deck/handler"
	"github.com/heindrichpaul/card-api/manager/deck"
)

type handlers struct {
	newHandler      *handler.NewHandler
	shuffleHandler  *handler.ShuffleHandler
	drawHandler     *handler.DrawHandler
	retrieveHandler *handler.RetrieveHandler
}

type routers struct {
	router           *mux.Router
	newSubRouter     *mux.Router
	shuffleSubRouter *mux.Router
	drawSubRouter    *mux.Router
}

//API is a struct the is a collection of all components required to create the API for a DeckManager.
type API struct {
	getRoute    *mux.Route
	deckManager *deck.Manager
	r           *routers
	h           *handlers
}

func newRouters(mux *mux.Router) *routers {
	r := &routers{
		router: mux.PathPrefix("/pile").Subrouter(),
	}
	r.newSubRouter = r.router.PathPrefix("/new").Methods("GET").Subrouter()
	r.shuffleSubRouter = r.router.PathPrefix("/shuffle").Methods("POST").Subrouter()
	r.drawSubRouter = r.router.PathPrefix("/draw").Methods("GET").Subrouter()
	return r
}

func newHandlers(deckManager *deck.Manager) *handlers {
	h := &handlers{
		newHandler:      handler.CreateNewHandler(deckManager),
		retrieveHandler: handler.CreateRetrieveHandler(deckManager),
		shuffleHandler:  handler.CreateShuffleHandler(deckManager),
		drawHandler:     handler.CreateDrawHandler(deckManager),
	}
	return h
}

//NewDeckAPI returns a pointer to a newly initialized deck.API.
func NewDeckAPI(mux *mux.Router, deckManager *deck.Manager) *API {
	api := &API{
		r:           newRouters(mux),
		h:           newHandlers(deckManager),
		deckManager: deckManager,
	}

	api.getRoute = api.r.router.Methods("GET")
	return api
}

//Register registers all paths required by the deck.API.
func (z *API) Register() {
	z.registerNewPaths()
	z.getRoute.Path("/{id}").Handler(z.h.retrieveHandler)
	z.r.shuffleSubRouter.Handle("/{id}", z.h.shuffleHandler)
	z.r.drawSubRouter.Handle("/{id}/{amount:[0-9]+}", z.h.drawHandler)
}

func (z *API) registerNewPaths() {
	z.registerShufflePaths()
	z.registerUnshuffledPaths()
}

func (z *API) registerShufflePaths() {
	z.r.newSubRouter.Handle("/", z.h.newHandler).Queries("shuffle", "{shuffle}")
	z.r.newSubRouter.Handle("/", z.h.newHandler).Queries("amount", "{amount}", "shuffle", "{shuffle}")
	z.r.newSubRouter.Handle("/", z.h.newHandler).Queries("jokers", "{jokers}", "shuffle", "{shuffle}")
	z.r.newSubRouter.Handle("/", z.h.newHandler).Queries("amount", "{amount}", "jokers", "{jokers},", "shuffle", "{shuffle}")
}

func (z *API) registerUnshuffledPaths() {
	z.r.newSubRouter.Handle("/", z.h.newHandler)
	z.r.newSubRouter.Handle("/", z.h.newHandler).Queries("amount", "{amount}")
	z.r.newSubRouter.Handle("/", z.h.newHandler).Queries("jokers", "{jokers}")
	z.r.newSubRouter.Handle("/", z.h.newHandler).Queries("amount", "{amount}", "jokers", "{jokers},")
}

package pile

import (
	"github.com/gorilla/mux"
	"github.com/heindrichpaul/card-api/api/pile/handler"
	"github.com/heindrichpaul/card-api/manager/pile"
)

type handlers struct {
	newHandler      *handler.NewHandler
	retrieveHandler *handler.RetrieveHandler
	shuffleHandler  *handler.ShuffleHandler
}

type routers struct {
	router           *mux.Router
	newSubRouter     *mux.Router
	shuffleSubRouter *mux.Router
}

//API is a struct the is a collection of all components required to create the API for a PileManager.
type API struct {
	r           *routers
	h           *handlers
	pileManager *pile.Manager
	getRoute    *mux.Route
}

func newRouters(mux *mux.Router) *routers {
	r := &routers{
		router: mux.PathPrefix("/pile").Subrouter(),
	}
	r.newSubRouter = r.router.PathPrefix("/new").Methods("GET").Subrouter()
	r.shuffleSubRouter = r.router.PathPrefix("/shuffle").Methods("POST").Subrouter()
	return r
}

func newHandlers(pileManager *pile.Manager) *handlers {
	h := &handlers{
		newHandler:      handler.CreateNewHandler(pileManager),
		retrieveHandler: handler.CreateRetrieveHandler(pileManager),
		shuffleHandler:  handler.CreateShuffleHandler(pileManager),
	}
	return h
}

//NewPileAPI returns a pointer to a newly initialized pile.API.
func NewPileAPI(mux *mux.Router, pileManager *pile.Manager) *API {
	api := &API{
		r:           newRouters(mux),
		pileManager: pileManager,
		h:           newHandlers(pileManager),
	}

	api.getRoute = api.r.router.Methods("GET")

	return api
}

//Register registers all paths required by the pile.API.
func (z *API) Register() {
	z.getRoute.Path("/{id}").Handler(z.h.retrieveHandler)
	z.r.shuffleSubRouter.Handle("/{id}", z.h.shuffleHandler)
	z.r.newSubRouter.Handle("/", z.h.newHandler)
}

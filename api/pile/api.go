package pile

import (
	"github.com/gorilla/mux"
	"github.com/heindrichpaul/card-api/api/pile/handler"
	"github.com/heindrichpaul/card-api/manager/pile"
)

//API is a struct the is a collection of all components required to create the API for a PileManager.
type API struct {
	router           *mux.Router
	pileManager      *pile.Manager
	getRoute         *mux.Route
	newSubRouter     *mux.Router
	shuffleSubRouter *mux.Router
	newHandler       *handler.NewHandler
	retrieveHandler  *handler.RetrieveHandler
	shuffleHandler   *handler.ShuffleHandler
}

//NewPileAPI returns a pointer to a newly initialized pile.API.
func NewPileAPI(mux *mux.Router, pileManager *pile.Manager) *API {
	api := &API{
		router:      mux.PathPrefix("/pile").Subrouter(),
		pileManager: pileManager,
	}
	api.getRoute = api.router.Methods("GET")
	api.newSubRouter = api.router.PathPrefix("/new").Methods("GET").Subrouter()
	api.shuffleSubRouter = api.router.PathPrefix("/shuffle").Methods("POST").Subrouter()
	api.newHandler = handler.CreateNewHandler(api.pileManager)
	api.retrieveHandler = handler.CreateRetrieveHandler(api.pileManager)
	api.shuffleHandler = handler.CreateShuffleHandler(api.pileManager)
	return api
}

//Register registers all paths required by the pile.API.
func (z *API) Register() {
	z.getRoute.Path("/{id}").Handler(z.retrieveHandler)
	z.shuffleSubRouter.Handle("/{id}", z.shuffleHandler)
	z.newSubRouter.Handle("/", z.newHandler)
}

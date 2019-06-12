package pile

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/heindrichpaul/card-api/apierror"
	"github.com/heindrichpaul/card-api/apiutilities"
	"github.com/heindrichpaul/card-api/manager/pile"
)

//API is a struct the is a collection of all components required to create the API for a PileManager.
type API struct {
	router      *mux.Router
	pileManager *pile.Manager
}

//NewPileAPI returns a pointer to a newly initialized pile.API.
func NewPileAPI(mux *mux.Router, pileManager *pile.Manager) *API {
	api := &API{
		router:      mux.PathPrefix("/pile").Subrouter(),
		pileManager: pileManager,
	}

	return api
}

//Register registers all paths required by the pile.API.
func (z *API) Register() {
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newPileHandler)
	z.router.Path("/{id}").Methods("GET").HandlerFunc(z.retrievePileHandler)
}

func (z *API) newPileHandler(w http.ResponseWriter, r *http.Request) {
	pile := z.pileManager.RequestNewPile()

	apiutilities.HandleResponse(w, pile)
}

func (z *API) retrievePileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pile := z.pileManager.FindPileByID(vars["id"])
	if pile == nil {
		e := apierror.NewAPIError(fmt.Sprintf("Could not find pile with id: %s", vars["id"]), apierror.NotFoundError)
		apiutilities.HandleResponse(w, e)
		return
	}
	apiutilities.HandleResponse(w, pile)
}

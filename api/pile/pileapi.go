package pile

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/heindrichpaul/card-api/apierror"
	"github.com/heindrichpaul/card-api/apiutilities"
	"github.com/heindrichpaul/card-api/manager/pile"
)

type PileAPI struct {
	router      *mux.Router
	pileManager *pile.PileManager
}

func NewPileAPI(mux *mux.Router, pileManager *pile.PileManager) *PileAPI {
	pAPI := &PileAPI{
		router:      mux.PathPrefix("/pile").Subrouter(),
		pileManager: pileManager,
	}

	return pAPI
}

func (z *PileAPI) Register() {
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newPileHandler)
	z.router.Path("/{id}").Methods("GET").HandlerFunc(z.retrievePileHandler)
}

func (z *PileAPI) newPileHandler(w http.ResponseWriter, r *http.Request) {
	pile := z.pileManager.RequestNewPile()

	apiutilities.HandleResponse(w, r, pile)
}

func (z *PileAPI) retrievePileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pile := z.pileManager.FindPileById(vars["id"])
	if pile == nil {
		e := apierror.NewAPIError(fmt.Sprintf("Could not find pile with id: %s", vars["id"]), apierror.NotFoundError)
		apiutilities.HandleResponse(w, r, e)
		return
	}
	apiutilities.HandleResponse(w, r, pile)
}

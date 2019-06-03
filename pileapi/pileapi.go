package pileapi

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/heindrichpaul/card-api/apiutilities"
	"github.com/heindrichpaul/card-api/pilemanager"
)

type PileAPI struct {
	router      *mux.Router
	pileManager *pilemanager.PileManager
}

func NewPileAPI(mux *mux.Router, pileManager *pilemanager.PileManager) *PileAPI {
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

	pileJSON, err := pile.Marshal()
	if err != nil {
		e := apiutilities.NewAPIError("Could not marshal pile", "1")
		apiutilities.HandleError(w, r, e)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(pileJSON))
}

func (z *PileAPI) retrievePileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pile := z.pileManager.FindPileById(vars["id"])
	if pile == nil {
		apiutilities.HandleError(w, r, apiutilities.NewAPIError(fmt.Sprintf("Could not find pile with id: %s", vars["id"]), "1"))
		return
	}

	pileJSON, err := pile.Marshal()
	if err != nil {
		apiutilities.HandleError(w, r, apiutilities.NewAPIError("Could not marshal pile", "1"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(pileJSON))

}

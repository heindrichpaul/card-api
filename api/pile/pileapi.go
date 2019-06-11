package pile

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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
	/*pileJSON, err := pile.Marshal()
	if err != nil {
		e := apiutilities.NewAPIError("Could not marshal pile", "1")
		apiutilities.HandleError(w, r, e)
		return
	}*/

}

func (z *PileAPI) retrievePileHandler(w http.ResponseWriter, r *http.Request) {
	var x apiutilities.Marshalable
	defer apiutilities.HandleResponse(w, r, x)
	vars := mux.Vars(r)
	x = z.pileManager.FindPileById(vars["id"])
	if x == nil {
		x = apiutilities.NewAPIError(fmt.Sprintf("Could not find pile with id: %s", vars["id"]), "1")
		return
	}

	/*pileJSON, err := pile.Marshal()
	if err != nil {
		apiutilities.HandleError(w, r, apiutilities.NewAPIError("Could not marshal pile", "1"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(pileJSON))*/

}

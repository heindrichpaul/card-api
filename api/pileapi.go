package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/heindrichpaul/card-api/deckmanager"
)

type pileAPI struct {
	router      *mux.Router
	deckManager *deckmanager.DeckManager
}

func newPileAPI(mux *mux.Router, deckM *deckmanager.DeckManager) *pileAPI {
	pAPI := &pileAPI{
		router:      mux.PathPrefix("/pile").Subrouter(),
		deckManager: deckM,
	}

	return pAPI
}

func (z *pileAPI) register() {
	z.router.Path("/new").Methods("GET").HandlerFunc(z.newPileHandler)
	z.router.Path("/{id}").Methods("GET").HandlerFunc(z.retrievePileHandler)
}

func (z *pileAPI) newPileHandler(w http.ResponseWriter, r *http.Request) {
	pile := z.deckManager.RequestNewPile()

	pileJSON, err := pile.Marshal()
	if err != nil {
		e := newAPIError("Could not marshal pile", "1")
		handleError(w, r, e)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(pileJSON))
}

func (z *pileAPI) retrievePileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deck := z.deckManager.FindDeckById(id)
	if deck == nil {
		e := newAPIError(fmt.Sprintf("Could not find deck with id: %s", id), "1")
		handleError(w, r, e)
		return
	}

	deckJSON, err := deck.Marshal()
	if err != nil {
		e := newAPIError("Could not marshal deck", "1")
		handleError(w, r, e)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(deckJSON))

}

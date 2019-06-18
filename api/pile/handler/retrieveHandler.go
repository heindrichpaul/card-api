package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/heindrichpaul/card-api/apierror"
	"github.com/heindrichpaul/card-api/apiutilities"
	"github.com/heindrichpaul/card-api/manager/pile"
)

type RetrieveHandler struct {
	pileManager *pile.Manager
}

func CreateRetrieveHandler(manager *pile.Manager) *RetrieveHandler {
	z := &RetrieveHandler{
		pileManager: manager,
	}
	return z
}

func (z *RetrieveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pile := z.pileManager.FindPileByID(vars["id"])
	if pile == nil {
		e := apierror.NewAPIError(fmt.Sprintf("Could not find pile with id: %s", vars["id"]), apierror.NotFoundError)
		apiutilities.HandleResponse(w, e)
		return
	}
	apiutilities.HandleResponse(w, pile)
}

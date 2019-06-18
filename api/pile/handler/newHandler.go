package handler

import (
	"net/http"

	"github.com/heindrichpaul/card-api/apiutilities"
	"github.com/heindrichpaul/card-api/manager/pile"
)

type NewHandler struct {
	pileManager *pile.Manager
}

func CreateNewHandler(manager *pile.Manager) *NewHandler {
	z := &NewHandler{
		pileManager: manager,
	}
	return z
}

func (z *NewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pile := z.pileManager.RequestNewPile()
	apiutilities.HandleResponse(w, pile)
}

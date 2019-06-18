package handler

import (
	"net/http"

	"github.com/heindrichpaul/card-api/apiutilities"
	"github.com/heindrichpaul/card-api/manager/deck"
)

type RetrieveHandler struct {
	deckManager *deck.Manager
}

func CreateRetrieveHandler(manager *deck.Manager) *RetrieveHandler {
	z := &RetrieveHandler{
		deckManager: manager,
	}
	return z
}

func (z *RetrieveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := apiutilities.GetIDFromRequest(r)
	if z.deckManager.DoesDeckExist(id) {
		deck := z.deckManager.FindDeckByID(id)
		apiutilities.HandleResponse(w, deck)
	}
}

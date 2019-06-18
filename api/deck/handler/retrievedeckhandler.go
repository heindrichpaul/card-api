package handler

import (
	"net/http"

	"github.com/heindrichpaul/card-api/apiutilities"
	"github.com/heindrichpaul/card-api/manager/deck"
)

type RetrieveDeckHandler struct {
	deckManager *deck.Manager
}

func CreateRetrieveDeckHandler(manager *deck.Manager) *RetrieveDeckHandler {
	z := &RetrieveDeckHandler{
		deckManager: manager,
	}
	return z
}

func (z *RetrieveDeckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := apiutilities.GetIDFromRequest(r)
	if z.deckManager.DoesDeckExist(id) {
		deck := z.deckManager.FindDeckByID(id)
		apiutilities.HandleResponse(w, deck)
	}
}

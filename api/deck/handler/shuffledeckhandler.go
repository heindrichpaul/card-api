package handler

import (
	"net/http"

	"github.com/heindrichpaul/card-api/apiutilities"
	"github.com/heindrichpaul/card-api/manager/deck"
)

type ShuffleDeckHandler struct {
	deckManager *deck.Manager
}

func CreateShuffleDeckHandler(manager *deck.Manager) *ShuffleDeckHandler {
	z := &ShuffleDeckHandler{
		deckManager: manager,
	}
	return z
}

func (z *ShuffleDeckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := apiutilities.GetIDFromRequest(r)
	if z.deckManager.DoesDeckExist(id) {
		deck := z.deckManager.FindDeckByID(id)
		deck = z.deckManager.ReshuffleDeck(deck)
		apiutilities.HandleResponse(w, deck)
	}
}

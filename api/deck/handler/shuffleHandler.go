package handler

import (
	"net/http"

	"github.com/heindrichpaul/card-api/apiutilities"
	"github.com/heindrichpaul/card-api/manager/deck"
)

type ShuffleHandler struct {
	deckManager *deck.Manager
}

func CreateShuffleHandler(manager *deck.Manager) *ShuffleHandler {
	z := &ShuffleHandler{
		deckManager: manager,
	}
	return z
}

func (z *ShuffleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := apiutilities.GetIDFromRequest(r)
	if z.deckManager.DoesDeckExist(id) {
		deck := z.deckManager.FindDeckByID(id)
		deck = z.deckManager.ReshuffleDeck(deck)
		apiutilities.HandleResponse(w, deck)
	}
}

package handler

import (
	"net/http"

	"github.com/heindrichpaul/card-api/apiutilities"
	"github.com/heindrichpaul/card-api/manager/pile"
)

type ShuffleHandler struct {
	pileManager *pile.Manager
}

func CreateShuffleHandler(manager *pile.Manager) *ShuffleHandler {
	z := &ShuffleHandler{
		pileManager: manager,
	}
	return z
}

func (z *ShuffleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := apiutilities.GetIDFromRequest(r)
	if z.pileManager.DoesPileExist(id) {
		pile := z.pileManager.FindPileByID(id)
		pile = z.pileManager.ReshufflePile(pile)
		apiutilities.HandleResponse(w, pile)
	}
}

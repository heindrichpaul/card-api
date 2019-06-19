package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/heindrichpaul/deckofcards"

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
	view := apiutilities.GetBooleanValue(r.URL.Query(), "view")
	id := apiutilities.GetIDFromRequest(r)
	var pile *deckofcards.Pile
	var cards deckofcards.Cards
	switch view {
	case true:
		cards = z.pileManager.RetrieveCardsInPile(id)
		if cards == nil {
			break
		}
		json.NewEncoder(w).Encode(cards)
		return
	default:
		pile = z.pileManager.FindPileByID(id)
		if pile == nil {
			break
		}
		apiutilities.HandleResponse(w, pile)
		return
	}
	e := apierror.NewAPIError(fmt.Sprintf("Could not find pile with id: %s", id), apierror.NotFoundError)
	apiutilities.HandleResponse(w, e)

}

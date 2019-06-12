package deck

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/heindrichpaul/card-api/apierror"
	"github.com/heindrichpaul/card-api/apiutilities"
	"github.com/heindrichpaul/deckofcards"
)

func (z *API) createDeck(shuffle, jokers bool, amountOfDecks int) (deck *deckofcards.Deck) {
	if shuffle {
		if jokers {
			deck = z.deckManager.RequestNumberOfShuffledDecksWithJokers(amountOfDecks)
		} else {
			deck = z.deckManager.RequestNumberOfShuffledDecks(amountOfDecks)
		}
	} else {
		if jokers {
			deck = z.deckManager.RequestNumberOfDecksWithJokers(amountOfDecks)
		} else {
			deck = z.deckManager.RequestNumberOfDecks(amountOfDecks)
		}
	}
	return
}

func (z *API) findAndValidateDeck(w http.ResponseWriter, id string) (deck *deckofcards.Deck, ok bool) {
	deck = z.deckManager.FindDeckByID(id)
	if deck == nil {
		e := apierror.NewAPIError(fmt.Sprintf("Could not find deck with id: %s", id), apierror.NotFoundError)
		apiutilities.HandleResponse(w, e)
		return nil, false
	}
	return deck, true
}

func (z *API) getNewDeckQueryValues(r *http.Request) (shuffle bool, jokers bool, amount int) {
	amount = apiutilities.GetIntWithDefaultValueOfOne(r.URL.Query(), "amount")
	jokers = apiutilities.GetBooleanValue(r.URL.Query(), "jokers")
	shuffle = apiutilities.GetBooleanValue(r.URL.Query(), "shuffle")
	return
}

func (z *API) getIDFromRequest(r *http.Request) (id string) {
	vars := mux.Vars(r)
	id = vars["id"]
	return
}

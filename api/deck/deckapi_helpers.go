package deck

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/heindrichpaul/card-api/apierror"
	"github.com/heindrichpaul/card-api/apiutilities"
	"github.com/heindrichpaul/deckofcards"
)

func (z *DeckAPI) createDeck(shuffle, jokers bool, amountOfDecks int) (deck *deckofcards.Deck) {
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

func (z *DeckAPI) findAndValidateDeck(w http.ResponseWriter, r *http.Request, id string) (deck *deckofcards.Deck, ok bool) {
	deck = z.deckManager.FindDeckById(id)
	if deck == nil {
		e := apierror.NewAPIError(fmt.Sprintf("Could not find deck with id: %s", id), apierror.NotFoundError)
		apiutilities.HandleResponse(w, r, e)
		return nil, false
	}
	return deck, true
}

func (z *DeckAPI) getNewDeckQueryValues(r *http.Request) (shuffle bool, jokers bool, amount int) {
	amount = apiutilities.GetIntWithDefaultValueOfOne(r.URL.Query(), "amount")
	jokers = apiutilities.GetBooleanValue(r.URL.Query(), "jokers")
	shuffle = apiutilities.GetBooleanValue(r.URL.Query(), "shuffle")
	return
}

func (z *DeckAPI) getIDFromRequest(r *http.Request) (id string) {
	vars := mux.Vars(r)
	id = vars["id"]
	return
}

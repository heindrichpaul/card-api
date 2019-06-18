package handler

import (
	"net/http"

	"github.com/heindrichpaul/card-api/apiutilities"
	"github.com/heindrichpaul/card-api/manager/deck"
	"github.com/heindrichpaul/deckofcards"
)

type NewDeckHandler struct {
	deckManager *deck.Manager
}

func CreateNewDeckHandler(manager *deck.Manager) *NewDeckHandler {
	z := &NewDeckHandler{
		deckManager: manager,
	}
	return z
}

func (z *NewDeckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	deck := z.createDeck(z.getNewDeckQueryValues(r))
	apiutilities.HandleResponse(w, deck)
}

func (z *NewDeckHandler) createDeck(shuffle, jokers bool, amountOfDecks int) (deck *deckofcards.Deck) {
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

func (z *NewDeckHandler) getNewDeckQueryValues(r *http.Request) (shuffle bool, jokers bool, amount int) {
	amount = apiutilities.GetIntWithDefaultValueOfOne(r.URL.Query(), "amount")
	jokers = apiutilities.GetBooleanValue(r.URL.Query(), "jokers")
	shuffle = apiutilities.GetBooleanValue(r.URL.Query(), "shuffle")
	return
}

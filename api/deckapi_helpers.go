package api

import (
	"fmt"
	"net/http"

	"github.com/heindrichpaul/deckofcards"
)

func (z *deckAPI) createDeck(shuffle, jokers bool, amountOfDecks int) (deck *deckofcards.Deck) {
	if shuffle {
		deck = z.createShuffledDeck(amountOfDecks, jokers)
	} else {
		deck = z.createUnshuffledDeck(amountOfDecks, jokers)
	}
	return
}

func (z *deckAPI) createUnshuffledDeck(amount int, jokers bool) (deck *deckofcards.Deck) {
	if jokers {
		deck = z.deckManager.RequestNumberOfDecksWithJokers(amount)
	} else {
		deck = z.deckManager.RequestNumberOfDecks(amount)
	}
	return
}

func (z *deckAPI) createShuffledDeck(amount int, jokers bool) (deck *deckofcards.Deck) {
	if jokers {
		deck = z.deckManager.RequestNumberOfShuffledDecksWithJokers(amount)
	} else {
		deck = z.deckManager.RequestNumberOfShuffledDecks(amount)
	}
	return
}

func (z *deckAPI) findAndValidateDeck(w http.ResponseWriter, r *http.Request, id string) (deck *deckofcards.Deck, ok bool) {
	deck = z.deckManager.FindDeckById(id)
	if deck == nil {
		e := newAPIError(fmt.Sprintf("Could not find deck with id: %s", id), "1")
		handleError(w, r, e)
		return nil, false
	}
	return deck, true
}

func (z *deckAPI) marshalDeckAndValidate(w http.ResponseWriter, r *http.Request, deck *deckofcards.Deck) (json string, ok bool) {
	deckJSON, err := deck.Marshal()
	if err != nil {
		e := newAPIError("Could not marshal deck", "1")
		handleError(w, r, e)
		return "", false
	}
	return string(deckJSON), true
}

func getQueryValues(r *http.Request) (shuffle bool, jokers bool, amount int) {
	amount = getIntWithDefaultValueAs1(r.URL.Query(), "amount")
	jokers = getBooleanValue(r.URL.Query(), "jokers")
	shuffle = getBooleanValue(r.URL.Query(), "shuffle")
	return
}

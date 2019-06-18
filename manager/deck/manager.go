package deck

import (
	"github.com/heindrichpaul/card-api/interfaces"
	"github.com/heindrichpaul/deckofcards"
)

//Manager is a struct that handles all functionality around decks.
type Manager struct {
	persistanceManger interfaces.PersistenceManager
}

//NewDeckManager returns a pointer to a new deck.Manager struct
func NewDeckManager(p interfaces.PersistenceManager) *Manager {
	d := &Manager{
		persistanceManger: p,
	}
	return d
}

//RequestNumberOfDecks returns a deck with the amount of cards for the amount of requested decks.
func (z *Manager) RequestNumberOfDecks(number int) *deckofcards.Deck {
	deck := deckofcards.NewDeck(number)
	if deck.Success {
		z.persistanceManger.PersistDeck(deck)
	}
	return deck
}

//RequestNumberOfShuffledDecks returns a shuffled deck with the amount of cards for the amount of requested decks.
func (z *Manager) RequestNumberOfShuffledDecks(number int) *deckofcards.Deck {
	deck := z.RequestNumberOfDecks(number)
	if deck.Success {
		deck = deckofcards.ShuffleDeck(deck)
		z.persistanceManger.PersistDeck(deck)
	}
	return deck
}

//RequestNumberOfDecksWithJokers returns a deck with the amount of cards for the amount of requested decks. It also includes two jokers per requested deck.
func (z *Manager) RequestNumberOfDecksWithJokers(number int) *deckofcards.Deck {
	deck := deckofcards.NewDeckWithJokers(number)
	if deck.Success {
		z.persistanceManger.PersistDeck(deck)
	}
	return deck
}

//RequestNumberOfShuffledDecksWithJokers returns a shuffled deck with the amount of cards for the amount of requested decks. It also includes two jokers per requested deck.
func (z *Manager) RequestNumberOfShuffledDecksWithJokers(number int) *deckofcards.Deck {
	deck := z.RequestNumberOfDecksWithJokers(number)
	if deck.Success {
		deck = deckofcards.ShuffleDeck(deck)
		z.persistanceManger.PersistDeck(deck)
	}
	return deck
}

//ReshuffleDeck reshuffles an existing deck.
func (z *Manager) ReshuffleDeck(deck *deckofcards.Deck) *deckofcards.Deck {
	deck = deckofcards.ShuffleDeck(deck)
	if deck.Success {
		z.persistanceManger.PersistDeck(deck)
	}
	return deck

}

//FindDeckByID returns the deck with the given ID
func (z *Manager) FindDeckByID(ID string) *deckofcards.Deck {
	deck, ok := z.persistanceManger.RetrieveDeck(ID)
	if !ok {
		return nil
	}
	return deck
}

//DrawFromDeck draw the amount of cards from the deck with the given ID.
func (z *Manager) DrawFromDeck(ID string, amount int) *deckofcards.Draw {
	deck, ok := z.persistanceManger.RetrieveDeck(ID)
	defer z.persistanceManger.PersistDeck(deck)
	if !ok {
		return nil
	}
	draw := deck.Draw(amount)
	return draw
}

//DoesDeckExist returns true if the deck with the given ID exists.
func (z *Manager) DoesDeckExist(ID string) bool {
	_, ok := z.persistanceManger.RetrieveDeck(ID)
	return ok
}

package deckmanager

import (
	persistancemanager "github.com/heindrichpaul/card-api/persistanceManager"
	"github.com/heindrichpaul/deckofcards"
)

type DeckManager struct {
	persistanceManger *persistancemanager.PersistanceManage
}

func NewDeckManager() *DeckManager {
	d := &DeckManager{
		persistanceManger: persistancemanager.NewPersistanceManager(),
	}
	return d
}

func (z *DeckManager) RequestNumberOfDecks(number int) *deckofcards.Deck {
	deck := deckofcards.NewDeck(number)
	if deck.Success {
		z.persistanceManger.PersistDeck(deck)
	}
	return deck
}

func (z *DeckManager) RequestNumberOfShuffledDecks(number int) *deckofcards.Deck {
	deck := z.RequestNumberOfDecks(number)
	if deck.Success {
		deck = deckofcards.ShuffleDeck(deck)
		z.persistanceManger.PersistDeck(deck)
	}
	return deck
}

func (z *DeckManager) RequestNumberOfDecksWithJokers(number int) *deckofcards.Deck {
	deck := deckofcards.NewDeckWithJokers(number)
	if deck.Success {
		z.persistanceManger.PersistDeck(deck)
	}
	return deck
}

func (z *DeckManager) RequestNumberOfShuffledDecksWithJokers(number int) *deckofcards.Deck {
	deck := z.RequestNumberOfDecksWithJokers(number)
	if deck.Success {
		deck = deckofcards.ShuffleDeck(deck)
		z.persistanceManger.PersistDeck(deck)
	}
	return deck
}

func (z *DeckManager) ReshuffleDeck(deck *deckofcards.Deck) *deckofcards.Deck {
	deck = deckofcards.ShuffleDeck(deck)
	if deck.Success {
		z.persistanceManger.PersistDeck(deck)
	}
	return deck

}

func (z *DeckManager) RequestSingleUnshuffledDeck() *deckofcards.Deck {
	deck := deckofcards.NewDeck(1)
	if deck.Success {
		z.persistanceManger.PersistDeck(deck)
	}
	return deck
}

func (z *DeckManager) FindDeckById(Id string) *deckofcards.Deck {
	deck, ok := z.persistanceManger.RetrieveDeck(Id)
	if !ok {
		return nil
	}
	return deck
}

func (z *DeckManager) DrawFromDeck(Id string, amount int) *deckofcards.Draw {
	deck, ok := z.persistanceManger.RetrieveDeck(Id)
	defer z.persistanceManger.PersistDeck(deck)
	if !ok {
		return nil
	}
	draw := deck.Draw(amount)
	return draw
}

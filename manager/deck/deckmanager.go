package deck

import (
	persistancemanager "github.com/heindrichpaul/card-api/manager/persistance"
	"github.com/heindrichpaul/deckofcards"
)

type DeckManager struct {
	persistanceManger *persistancemanager.PersistanceManage
}

func NewDeckManager(p *persistancemanager.PersistanceManage) *DeckManager {
	d := &DeckManager{
		persistanceManger: p,
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

func (z *DeckManager) DoesDeckExist(Id string) bool {
	_, ok := z.persistanceManger.RetrieveDeck(Id)
	return ok
}

package persistance

import "github.com/heindrichpaul/deckofcards"

type PersistanceManage struct {
	piles map[string]*deckofcards.Pile
	decks map[string]*deckofcards.Deck
}

func NewPersistanceManager() *PersistanceManage {
	p := &PersistanceManage{
		piles: make(map[string]*deckofcards.Pile),
		decks: make(map[string]*deckofcards.Deck),
	}
	return p
}

func (z *PersistanceManage) PersistDeck(deck *deckofcards.Deck) bool {
	z.decks[deck.DeckID] = deck
	return true
}

func (z *PersistanceManage) RetrieveDeck(id string) (*deckofcards.Deck, bool) {
	deck, ok := z.decks[id]
	return deck, ok
}

func (z *PersistanceManage) PersistPile(pile *deckofcards.Pile) bool {
	z.piles[pile.PileID] = pile
	return true
}

func (z *PersistanceManage) RetrievePile(id string) (*deckofcards.Pile, bool) {
	pile, ok := z.piles[id]
	return pile, ok
}

package persistence

import "github.com/heindrichpaul/deckofcards"

//MapManager is a struct that implements the interfaces.PersistenceManager interface with a internal map as a way of persistence.
type MapManager struct {
	piles map[string]*deckofcards.Pile
	decks map[string]*deckofcards.Deck
}

//NewMapPersistenceManager returns a pointer to a newly created MapManager.
func NewMapPersistenceManager() *MapManager {
	p := &MapManager{
		piles: make(map[string]*deckofcards.Pile),
		decks: make(map[string]*deckofcards.Deck),
	}
	return p
}

//PersistDeck persists the deck to the map.
func (z *MapManager) PersistDeck(deck *deckofcards.Deck) bool {
	z.decks[deck.DeckID] = deck
	return true
}

//RetrieveDeck retrieves the deck with the given ID from the map.
func (z *MapManager) RetrieveDeck(ID string) (*deckofcards.Deck, bool) {
	deck, ok := z.decks[ID]
	return deck, ok
}

//PersistPile persists the pile to the map.
func (z *MapManager) PersistPile(pile *deckofcards.Pile) bool {
	z.piles[pile.PileID] = pile
	return true
}

//RetrievePile retrieves the pile with the given ID from the map.
func (z *MapManager) RetrievePile(ID string) (*deckofcards.Pile, bool) {
	pile, ok := z.piles[ID]
	return pile, ok
}

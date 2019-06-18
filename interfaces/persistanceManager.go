package interfaces

import "github.com/heindrichpaul/deckofcards"

//PersistanceManager is an interface that can be used to swap to different ways of persistance
type PersistanceManager interface {
	PersistDeck(deck *deckofcards.Deck) bool
	RetrieveDeck(id string) (*deckofcards.Deck, bool)
	PersistPile(pile *deckofcards.Pile) bool
	RetrievePile(id string) (*deckofcards.Pile, bool)
}

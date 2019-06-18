package pile

import (
	"fmt"

	"github.com/heindrichpaul/card-api/interfaces"
	"github.com/heindrichpaul/deckofcards"
)

//Manager is a struct that handles all functionality around piles.
type Manager struct {
	persistanceManger interfaces.PersistenceManager
}

//NewPileManager returns a pointer to a new pile.Manager struct
func NewPileManager(p interfaces.PersistenceManager) *Manager {
	d := &Manager{
		persistanceManger: p,
	}
	return d
}

//RequestNewPile returns a new pile.
func (z *Manager) RequestNewPile() *deckofcards.Pile {
	pile := deckofcards.NewPile()
	z.persistanceManger.PersistPile(pile)
	return pile
}

//ReshufflePile shuffles an existing pile.
func (z *Manager) ReshufflePile(pile *deckofcards.Pile) *deckofcards.Pile {
	defer z.persistanceManger.PersistPile(pile)
	pile = deckofcards.ShufflePile(pile)
	return pile
}

//AddCardsToPile adds the cards from the given draw to the pile with the given ID.
func (z *Manager) AddCardsToPile(ID string, draw *deckofcards.Draw, cards deckofcards.Cards) {
	pile, ok := z.persistanceManger.RetrievePile(ID)
	if ok {
		defer z.persistanceManger.PersistPile(pile)
		pile.AddCardsToPile(draw, cards)
	}
}

//GetCardAtIDFromPile returns a card at the given index in a pile with the given ID.
func (z *Manager) GetCardAtIDFromPile(ID string, index int) (*deckofcards.Draw, error) {
	pile, ok := z.persistanceManger.RetrievePile(ID)
	if !ok {
		return nil, fmt.Errorf("Unable to retrieve the pile")
	}
	defer z.persistanceManger.PersistPile(pile)
	return pile.GetCardAtID(index)
}

//GetCardsFromPile returns the requested cards from the pile with the given ID.
func (z *Manager) GetCardsFromPile(ID string, cards deckofcards.Cards) *deckofcards.Draw {
	pile, ok := z.persistanceManger.RetrievePile(ID)
	if !ok {
		return nil
	}
	defer z.persistanceManger.PersistPile(pile)
	return pile.GetCardsFromPile(cards)
}

//GetAllCardsFromPile returns all cards from the pile with the given ID.
func (z *Manager) GetAllCardsFromPile(ID string) *deckofcards.Draw {
	pile, ok := z.persistanceManger.RetrievePile(ID)
	if !ok {
		return nil
	}
	defer z.persistanceManger.PersistPile(pile)
	return pile.PickAllCardsFromPile()
}

//GetAmountOfCardsFromBottomOfPile returns the requested amount of cards from the bottom of the pile with the given ID.
func (z *Manager) GetAmountOfCardsFromBottomOfPile(ID string, amount int) *deckofcards.Draw {
	pile, ok := z.persistanceManger.RetrievePile(ID)
	if !ok {
		return nil
	}
	defer z.persistanceManger.PersistPile(pile)
	return pile.PickAmountOfCardsFromBottomOfPile(amount)
}

//GetAmountOfCardsFromBTopOfPile returns the requested amount of cards from the top of the pile with the given ID.
func (z *Manager) GetAmountOfCardsFromBTopOfPile(ID string, amount int) *deckofcards.Draw {
	pile, ok := z.persistanceManger.RetrievePile(ID)
	if !ok {
		return nil
	}
	defer z.persistanceManger.PersistPile(pile)
	return pile.PickAmountOfCardsFromTopOfPile(amount)
}

//RetrieveCardsInPile returns an array of all cards in the list. It is not a way of drawing cards from a pile. It should be used to view the contents of the pile with the given ID.
func (z *Manager) RetrieveCardsInPile(ID string) deckofcards.Cards {
	pile, ok := z.persistanceManger.RetrievePile(ID)
	if !ok {
		return nil
	}
	defer z.persistanceManger.PersistPile(pile)
	return pile.RetrieveCardsInPile()
}

//FindPileByID returns the pile with the given ID.
func (z *Manager) FindPileByID(ID string) *deckofcards.Pile {
	pile, ok := z.persistanceManger.RetrievePile(ID)
	if !ok {
		return nil
	}
	return pile
}

//DoesPileExist returns true if the pile with the given ID exists.
func (z *Manager) DoesPileExist(ID string) bool {
	_, ok := z.persistanceManger.RetrieveDeck(ID)
	return ok
}

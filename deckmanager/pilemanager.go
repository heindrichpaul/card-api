package deckmanager

import (
	"fmt"

	"github.com/heindrichpaul/deckofcards"
)

func (z *DeckManager) RequestNewPile() *deckofcards.Pile {
	pile := deckofcards.NewPile()
	z.persistanceManger.PersistPile(pile)
	return pile
}

func (z *DeckManager) ShufflePile(Id string) *deckofcards.Pile {
	pile, ok := z.persistanceManger.RetrievePile(Id)
	if !ok {
		return nil
	}
	defer z.persistanceManger.PersistPile(pile)
	pile = deckofcards.ShufflePile(pile)
	return pile
}

func (z *DeckManager) AddCardsToPile(Id string, draw *deckofcards.Draw, cards deckofcards.Cards) {
	pile, ok := z.persistanceManger.RetrievePile(Id)
	if ok {
		defer z.persistanceManger.PersistPile(pile)
		pile.AddCardsToPile(draw, cards)
	}
}

func (z *DeckManager) GetCardAtIDFromPile(Id string, index int) (*deckofcards.Draw, error) {
	pile, ok := z.persistanceManger.RetrievePile(Id)
	if !ok {
		return nil, fmt.Errorf("Unable to retrieve the pile")
	}
	defer z.persistanceManger.PersistPile(pile)
	return pile.GetCardAtID(index)
}

func (z *DeckManager) GetCardsFromPile(Id string, cards deckofcards.Cards) *deckofcards.Draw {
	pile, ok := z.persistanceManger.RetrievePile(Id)
	if !ok {
		return nil
	}
	defer z.persistanceManger.PersistPile(pile)
	return pile.GetCardsFromPile(cards)
}

func (z *DeckManager) PickAllCardsFromPile(Id string) *deckofcards.Draw {
	pile, ok := z.persistanceManger.RetrievePile(Id)
	if !ok {
		return nil
	}
	defer z.persistanceManger.PersistPile(pile)
	return pile.PickAllCardsFromPile()
}

func (z *DeckManager) PickAmountOfCardsFromBottomOfPile(Id string, amount int) *deckofcards.Draw {
	pile, ok := z.persistanceManger.RetrievePile(Id)
	if !ok {
		return nil
	}
	defer z.persistanceManger.PersistPile(pile)
	return pile.PickAmountOfCardsFromBottomOfPile(amount)
}

func (z *DeckManager) PickAmountOfCardsFromBTopOfPile(Id string, amount int) *deckofcards.Draw {
	pile, ok := z.persistanceManger.RetrievePile(Id)
	if !ok {
		return nil
	}
	defer z.persistanceManger.PersistPile(pile)
	return pile.PickAmountOfCardsFromTopOfPile(amount)
}

func (z *DeckManager) RetrieveCardsInPile(Id string) deckofcards.Cards {
	pile, ok := z.persistanceManger.RetrievePile(Id)
	if !ok {
		return nil
	}
	defer z.persistanceManger.PersistPile(pile)
	return pile.RetrieveCardsInPile()
}

func (z *DeckManager) FindPileById(Id string) *deckofcards.Pile {
	pile, ok := z.persistanceManger.RetrievePile(Id)
	if !ok {
		return nil
	}
	return pile
}

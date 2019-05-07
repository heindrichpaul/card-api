package deckmanager

import "github.com/heindrichpaul/deckofcards"

func (z *DeckManager) RequestNewPile() *deckofcards.Pile {
	pile := deckofcards.NewPile()
	z.persistanceManger.PersistPile(pile)
	return pile
}

func (z *DeckManager) ShufflePile(pile *deckofcards.Pile) *deckofcards.Pile {
	pile = deckofcards.ShufflePile(pile)
	z.persistanceManger.PersistPile(pile)
	return pile
}

func (z *DeckManager) AddCardsToPile(pile *deckofcards.Pile, draw *deckofcards.Draw, cards deckofcards.Cards) {
	pile.AddCardsToPile(draw, cards)
}

func (z *DeckManager) GetCardAtIDFromPile(pile *deckofcards.Pile, index int) (*deckofcards.Draw, error) {
	return pile.GetCardAtID(index)
}

func (z *DeckManager) GetCardsFromPile(pile *deckofcards.Pile, cards deckofcards.Cards) *deckofcards.Draw {
	return pile.GetCardsFromPile(cards)
}

func (z *DeckManager) PickAllCardsFromPile(pile *deckofcards.Pile) *deckofcards.Draw {
	return pile.PickAllCardsFromPile()
}

func (z *DeckManager) PickAmountOfCardsFromBottomOfPile(pile *deckofcards.Pile, amount int) *deckofcards.Draw {
	return pile.PickAmountOfCardsFromBottomOfPile(amount)
}

func (z *DeckManager) PickAmountOfCardsFromBTopOfPile(pile *deckofcards.Pile, amount int) *deckofcards.Draw {
	return pile.PickAmountOfCardsFromTopOfPile(amount)
}

func (z *DeckManager) RetrieveCardsInPile(pile *deckofcards.Pile) deckofcards.Cards {
	return pile.RetrieveCardsInPile()
}

func (z *DeckManager) FindPileById(Id string) *deckofcards.Pile {
	pile, ok := z.persistanceManger.RetrievePile(Id)
	if !ok {
		return nil
	}
	return pile
}

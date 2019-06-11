package pile

import (
	"fmt"

	persistancemanager "github.com/heindrichpaul/card-api/manager/persistance"
	"github.com/heindrichpaul/deckofcards"
)

type PileManager struct {
	persistanceManger *persistancemanager.PersistanceManage
}

func NewPileManager(p *persistancemanager.PersistanceManage) *PileManager {
	d := &PileManager{
		persistanceManger: p,
	}
	return d
}

func (z *PileManager) RequestNewPile() *deckofcards.Pile {
	pile := deckofcards.NewPile()
	z.persistanceManger.PersistPile(pile)
	return pile
}

func (z *PileManager) ShufflePile(Id string) *deckofcards.Pile {
	pile, ok := z.persistanceManger.RetrievePile(Id)
	if !ok {
		return nil
	}
	defer z.persistanceManger.PersistPile(pile)
	pile = deckofcards.ShufflePile(pile)
	return pile
}

func (z *PileManager) AddCardsToPile(Id string, draw *deckofcards.Draw, cards deckofcards.Cards) {
	pile, ok := z.persistanceManger.RetrievePile(Id)
	if ok {
		defer z.persistanceManger.PersistPile(pile)
		pile.AddCardsToPile(draw, cards)
	}
}

func (z *PileManager) GetCardAtIDFromPile(Id string, index int) (*deckofcards.Draw, error) {
	pile, ok := z.persistanceManger.RetrievePile(Id)
	if !ok {
		return nil, fmt.Errorf("Unable to retrieve the pile")
	}
	defer z.persistanceManger.PersistPile(pile)
	return pile.GetCardAtID(index)
}

func (z *PileManager) GetCardsFromPile(Id string, cards deckofcards.Cards) *deckofcards.Draw {
	pile, ok := z.persistanceManger.RetrievePile(Id)
	if !ok {
		return nil
	}
	defer z.persistanceManger.PersistPile(pile)
	return pile.GetCardsFromPile(cards)
}

func (z *PileManager) PickAllCardsFromPile(Id string) *deckofcards.Draw {
	pile, ok := z.persistanceManger.RetrievePile(Id)
	if !ok {
		return nil
	}
	defer z.persistanceManger.PersistPile(pile)
	return pile.PickAllCardsFromPile()
}

func (z *PileManager) PickAmountOfCardsFromBottomOfPile(Id string, amount int) *deckofcards.Draw {
	pile, ok := z.persistanceManger.RetrievePile(Id)
	if !ok {
		return nil
	}
	defer z.persistanceManger.PersistPile(pile)
	return pile.PickAmountOfCardsFromBottomOfPile(amount)
}

func (z *PileManager) PickAmountOfCardsFromBTopOfPile(Id string, amount int) *deckofcards.Draw {
	pile, ok := z.persistanceManger.RetrievePile(Id)
	if !ok {
		return nil
	}
	defer z.persistanceManger.PersistPile(pile)
	return pile.PickAmountOfCardsFromTopOfPile(amount)
}

func (z *PileManager) RetrieveCardsInPile(Id string) deckofcards.Cards {
	pile, ok := z.persistanceManger.RetrievePile(Id)
	if !ok {
		return nil
	}
	defer z.persistanceManger.PersistPile(pile)
	return pile.RetrieveCardsInPile()
}

func (z *PileManager) FindPileById(Id string) *deckofcards.Pile {
	pile, ok := z.persistanceManger.RetrievePile(Id)
	if !ok {
		return nil
	}
	return pile
}

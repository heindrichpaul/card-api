package domain

import (
	"fmt"
	"strings"

	"github.com/twinj/uuid"
)

type Pile struct {
	stack  []*pileObject
	PileID string
}

type pileObject struct {
	deckID      string
	stillInPile bool
	card        *Card
}

func (z *pileObject) String() string {
	var printString []string
	printString = append(printString, fmt.Sprintf("DeckID: %s", z.deckID))
	printString = append(printString, fmt.Sprintf("%s", z.card))

	return strings.Join(printString, "\n")
}

func (z *Pile) AddCardsToPile(draw *Draw, cards []*Card) {

	if draw != nil && draw.Success && len(draw.Cards) != 0 {
		if len(draw.Cards) >= len(cards) {

			for _, card := range cards {
				found := false
				for i, f := range draw.Cards {
					if f.Value == card.Value && f.Suit == card.Suit {
						draw.Cards = append(draw.Cards[:i], draw.Cards[i+1:]...)
						found = true
					}
				}
				if found {
					p := &pileObject{
						deckID:      draw.DeckID,
						stillInPile: true,
						card:        card,
					}
					z.stack = append(z.stack, p)
				}
			}
		}
	} else {
	}
}

func NewPile() *Pile {
	return &Pile{
		PileID: uuid.NewV4().String(),
		stack:  make([]*pileObject, 0),
	}
}

func (z *Pile) ListCardsInPile() string {
	var printString []string
	printString = append(printString, fmt.Sprintf("PileID: %s", z.PileID))

	for _, stackObject := range z.stack {
		if !stackObject.card.drawn {
			printString = append(printString, fmt.Sprintf("%s", stackObject))
		}
	}

	return strings.Join(printString, "\n")
}

func (z *Pile) PickAmountOfCardsFromPile(amount int) []*Card {
	return nil
}

func (z *Pile) PickAllCardsFromPile() []*Card {
	return nil
}

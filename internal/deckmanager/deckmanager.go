package deckmanager

import "github.com/heindrichpaul/deckofcards"

var decks map[string]*deckofcards.Deck
var piles map[string]*deckofcards.Pile

func RequestNumberOfDecks(number int) *deckofcards.Deck {
	deck := deckofcards.NewDeck(number)
	if deck.Success {
		decks[deck.DeckID] = deck
	}
	return deck
}

func RequestNumberOfShuffledDecks(number int) *deckofcards.Deck {
	deck := RequestNumberOfDecks(number)
	if deck.Success {
		deck = deckofcards.ShuffleDeck(deck)
		decks[deck.DeckID] = deck
	}
	return deck
}

func RequestNumberOfDecksWithJockers(number int) *deckofcards.Deck {
	deck := deckofcards.NewDeckWithJockers(number)
	if deck.Success {
		decks[deck.DeckID] = deck
	}
	return deck
}

func RequestNumberOfShuffledDecksWithJockers(number int) *deckofcards.Deck {
	deck := RequestNumberOfDecksWithJockers(number)
	if deck.Success {
		deck = deckofcards.ShuffleDeck(deck)
		decks[deck.DeckID] = deck
	}
	return deck
}

func ReshuffleDeck(deck *deckofcards.Deck) *deckofcards.Deck {
	deck = deckofcards.ShuffleDeck(deck)
	if deck.Success {
		decks[deck.DeckID] = deck
	}
	return deck

}

func RequestSingleUnshuffledDeck() *deckofcards.Deck {
	deck := deckofcards.NewDeck(1)
	if deck.Success {
		decks[deck.DeckID] = deck
	}
	return deck
}

func FindDeckById(Id string) *deckofcards.Deck {
	deck, ok := decks[Id]
	if !ok {
		return nil
	}
	return deck
}

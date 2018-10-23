package deckmanager

import "github.com/heindrichpaul/deckOfCards"

var decks map[string]*deckOfCards.Deck
var piles map[string]*deckOfCards.Pile

func RequestNumberOfDecks(number int) *deckOfCards.Deck {
	deck := deckOfCards.NewDeck(number)
	if deck.Success {
		decks[deck.DeckID] = deck
	}
	return deck
}

func RequestNumberOfShuffledDecks(number int) *deckOfCards.Deck {
	deck := RequestNumberOfDecks(number)
	if deck.Success {
		deck = deckOfCards.ShuffleDeck(deck)
		decks[deck.DeckID] = deck
	}
	return deck
}

func RequestNumberOfDecksWithJockers(number int) *deckOfCards.Deck {
	deck := deckOfCards.NewDeckWithJockers(number)
	if deck.Success {
		decks[deck.DeckID] = deck
	}
	return deck
}

func RequestNumberOfShuffledDecksWithJockers(number int) *deckOfCards.Deck {
	deck := RequestNumberOfDecksWithJockers(number)
	if deck.Success {
		deck = deckOfCards.ShuffleDeck(deck)
		decks[deck.DeckID] = deck
	}
	return deck
}

func ReshuffleDeck(deck *deckOfCards.Deck) *deckOfCards.Deck {
	deck = deckOfCards.ShuffleDeck(deck)
	if deck.Success {
		decks[deck.DeckID] = deck
	}
	return deck

}

func RequestSingleUnshuffledDeck() *deckOfCards.Deck {
	deck := deckOfCards.NewDeck(1)
	if deck.Success {
		decks[deck.DeckID] = deck
	}
	return deck
}

func FindDeckById(Id string) *deckOfCards.Deck {
	deck, ok := decks[Id]
	if !ok {
		return nil
	}
	return deck
}

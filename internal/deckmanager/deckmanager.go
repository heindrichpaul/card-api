package deckmanager

import (
	"github.com/heindrichpaul/blackjack/internal/domain"
)

var decks map[string]*domain.Deck
var piles map[string]*domain.Pile

func RequestNumberOfDecks(number int) *domain.Deck {
	deck := domain.NewDeck(number)
	if deck.Success {
		decks[deck.DeckID] = deck
	}
	return deck
}

func RequestNumberOfShuffledDecks(number int) *domain.Deck {
	deck := RequestNumberOfDecks(number)
	if deck.Success {
		deck = domain.ShuffleDeck(deck)
		decks[deck.DeckID] = deck
	}
	return deck
}

func RequestNumberOfDecksWithJockers(number int) *domain.Deck {
	deck := domain.NewDeckWithJockers(number)
	if deck.Success {
		decks[deck.DeckID] = deck
	}
	return deck
}

func RequestNumberOfShuffledDecksWithJockers(number int) *domain.Deck {
	deck := RequestNumberOfDecksWithJockers(number)
	if deck.Success {
		deck = domain.ShuffleDeck(deck)
		decks[deck.DeckID] = deck
	}
	return deck
}

func ReshuffleDeck(deck *domain.Deck) *domain.Deck {
	deck = domain.ShuffleDeck(deck)
	if deck.Success {
		decks[deck.DeckID] = deck
	}
	return deck

}

func RequestSingleUnshuffledDeck() *domain.Deck {
	deck := domain.NewDeck(1)
	if deck.Success {
		decks[deck.DeckID] = deck
	}
	return deck
}

func FindDeckById(Id string) *domain.Deck {
	deck, ok := decks[Id]
	if !ok {
		return nil
	}
	return deck
}

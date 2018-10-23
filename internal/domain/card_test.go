package domain

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestNewCard(t *testing.T) {

	var suits = [...]string{"S", "D", "C", "H"}
	for _, suit := range suits {
		for i := 2; i < 10; i++ {
			cardCreatorHelper(suit, strconv.Itoa(i), t)
		}

		cardCreatorHelper(suit, "0", t)
		cardCreatorHelper(suit, "A", t)
		cardCreatorHelper(suit, "K", t)
		cardCreatorHelper(suit, "Q", t)
		cardCreatorHelper(suit, "J", t)
	}

	cardCreatorHelper("*", "*", t)
}

func cardCreatorHelper(suit, value string, t *testing.T) {
	t.Log("Now running " + fmt.Sprintf("%s%s", value, suit) + ": " + time.Now().String())
	card, err := newCard(value, suit)
	if err != nil {
		t.Logf("Failed to create card for: %s%s\n", value, suit)
		t.FailNow()
	}
	if !strings.EqualFold(card.Code, fmt.Sprintf("%s%s", value, suit)) {
		t.Logf("Failed to verify card code for: %s%s expected: %s\n", value, suit, card.Code)
		t.FailNow()
	}
	switch suit {
	case "S":
		if !strings.EqualFold(card.Suit, "SPADES") {
			t.Logf("Failed to verify card suit for: %s%s expected: SPADES\n", value, suit)
			t.FailNow()
		}
	case "D":
		if !strings.EqualFold(card.Suit, "DIAMONDS") {
			t.Logf("Failed to verify card suit for: %s%s expected: DIAMONDS\n", value, suit)
			t.FailNow()
		}
	case "C":
		if !strings.EqualFold(card.Suit, "CLUBS") {
			t.Logf("Failed to verify card suit for: %s%s expected: CLUBS\n", value, suit)
			t.FailNow()
		}
	case "H":
		if !strings.EqualFold(card.Suit, "HEARTS") {
			t.Logf("Failed to verify card suit for: %s%s expected: HEARTS\n", value, suit)
			t.FailNow()
		}
	case "*":
		if !strings.EqualFold(card.Suit, "NONE") {
			t.Logf("Failed to verify card suit for: %s%s expected: NONE\n", value, suit)
			t.FailNow()
		}
	}

	switch value {
	case "A":
		if !strings.EqualFold(card.Value, "ACE") {
			t.Logf("Failed to verify card code for: %s%s expected: ACE but received: %s\n", value, suit, card.Value)
			t.FailNow()
		}
	case "K":
		if !strings.EqualFold(card.Value, "KING") {
			t.Logf("Failed to verify card value for: %s%s expected: KING but received: %s\n", value, suit, card.Value)
			t.FailNow()
		}
	case "Q":
		if !strings.EqualFold(card.Value, "QUEEN") {
			t.Logf("Failed to verify card value for: %s%s expected: QUEEN but received: %s\n", value, suit, card.Value)
			t.FailNow()
		}
	case "J":
		if !strings.EqualFold(card.Value, "JACK") {
			t.Logf("Failed to verify card value for: %s%s expected: JACK but received: %s\n", value, suit, card.Value)
			t.FailNow()
		}
	case "0":
		if !strings.EqualFold(card.Value, "10") {
			t.Logf("Failed to verify card value for: %s%s expected: 10 but received: %s\n", value, suit, card.Value)
			t.FailNow()
		}
	case "*":
		if !strings.EqualFold(card.Value, "JOCKER") {
			t.Logf("Failed to verify card value for: %s%s expected: JOCKER but received %s\n", value, suit, card.Value)
			t.FailNow()
		}
	default:
		if !strings.EqualFold(card.Value, value) {
			t.Logf("Failed to verify card code for: %s%s expected: %s but received: %s\n", value, suit, value, card.Value)
			t.FailNow()
		}
	}

	if card.drawn {
		t.Logf("Failed to verify card drawn flag for: %s%s expected: false but received: %t\n", value, suit, card.drawn)
		t.FailNow()
	}

	if strings.Compare(card.Image, "") != 0 {
		resp, err := http.Get(card.Image)
		if err != nil {
			t.Error(err.Error())
		}
		if resp.StatusCode != 200 {
			t.Errorf("Unable to find image %s\n", card.Image)
		}
	}

	t.Log("Finished running " + fmt.Sprintf("%s%s", value, suit) + ": " + time.Now().String())
}

func TestNewCardWithInvalidSuit(t *testing.T) {
	suit := ""
	value := "0"
	expectedError := fmt.Sprintf("Card suit (%s), value (%s): invalid suit.", suit, value)
	card, err := newCard(value, suit)
	if card == nil {
		if !strings.EqualFold(err.Error(), expectedError) {
			t.Logf("expected:[%s] but received:[%s]\n", expectedError, err.Error())
			t.FailNow()
		}
	}
}

func TestNewCardWithInvalidValue(t *testing.T) {
	suit := "S"
	value := "^"
	expectedError := fmt.Sprintf("Card suit (%s), value (%s): invalid value.", suit, value)
	card, err := newCard(value, suit)
	if card == nil {
		if !strings.Contains(err.Error(), expectedError) {
			t.Logf("expected:[%s] but received:[%s]\n", expectedError, err.Error())
			t.FailNow()
		}
	}
}

func TestCardString(t *testing.T) {
	card, err := newCard("*", "*")
	if err != nil {
		t.Logf("Failed to create card: %s\n", err.Error())
		t.FailNow()
	}
	actualString := fmt.Sprintf("%s", card)
	expectedString := fmt.Sprintf("%s - %s", card.Suit, card.Value)
	if !strings.EqualFold(actualString, expectedString) {
		t.Logf("expected:[%s] but received:[%s]\n", expectedString, actualString)
	}
}

func TestDraw(t *testing.T) {
	card, err := newCard("*", "*")
	if err != nil {
		t.Logf("Failed to create card: %s\n", err.Error())
		t.FailNow()
	}

	if card.drawn {
		t.Logf("Card not properly initialized. Expected drawn property to be false after creation.\n")
		t.FailNow()
	}

	drawnCard := card.draw()
	if !drawnCard.drawn {
		t.Logf("Card not properly drawn. Expected drawn property to be true after a draw.\n")
		t.FailNow()
	}
	if !strings.EqualFold(drawnCard.Code, card.Code) {
		t.Logf("expected:[%s] but received:[%s]\n", card.Code, drawnCard.Code)
		t.FailNow()
	}
	if !strings.EqualFold(drawnCard.Image, card.Image) {
		t.Logf("expected:[%s] but received:[%s]\n", card.Image, drawnCard.Image)
		t.FailNow()
	}
	if !strings.EqualFold(drawnCard.Value, card.Value) {
		t.Logf("expected:[%s] but received:[%s]\n", card.Value, drawnCard.Value)
		t.FailNow()
	}
	if !strings.EqualFold(drawnCard.Suit, card.Suit) {
		t.Logf("expected:[%s] but received:[%s]\n", card.Suit, drawnCard.Suit)
		t.FailNow()
	}
}

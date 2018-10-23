package domain

import (
	"fmt"
	"regexp"
	"strings"
)

type Card struct {
	Image string `json:"image"`
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
	drawn bool
}

type cardError struct {
	err   string
	value string
	suit  string
}

func newCard(value, suit string) (card *Card, err error) {

	values := regexp.MustCompile(`[2-9]|0|A|K|Q|J|\*`)
	suites := regexp.MustCompile(`S|D|C|H|\*`)

	card = &Card{
		Code:  "",
		Image: "",
		Value: "",
		Suit:  "",
		drawn: false,
	}

	if !suites.MatchString(suit) {
		return nil, &cardError{"invalid suit.", value, suit}
	} else {
		switch suit {
		case "S":
			card.Suit = "SPADES"
		case "D":
			card.Suit = "DIAMONDS"
		case "C":
			card.Suit = "CLUBS"
		case "H":
			card.Suit = "HEARTS"
		default:
			card.Suit = "NONE"
		}
	}

	if !values.MatchString(value) {
		return nil, &cardError{"invalid value.", value, suit}
	} else {
		switch value {
		case "A":
			card.Value = "ACE"
		case "K":
			card.Value = "KING"
		case "Q":
			card.Value = "QUEEN"
		case "J":
			card.Value = "JACK"
		case "0":
			card.Value = "10"
		case "*":
			card.Value = "JOCKER"
		default:
			card.Value = value
		}
	}
	card.Code = fmt.Sprintf("%s%s", value, suit)
	if !strings.EqualFold("*", value) && !strings.EqualFold("*", suit) {
		card.Image = fmt.Sprintf("https://deckofcardsapi.com/static/img/%s.png", card.Code)
	} else {
		card.Image = ""
	}

	return
}

func (z *Card) String() string {
	return fmt.Sprintf("%s - %s", z.Suit, z.Value)
}

func (z *Card) draw() *Card {

	z.drawn = true
	card := &Card{
		Code:  z.Code,
		Image: z.Image,
		Value: z.Value,
		Suit:  z.Suit,
		drawn: z.drawn,
	}

	return card
}

func (z *cardError) Error() string {
	return fmt.Sprintf(`Card suit (%s), value (%s): %s`, z.suit, z.value, z.err)
}

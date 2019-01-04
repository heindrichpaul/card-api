package main

import (
	"fmt"

	"github.com/heindrichpaul/card-api/deckmanager"
)

func main() {
	deck := deckmanager.RequestNumberOfDecks(1)
	draw := deck.Draw(2)
	fmt.Println(draw)

}

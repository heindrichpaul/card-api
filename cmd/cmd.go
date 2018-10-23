package main

import (
	"fmt"

	"github.com/heindrichpaul/blackjack/internal/domain"
)

func main() {
	card, err := domain.NewCard("*", "*")
	if err != nil {
		return
	}
	fmt.Println(card)

}

package api

import (
	"github.com/gorilla/mux"
	"github.com/heindrichpaul/card-api/api/deck"
	"github.com/heindrichpaul/card-api/api/pile"
	"github.com/heindrichpaul/card-api/manager"
)

type CardDeckApi struct {
	Router  *mux.Router
	deckApi *deck.DeckAPI
	pileApi *pile.PileAPI
}

func NewAPI() *CardDeckApi {
	carddeckApi := &CardDeckApi{
		Router: mux.NewRouter(),
	}

	carddeckApi.registerAPIs()
	return carddeckApi
}

func (z *CardDeckApi) registerAPIs() {
	manager := manager.NewManager()
	z.deckApi = deck.NewDeckAPI(z.Router, manager.DeckManager)
	z.deckApi.Register()
	z.pileApi = pile.NewPileAPI(z.Router, manager.PileManager)
	z.pileApi.Register()
}

package api

import (
	"github.com/gorilla/mux"
	"github.com/heindrichpaul/card-api/api/deck"
	"github.com/heindrichpaul/card-api/api/pile"
	"github.com/heindrichpaul/card-api/manager"
)

type CardeckApi struct {
	Router  *mux.Router
	deckApi *deck.DeckAPI
	pileApi *pile.PileAPI
}

func NewAPI() *CardeckApi {
	cardeckApi := &CardeckApi{
		Router: mux.NewRouter(),
	}

	cardeckApi.registerAPIs()
	return cardeckApi
}

func (z *CardeckApi) registerAPIs() {
	manager := manager.NewManager()
	z.deckApi = deck.NewDeckAPI(z.Router, manager.DeckManager)
	z.deckApi.Register()
	z.pileApi = pile.NewPileAPI(z.Router, manager.PileManager)
	z.pileApi.Register()
}

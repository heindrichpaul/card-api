package api

import (
	"github.com/gorilla/mux"
	"github.com/heindrichpaul/card-api/deckapi"
	"github.com/heindrichpaul/card-api/manager"
	"github.com/heindrichpaul/card-api/pileapi"
)

type CardeckApi struct {
	Router  *mux.Router
	deckApi *deckapi.DeckAPI
	pileApi *pileapi.PileAPI
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
	z.deckApi = deckapi.NewDeckAPI(z.Router, manager.DeckManager)
	z.deckApi.Register()
	z.pileApi = pileapi.NewPileAPI(z.Router, manager.PileManager)
	z.pileApi.Register()
}

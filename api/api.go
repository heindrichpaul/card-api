package api

import (
	"github.com/gorilla/mux"
	"github.com/heindrichpaul/card-api/api/deck"
	"github.com/heindrichpaul/card-api/api/pile"
	"github.com/heindrichpaul/card-api/interfaces"
	"github.com/heindrichpaul/card-api/manager"
)

//CardDeckAPI is a struct that contains the router and the api's for the pile and deck managers.
type CardDeckAPI struct {
	Manager *manager.Manager
	Router  *mux.Router
	deckAPI *deck.API
	pileAPI *pile.API
}

//NewAPI returns a pointer to a new CardDeckAPI struct with a router and subroutes for the different managers.
func NewAPI() *CardDeckAPI {
	cardDeckAPI := &CardDeckAPI{
		Manager: manager.NewManager(),
		Router:  mux.NewRouter(),
	}

	cardDeckAPI.registerAPIs()
	return cardDeckAPI
}

//NewAPIWithPersistenceManager returns a pointer to a new CardDeckAPI struct with a router and subroutes for the different managers. It also uses the specified persistence manager for persistence.
func NewAPIWithPersistenceManager(persistence interfaces.PersistenceManager) *CardDeckAPI {
	cardDeckAPI := &CardDeckAPI{
		Manager: manager.NewManagerWithPersistenceManager(persistence),
		Router:  mux.NewRouter(),
	}

	cardDeckAPI.registerAPIs()
	return cardDeckAPI
}

func (z *CardDeckAPI) registerAPIs() {
	z.deckAPI = deck.NewDeckAPI(z.Router, z.Manager.DeckManager)
	z.deckAPI.Register()
	z.pileAPI = pile.NewPileAPI(z.Router, z.Manager.PileManager)
	z.pileAPI.Register()
}

package manager

import (
	"github.com/heindrichpaul/card-api/manager/deck"
	persistancemanager "github.com/heindrichpaul/card-api/manager/persistance"
	"github.com/heindrichpaul/card-api/manager/pile"
)

type Manager struct {
	DeckManager *deck.DeckManager
	PileManager *pile.PileManager
}

func NewManager() *Manager {
	p := persistancemanager.NewPersistanceManager()
	m := &Manager{
		DeckManager: deck.NewDeckManager(p),
		PileManager: pile.NewPileManager(p),
	}

	return m
}

package manager

import (
	"github.com/heindrichpaul/card-api/deckmanager"
	persistancemanager "github.com/heindrichpaul/card-api/persistanceManager"
	"github.com/heindrichpaul/card-api/pilemanager"
)

type Manager struct {
	DeckManager *deckmanager.DeckManager
	PileManager *pilemanager.PileManager
}

func NewManager() *Manager {
	p := persistancemanager.NewPersistanceManager()
	m := &Manager{
		DeckManager: deckmanager.NewDeckManager(p),
		PileManager: pilemanager.NewPileManager(p),
	}

	return m
}

package manager

import (
	"github.com/heindrichpaul/card-api/interfaces"
	"github.com/heindrichpaul/card-api/manager/deck"
	persistancemanager "github.com/heindrichpaul/card-api/manager/persistance"
	"github.com/heindrichpaul/card-api/manager/pile"
)

//Manager is a struct that wraps all managers for easy transport between objects.
type Manager struct {
	PersistanceManager interfaces.PersistanceManager
	DeckManager        *deck.Manager
	PileManager        *pile.Manager
}

//NewManager returns a pointer to a new Manager struct.
func NewManager() *Manager {
	m := &Manager{
		PersistanceManager: persistancemanager.NewMapPersistanceManager(),
	}

	m.DeckManager = deck.NewDeckManager(m.PersistanceManager)
	m.PileManager = pile.NewPileManager(m.PersistanceManager)

	return m
}

//NewManagerWithPersistanceManager returns a pointer to a new Manager struct with the specified persistance manager.
func NewManagerWithPersistanceManager(persistance interfaces.PersistanceManager) *Manager {
	m := &Manager{
		PersistanceManager: persistance,
		DeckManager:        deck.NewDeckManager(persistance),
		PileManager:        pile.NewPileManager(persistance),
	}

	return m
}

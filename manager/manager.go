package manager

import (
	"github.com/heindrichpaul/card-api/interfaces"
	"github.com/heindrichpaul/card-api/manager/deck"
	persistencemanager "github.com/heindrichpaul/card-api/manager/persistance"
	"github.com/heindrichpaul/card-api/manager/pile"
)

//Manager is a struct that wraps all managers for easy transport between objects.
type Manager struct {
	PersistenceManager interfaces.PersistenceManager
	DeckManager        *deck.Manager
	PileManager        *pile.Manager
}

//NewManager returns a pointer to a new Manager struct.
func NewManager() *Manager {
	m := &Manager{
		PersistenceManager: persistencemanager.NewMapPersistenceManager(),
	}

	m.DeckManager = deck.NewDeckManager(m.PersistenceManager)
	m.PileManager = pile.NewPileManager(m.PersistenceManager)

	return m
}

//NewManagerWithPersistenceManager returns a pointer to a new Manager struct with the specified persistance manager.
func NewManagerWithPersistenceManager(persistance interfaces.PersistenceManager) *Manager {
	m := &Manager{
		PersistenceManager: persistance,
		DeckManager:        deck.NewDeckManager(persistance),
		PileManager:        pile.NewPileManager(persistance),
	}

	return m
}

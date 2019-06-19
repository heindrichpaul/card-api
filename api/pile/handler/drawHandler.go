package handler

import (
	"net/http"

	"github.com/heindrichpaul/card-api/manager/pile"
)

type DrawHandler struct {
	pileManager *pile.Manager
}

func CreateDrawHandler(manager *pile.Manager) *DrawHandler {
	z := &DrawHandler{
		pileManager: manager,
	}
	return z
}

func (z *DrawHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

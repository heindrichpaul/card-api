package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/heindrichpaul/card-api/apierror"
	"github.com/heindrichpaul/card-api/apiutilities"
	"github.com/heindrichpaul/card-api/manager/deck"
)

type DrawDeckHandler struct {
	deckManager *deck.Manager
}

func CreateDrawDeckHandler(manager *deck.Manager) *DrawDeckHandler {
	z := &DrawDeckHandler{
		deckManager: manager,
	}
	return z
}

func (z *DrawDeckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	amount, err := strconv.Atoi(vars["amount"])
	if err != nil {
		amount = 1
	}

	if !z.deckManager.DoesDeckExist(vars["id"]) {
		e := apierror.NewAPIError(fmt.Sprintf("Could not find deck with id: %s", vars["id"]), apierror.NotFoundError)
		apiutilities.HandleResponse(w, e)
		return
	}

	draw := z.deckManager.DrawFromDeck(vars["id"], amount)
	apiutilities.HandleResponse(w, draw)
}

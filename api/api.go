package api

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/heindrichpaul/card-api/deckmanager"
)

type CardAPI struct {
	Router    *mux.Router
	dAPI      *deckAPI
	pAPI      *pileAPI
	logErrors bool
}

func NewAPI(log bool) *CardAPI {
	cardAPI := &CardAPI{
		Router:    mux.NewRouter(),
		logErrors: log,
	}

	cardAPI.registerAPIs()
	return cardAPI
}

func (z *CardAPI) registerAPIs() {
	deckManager := deckmanager.NewDeckManager()
	z.dAPI = newDeckAPI(z.Router, deckManager)
	z.dAPI.register()
	z.pAPI = newPileAPI(z.Router, deckManager)
	z.pAPI.register()
}

func handleError(w http.ResponseWriter, r *http.Request, e *apiError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	e.logRequest(r)
	eJSON, err := e.marshal()
	if err == nil {
		fmt.Fprintf(w, string(eJSON))
		return
	}
}

func getBooleanValue(v url.Values, key string) bool {
	boolean, err := strconv.ParseBool(v.Get(key))
	if err != nil {
		boolean = false
	}
	return boolean
}

func getIntWithDefaultValueAs1(v url.Values, key string) int {
	number, err := strconv.Atoi(v.Get(key))
	if err != nil {
		number = 1
	}
	return number
}

func getStringValue(v url.Values, key string) string {
	return v.Get(key)
}

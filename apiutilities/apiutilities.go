package apiutilities

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func HandleError(w http.ResponseWriter, r *http.Request, e *ApiError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	e.LogRequest(r)
	eJSON, err := e.marshal()
	if err == nil {
		fmt.Fprintf(w, string(eJSON))
		return
	}
}

func GetBooleanValue(v url.Values, key string) bool {
	boolean, err := strconv.ParseBool(v.Get(key))
	if err != nil {
		boolean = false
	}
	return boolean
}

func GetIntWithDefaultValueAs1(v url.Values, key string) int {
	number, err := strconv.Atoi(v.Get(key))
	if err != nil {
		number = 1
	}
	return number
}

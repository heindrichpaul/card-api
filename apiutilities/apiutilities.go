package apiutilities

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
)

//GetBooleanValue returns the boolean value from the key in the url.Values.
func GetBooleanValue(v url.Values, key string) bool {
	boolean, err := strconv.ParseBool(v.Get(key))
	if err != nil {
		boolean = false
	}
	return boolean
}

//GetIntWithDefaultValueOfOne returns the int value from the key in the url.Values. If it has an error it will return a default value of one.
func GetIntWithDefaultValueOfOne(v url.Values, key string) int {
	number, err := strconv.Atoi(v.Get(key))
	if err != nil {
		number = 1
	}
	return number
}

func GetIDFromRequest(r *http.Request) (id string) {
	vars := mux.Vars(r)
	id = vars["id"]
	return
}

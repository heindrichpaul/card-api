package apiutilities

import (
	"net/url"
	"strconv"
)

func GetBooleanValue(v url.Values, key string) bool {
	boolean, err := strconv.ParseBool(v.Get(key))
	if err != nil {
		boolean = false
	}
	return boolean
}

func GetIntWithDefaultValueOfOne(v url.Values, key string) int {
	number, err := strconv.Atoi(v.Get(key))
	if err != nil {
		number = 1
	}
	return number
}

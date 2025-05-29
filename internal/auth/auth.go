package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(header http.Header) (string,error) {
	val:=header.Get("Authorization")
	if val == "" {
		return "", errors.New("missing API key in Authorization header")
	}
	vals:= strings.Split(val, " ")
	if len(vals) != 2  {
		return "", errors.New("invalid API key format, expected")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("invalid API key format, expected ApiKey")
	}
	return vals[1], nil
}
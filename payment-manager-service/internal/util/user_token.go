package util

import (
	"net/http"
	"payment-manager-service/internal/delivery/http/exception"
	"strings"
)

// Returns bearer token from given header string
func ExtractBearerToken(header string) (string, error) {
	if header == "" {
		return "", exception.NewHttpError(http.StatusBadRequest, "missing authorization header")
	}

	bearerToken := strings.Split(header, " ")
	if len(bearerToken) != 2 {
		return "", exception.NewHttpError(http.StatusBadRequest, "incorrectly formatted authorization header")
	}

	return bearerToken[1], nil
}

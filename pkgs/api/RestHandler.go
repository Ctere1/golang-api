package api

import (
	"crypto/sha512"
	"crypto/subtle"
	"fmt"
	"net/http"
)

var (
	bearerPrefix = "Bearer "
	BearerToken  = "test"
)

func restApiAuth(handler func(w http.ResponseWriter, r *http.Request)) *RestAuthHandler {
	return &RestAuthHandler{next: handler}
}

type RestAuthHandler struct {
	next func(w http.ResponseWriter, r *http.Request)
}

func (h *RestAuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if isRestReqAuthorized(w, r) {
		h.next(w, r)
	} else {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Not Authorized")
	}
}

func isRestReqAuthorized(w http.ResponseWriter, r *http.Request) bool {
	var bearerToken = bearerPrefix + BearerToken
	userToken := r.Header.Get("Authorization")
	return secureCompare(bearerToken, userToken) && BearerToken != ""
}

func secureCompare(bearerToken string, userkey string) bool {
	givenSha := sha512.Sum512([]byte(bearerToken))
	actualSha := sha512.Sum512([]byte(userkey))

	return subtle.ConstantTimeCompare(givenSha[:], actualSha[:]) == 1
}

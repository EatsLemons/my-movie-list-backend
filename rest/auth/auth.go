package auth

import "net/http"

type Authenticator struct {
}

func NewAuthenticator() *Authenticator {
	return &Authenticator{}
}

func (a *Authenticator) LoginHndlr(w http.ResponseWriter, r *http.Request) {

}

func (a *Authenticator) LooutHndlr(w http.ResponseWriter, r *http.Request) {

}

func (a *Authenticator) CallBackHndlr(w http.ResponseWriter, r *http.Request) {

}

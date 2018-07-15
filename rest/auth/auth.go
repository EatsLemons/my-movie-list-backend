package auth

import (
	"math/rand"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

type Authenticator struct {
	config *oauth2.Config
}

func NewAuthenticator(cID, cSecret, rURL string) *Authenticator {
	return &Authenticator{
		config: &oauth2.Config{
			RedirectURL:  rURL,
			ClientID:     cID,
			ClientSecret: cSecret,
		},
	}
}

func (a *Authenticator) LoginHndlr(w http.ResponseWriter, r *http.Request) {
	authURL := a.config.AuthCodeURL(a.randString(15))
	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}

func (a *Authenticator) LogoutHndlr(w http.ResponseWriter, r *http.Request) {

}

func (a *Authenticator) CallBackHndlr(w http.ResponseWriter, r *http.Request) {

}

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (a *Authenticator) randString(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}

	return string(b)
}

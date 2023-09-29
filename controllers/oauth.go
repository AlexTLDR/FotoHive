package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
	"golang.org/x/oauth2"
)

type OAuth struct {
	ProviderConfigs map[string]*oauth2.Config
}

func (oa OAuth) Connect(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	provider = strings.ToLower(provider)
	conf, ok := oa.ProviderConfigs[provider]
	if !ok {
		http.Error(w, "Unknown provider - invalid OAuth2 Service", http.StatusBadRequest)
		return
	}
	state := csrf.Token(r)
	setCookie(w, "oauth_state", state)
	url := conf.AuthCodeURL(state, oauth2.SetAuthURLParam("redirect_uri", redirectURI(r, provider)))
	http.Redirect(w, r, url, http.StatusFound)
}

func redirectURI(r *http.Request, provider string) string {
	if r.Host == "localhost:3000" {
		return fmt.Sprintf("http://localhost:3000/oauth/%s/callback", provider)
	}
	return fmt.Sprintf("https://fotohive.alextldr.com/oauth/%s/callback", provider)
}

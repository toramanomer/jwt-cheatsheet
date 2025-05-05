package api

import (
	"encoding/json"
	"errors"
	"jwt-revocation/token"
	"jwt-revocation/user"
	"net/http"
)

type API struct {
	userStore  *user.UserStore
	tokenStore *token.TokenStore
}

const cookieName = "token"

func NewAPI(userStore *user.UserStore, tokenStore *token.TokenStore) *API {
	return &API{userStore, tokenStore}
}

func (api *API) Signup(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	_, err := api.userStore.CreateUser(req.Username)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrUserAlreadyExists):
			http.Error(w, "User already exists", http.StatusConflict)
			return
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func (api *API) Signin(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	user := api.userStore.GetUser(req.Username)
	if user == nil {
		http.Error(w, "Invalid credentials...", http.StatusUnauthorized)
		return
	}

	jwt, err := token.Generate(req.Username)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:   cookieName,
		Value:  jwt,
		MaxAge: 300,
	})

	w.WriteHeader(http.StatusOK)
}

func (api *API) Signout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return
	}

	claims, err := token.Verify(cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		return
	}

	api.tokenStore.Revoke(claims.ID)

	http.SetCookie(w, &http.Cookie{
		Name:   cookieName,
		Value:  "",
		MaxAge: -1,
	})
}

func (api *API) Protected(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	claims, err := token.Verify(cookie.Value)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	if api.tokenStore.IsRevoked(claims.ID) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}

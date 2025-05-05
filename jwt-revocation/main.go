package main

import (
	"jwt-revocation/api"
	"jwt-revocation/token"
	"jwt-revocation/user"
	"net/http"
	"time"
)

func main() {

	// ----- HTTP Routing
	mux := http.NewServeMux()

	userStore := user.NewUserStore()
	tokenStore := token.NewTokenStore()

	api := api.NewAPI(userStore, tokenStore)

	mux.HandleFunc("POST	/signup", api.Signup)
	mux.HandleFunc("POST	/signin", api.Signin)
	mux.HandleFunc("POST	/signout", api.Signout)
	mux.HandleFunc("GET		/protected", api.Protected)

	// ----- Start Server
	serverAddr := ":80"
	server := http.Server{
		Addr:         serverAddr,
		Handler:      mux,
		WriteTimeout: 3 * time.Second,
		ReadTimeout:  3 * time.Second,
	}

	server.ListenAndServe()
}

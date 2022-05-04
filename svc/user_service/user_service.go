package main

import (
	"log"
	"net/http"

	sgm "delinkcious/pkg/user_manager"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	store, err := sgm.NewDbUserStore("localhost", 5432, "postgres", "postgres")
	if err != nil {
		log.Fatal(err)
	}
	svc, err := sgm.NewUserManager(store)
	if err != nil {
		log.Fatal(err)
	}

	registerHandler := httptransport.NewServer(
		makeRegisterEndpoint(svc),
		decodeRegisterRequest,
		encodeResponse,
	)

	LoginHandler := httptransport.NewServer(
		makeLoginEndpoint(svc),
		decodeLoginRequest,
		encodeResponse,
	)

	LogoutHandler := httptransport.NewServer(
		makeLogoutEndpoint(svc),
		decodeLogoutRequest,
		encodeResponse,
	)

	http.Handle("/register", registerHandler)
	http.Handle("/login", LoginHandler)
	http.Handle("/logout", LogoutHandler)
	log.Fatal(http.ListenAndServe(":7070", nil))
}

package main

import (
	"cushon/handlers"
	"cushon/handlers/middleware"
	"cushon/repos"
	"cushon/services"
	"flag"
	"fmt"
	"net/http"
)

var (
	port   int
	jwtKey = []byte("mySecretKey")
)

func main() {
	// Resolve dependencies and inject them where needed
	dataRepo, err := repos.NewDataRepo()
	if err != nil {
		panic(err.Error())
	}
	jwtService, err := services.NewJwtService(jwtKey)
	if err != nil {
		panic(err.Error())
	}
	loginService, err := services.NewLoginService(&dataRepo)
	if err != nil {
		panic(err.Error())
	}
	fundsHandler := handlers.NewFundsHandler(&dataRepo)
	loginHandler := handlers.NewLoginHandler(&loginService, &jwtService)
	makeInvestmentHandler := handlers.NewMakeInvestmentHandler(&dataRepo)
	authenticationMiddleWare := middleware.NewAuthenticationMiddleWare(jwtKey)

	// Let program args define port
	flag.IntVar(&port, "port", 3000, "port to run the service on")
	flag.Parse()

	// Set up handlers and attach to routes
	mux := http.NewServeMux()
	mux.HandleFunc("/funds", authenticationMiddleWare.Handle(fundsHandler.Handle))
	mux.HandleFunc("/investments/newinvestment", authenticationMiddleWare.Handle(makeInvestmentHandler.Handle))
	mux.HandleFunc("/login", loginHandler.Handle)

	// Start listening on that address and serving responses based on handlers above
	addr := fmt.Sprintf(":%d", port)
	http.ListenAndServe(addr, mux)
}

package handlers

import (
	"cushon/consts"
	"cushon/interfaces"
	"cushon/types"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type loginHandler struct {
	loginService interfaces.LoginService
	jwtService   interfaces.JWTService
}

// NewloginHandler creates a new instance of loginHandler
func NewLoginHandler(loginService interfaces.LoginService, jwtService interfaces.JWTService) loginHandler {
	return loginHandler{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

// Login attempts to login a user with the credentials from the request body
// If successful it will set a JWT for the user containing the user ID for future requests
func (f loginHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var creds types.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID, err := f.loginService.Login(creds)
	if err != nil {
		errorMsg := fmt.Sprintf("error determining if creds: %+v is a valid login. %s", creds, err.Error())
		http.Error(w, errorMsg, http.StatusUnauthorized)
		return
	}

	token, err := f.jwtService.GenerateToken(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(consts.JWTExpiryDuration),
	})
}

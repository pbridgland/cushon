package handlers

import (
	"cushon/interfaces"
	"encoding/json"
	"fmt"
	"net/http"
)

type fundsHandler struct {
	dataRepo interfaces.DataRepo
}

// NewFundsHandler creates a new instance of fundsHandler
func NewFundsHandler(dataRepo interfaces.DataRepo) fundsHandler {
	return fundsHandler{
		dataRepo: dataRepo,
	}
}

// Handle gets the funds from the DB and serves these as a json response
func (f fundsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	funds, err := f.dataRepo.Funds()
	if err != nil {
		errorMsg := fmt.Sprintf("error getting funds. %s", err.Error())
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(funds)
	if err != nil {
		errorMsg := fmt.Sprintf("error marshalling funds. %s", err.Error())
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

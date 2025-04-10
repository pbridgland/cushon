package handlers

import (
	"cushon/interfaces"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/context"
)

type makeInvestmentRequestBody struct {
	FundID int
	Amount int
}

type makeInvestmentHandler struct {
	dataRepo interfaces.DataRepo
}

// NewMakeInvestmentHandler creates a new instance of makeInvestmentHandler
func NewMakeInvestmentHandler(dataRepo interfaces.DataRepo) makeInvestmentHandler {
	return makeInvestmentHandler{
		dataRepo: dataRepo,
	}
}

// Handle attempts to get a userID from the request context, and unmarshal a funID and amount from the request body
// if all that information is present it calls on the DB to make that investment
func (f makeInvestmentHandler) Handle(w http.ResponseWriter, r *http.Request) {
	untypedUserID := context.Get(r, "id")
	userID, ok := untypedUserID.(int)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var makeInvestmentBody makeInvestmentRequestBody
	err := json.NewDecoder(r.Body).Decode(&makeInvestmentBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = f.dataRepo.MakeInvestment(userID, makeInvestmentBody.FundID, makeInvestmentBody.Amount)
	if err != nil {
		errorMsg := fmt.Sprintf("error making investment for user %d fund %d for amount %d. %s", userID, makeInvestmentBody.FundID, makeInvestmentBody.Amount, err.Error())
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

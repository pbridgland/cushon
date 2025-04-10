package handlers

import (
	"bytes"
	"cushon/mocks"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/context"
)

func TestMakeInvestmentHandler(t *testing.T) {
	mockRepo := &mocks.DataRepo{}
	handler := NewMakeInvestmentHandler(mockRepo)

	tests := []struct {
		name           string
		userID         int // leave as 0 to have unset
		body           interface{}
		mockSetup      func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:   "Missing User ID in Context",
			userID: 0,
			body: makeInvestmentRequestBody{
				FundID: 1,
				Amount: 100,
			},
			mockSetup:      func() {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:   "Invalid JSON Body",
			userID: 42,
			body:   "invalid-json",
			mockSetup: func() {
				// No setup needed
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:   "DataRepo Error",
			userID: 42,
			body: makeInvestmentRequestBody{
				FundID: 10,
				Amount: 200,
			},
			mockSetup: func() {
				mockRepo.NextMakeInvestmentErr = errors.New("db failure")
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "error making investment for user 42 fund 10 for amount 200. db failure\n",
		},
		{
			name:   "Successful Investment",
			userID: 42,
			body: makeInvestmentRequestBody{
				FundID: 5,
				Amount: 300,
			},
			mockSetup: func() {
				mockRepo.NextMakeInvestmentErr = nil
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			var reqBody []byte
			switch v := tt.body.(type) {
			case string:
				reqBody = []byte(v)
			default:
				var err error
				reqBody, err = json.Marshal(v)
				if err != nil {
					t.Fatalf("error marshalling request body: %v", err)
				}
			}

			req := httptest.NewRequest(http.MethodPost, "/invest", bytes.NewBuffer(reqBody))
			w := httptest.NewRecorder()

			// set userID in request context
			if tt.userID != 0 {
				context.Set(req, "id", tt.userID)
			}

			handler.Handle(w, req)

			res := w.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			if tt.expectedBody != "" {
				body := w.Body.String()
				if body != tt.expectedBody {
					t.Errorf("expected body: %s, got: %s", tt.expectedBody, body)
				}
			}
		})
	}
}

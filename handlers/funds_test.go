package handlers

import (
	"cushon/mocks"
	"cushon/types"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleFundsHandler(t *testing.T) {
	mockRepo := mocks.DataRepo{}
	u := NewFundsHandler(&mockRepo)
	tests := []struct {
		name           string
		urlPath        string
		mockSetup      func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:    "Internal server error",
			urlPath: "/funds",
			mockSetup: func() {
				mockRepo.NextFundsErr = errors.New("test error")
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "error getting funds. test error\n",
		},
		{
			name:    "Successful Funds Req",
			urlPath: "/funds",
			mockSetup: func() {
				mockRepo.NextFundsErr = nil
				mockRepo.NextFundsResult = types.Funds{
					types.Fund{
						ID:   123,
						Name: "abc",
					},
				}
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `[{"id":123,"name":"abc"}]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			req := httptest.NewRequest(http.MethodGet, tt.urlPath, nil)
			w := httptest.NewRecorder()

			u.Handle(w, req)

			res := w.Result()
			defer res.Body.Close()

			if tt.expectedStatus != res.StatusCode {
				t.Errorf("expected status code to be %d but got %d", tt.expectedStatus, res.StatusCode)
			}

			bodyBytes := w.Body.Bytes()
			bodyString := string(bodyBytes)
			if tt.expectedBody != bodyString {
				t.Errorf("expected body to be %s but got %s", tt.expectedBody, bodyString)
			}
		})
	}
}

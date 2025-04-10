package handlers

import (
	"bytes"
	"cushon/mocks"
	"cushon/types"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginHandler(t *testing.T) {
	mockLoginService := &mocks.LoginService{}
	mockJWTService := &mocks.JWTService{}

	handler := NewLoginHandler(mockLoginService, mockJWTService)

	tests := []struct {
		name           string
		body           interface{}
		mockSetup      func()
		expectedStatus int
		expectCookie   bool
	}{
		{
			name: "Invalid JSON Body",
			body: "invalid-json",
			mockSetup: func() {
				// No setup needed
			},
			expectedStatus: http.StatusBadRequest,
			expectCookie:   false,
		},
		{
			name: "Invalid Login Credentials",
			body: types.Credentials{
				Username: "wrong",
				Password: "incorrect",
			},
			mockSetup: func() {
				mockLoginService.NextLoginResult = 0
				mockLoginService.NextLoginErr = errors.New("invalid credentials")
			},
			expectedStatus: http.StatusUnauthorized,
			expectCookie:   false,
		},
		{
			name: "JWT Generation Failure",
			body: types.Credentials{
				Username: "user",
				Password: "pass",
			},
			mockSetup: func() {
				mockLoginService.NextLoginResult = 42
				mockLoginService.NextLoginErr = nil
				mockJWTService.NextGenerateTokenResult = ""
				mockJWTService.NextGenerateTokenErr = errors.New("token error")
			},
			expectedStatus: http.StatusInternalServerError,
			expectCookie:   false,
		},
		{
			name: "Successful Login",
			body: types.Credentials{
				Username: "user",
				Password: "pass",
			},
			mockSetup: func() {
				mockLoginService.NextLoginResult = 42
				mockLoginService.NextLoginErr = nil
				mockJWTService.NextGenerateTokenResult = "valid.jwt.token"
				mockJWTService.NextGenerateTokenErr = nil
			},
			expectedStatus: http.StatusOK,
			expectCookie:   true,
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

			req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBody))
			w := httptest.NewRecorder()

			handler.Handle(w, req)

			res := w.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			cookieFound := false
			for _, cookie := range res.Cookies() {
				if cookie.Name == "token" {
					cookieFound = true
				}
			}
			if tt.expectCookie != cookieFound {
				t.Errorf("expected cookie presence to be %v, but got %v", tt.expectCookie, cookieFound)
			}
		})
	}
}

package middleware

import (
	"cushon/types"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/context"
)

func generateToken(t *testing.T, jwtKey []byte, claims *types.Claims, valid bool) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if !valid {
		token = jwt.New(jwt.SigningMethodNone)
	}
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		t.Fatalf("failed to sign token: %v", err)
	}
	return tokenStr
}

func TestAuthenticationMiddleware(t *testing.T) {
	jwtKey := []byte("test-secret")
	middleware := NewAuthenticationMiddleWare(jwtKey)

	tests := []struct {
		name           string
		cookie         *http.Cookie
		expectedStatus int
		expectContext  bool
	}{
		{
			name:           "No Cookie",
			cookie:         nil,
			expectedStatus: http.StatusUnauthorized,
			expectContext:  false,
		},
		{
			name: "Malformed Token",
			cookie: &http.Cookie{
				Name:  "token",
				Value: "invalid.token.here",
			},
			expectedStatus: http.StatusUnauthorized,
			expectContext:  false,
		},
		{
			name: "Valid Token",
			cookie: func() *http.Cookie {
				claims := &types.Claims{
					ID: 123,
					RegisteredClaims: jwt.RegisteredClaims{
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
					},
				}
				return &http.Cookie{
					Name:  "token",
					Value: generateToken(t, jwtKey, claims, true),
				}
			}(),
			expectedStatus: http.StatusOK,
			expectContext:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if tt.cookie != nil {
				req.AddCookie(tt.cookie)
			}

			w := httptest.NewRecorder()

			var handlerCalled bool
			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				handlerCalled = true
				if tt.expectContext {
					val := context.Get(r, "id")
					if val != 123 {
						t.Errorf("expected context value 123 but got %v", val)
					}
				}
				w.WriteHeader(http.StatusOK)
			})

			authHandler := middleware.Handle(nextHandler)
			authHandler.ServeHTTP(w, req)

			res := w.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			if tt.expectedStatus == http.StatusOK && !handlerCalled {
				t.Errorf("expected handler to be called but it wasn't")
			}
			if tt.expectedStatus != http.StatusOK && handlerCalled {
				t.Errorf("expected handler not to be called but it was")
			}
		})
	}
}

package middleware

import (
	"cushon/types"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/context"
)

type authenticationMiddleWare struct {
	jwtKey []byte
}

// NewAuthenticationMiddleWare creates a new instance of authenticationMiddleWare
func NewAuthenticationMiddleWare(jwtKey []byte) authenticationMiddleWare {
	return authenticationMiddleWare{
		jwtKey: jwtKey,
	}
}

// Handle checks if there is a cookie named "token" present,
// if so it parses this as a JWT and sets the userID stored in it in the request context
// if any of this fails it returns a 401 (Unauthorised) without proceeding onto the next handlerfunc
// if all of this succeeds it proceeds to the next handlerfunc
func (a authenticationMiddleWare) Handle(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tokenStr := c.Value
		claims := &types.Claims{}

		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return a.jwtKey, nil
		})

		if err != nil || !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if claims, ok := tkn.Claims.(*types.Claims); ok && tkn.Valid {
			context.Set(r, "id", claims.ID)
		}

		next.ServeHTTP(w, r)
	})
}

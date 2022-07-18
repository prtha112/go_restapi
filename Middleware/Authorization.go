package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Env struct {
	Signature string
}

func AuthorizationMiddleware(env Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		// signature = os.Getenv("JWT_SIGNATURE")
		s := c.Request.Header.Get("Authorization")

		token := strings.TrimPrefix(s, "Bearer ")

		if err := validateToken(token, env.Signature); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

func validateToken(token string, signature string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(signature), nil
	})

	return err
}

package handler

import (
	"jakpat-test-2/pkg/usecase"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	usecases usecase.Usecase
}

func NewAuthMiddleware(usecases usecase.Usecase) gin.HandlerFunc {
	return (&AuthMiddleware{
		usecases: usecases,
	}).loginCheck
}

func (m *AuthMiddleware) loginCheck(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := m.usecases.ParseToken(c.Request.Context(), headerParts[1])
	if err != nil {
		status := http.StatusInternalServerError
		if err == usecase.ErrInvalidAccessToken {
			status = http.StatusUnauthorized
		}

		c.AbortWithStatus(status)
		return
	}

	c.Set(usecase.CtxUserKey, user)
}

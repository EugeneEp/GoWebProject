package middleware

import (
	"GoWebProject/internal/model/errors"
	"GoWebProject/internal/usecase"
	"GoWebProject/pkg/crypt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authUseCase *usecase.AuthUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		items := strings.Split(header, " ")
		if len(items) < 2 {
			c.JSON(http.StatusInternalServerError, errors.ErrInvalidAuthData)
			authUseCase.Logger.Error(errors.ErrInvalidAuthData.Error())
			c.Abort()
			return
		}

		if _, err := crypt.VerifyToken(items[1], authUseCase.Config.GetString("app.api-key")); err != nil {
			c.JSON(http.StatusInternalServerError, errors.ErrInvalidAuthData)
			authUseCase.Logger.Error(err.Error())
			c.Abort()
			return
		}

		c.Next()
	}
}

package http

import (
	"GoWebProject/internal/model"
	"GoWebProject/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	UseCase *usecase.AuthUseCase
}

func NewAuthController(usecase *usecase.AuthUseCase) *AuthController {
	return &AuthController{
		UseCase: usecase,
	}
}

func (c *AuthController) Auth(ctx *gin.Context) {
	req := new(model.AuthRequest)

	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	token, err := c.UseCase.Auth(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err)
		return
	}

	ctx.JSON(http.StatusOK, token)
}

package http

import (
	"GoWebProject/internal/model"
	"GoWebProject/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UseCase *usecase.UserUseCase
}

func NewUserController(usecase *usecase.UserUseCase) *UserController {
	return &UserController{
		UseCase: usecase,
	}
}

func (c *UserController) Get(ctx *gin.Context) {
	user, err := c.UseCase.GetById(ctx, model.UserRequest{ID: ctx.Param("id")})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) GetAll(ctx *gin.Context) {
	users, err := c.UseCase.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) Create(ctx *gin.Context) {
	req := new(model.CreateUserRequest)

	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := c.UseCase.Create(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) Update(ctx *gin.Context) {
	req := new(model.UpdateUserRequest)
	req.ID = ctx.Param("id")

	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := c.UseCase.Update(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) Delete(ctx *gin.Context) {
	if err := c.UseCase.Delete(ctx, &model.UserRequest{ID: ctx.Param("id")}); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
}

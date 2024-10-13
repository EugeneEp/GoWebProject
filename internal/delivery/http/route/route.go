package route

import (
	"GoWebProject/internal/delivery/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	AuthController *http.AuthController
	UserController *http.UserController
	AuthMiddleware gin.HandlerFunc
	Server         *gin.Engine
}

func (r *Router) Setup() {
	auth := r.Server.Group("/api/v1/auth")
	auth.POST("/", r.AuthController.Auth)

	users := r.Server.Group("/api/v1/users")
	users.Use(r.AuthMiddleware)
	users.GET("", r.UserController.GetAll)
	users.GET("/:id", r.UserController.Get)
	users.POST("", r.UserController.Create)
	users.PATCH("/:id", r.UserController.Update)
	users.DELETE("/:id", r.UserController.Delete)
}

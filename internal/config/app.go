package config

import (
	"GoWebProject/internal/delivery/http"
	"GoWebProject/internal/delivery/http/middleware"
	"GoWebProject/internal/delivery/http/route"
	"GoWebProject/internal/repository"
	"GoWebProject/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppConfig struct {
	DB        *gorm.DB
	Server    *gin.Engine
	Config    *viper.Viper
	Logger    *zap.Logger
	Validator *validator.Validate
}

func App(app *AppConfig) {
	userRepository := repository.NewUserRepository()

	userUseCase := usecase.NewUserUseCase(app.DB, userRepository, app.Logger, app.Validator)
	authUseCase := usecase.NewAuthUseCase(app.Config, app.Logger, app.Validator)

	userController := http.NewUserController(userUseCase)
	authController := http.NewAuthController(authUseCase)

	router := route.Router{
		AuthController: authController,
		UserController: userController,
		AuthMiddleware: middleware.AuthMiddleware(authUseCase),
		Server:         app.Server,
	}

	router.Setup()
}

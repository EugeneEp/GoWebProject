package usecase

import (
	"GoWebProject/internal/model"
	"GoWebProject/internal/model/errors"
	"GoWebProject/pkg/crypt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type AuthUseCase struct {
	Config    *viper.Viper
	Logger    *zap.Logger
	Validator *validator.Validate
}

func NewAuthUseCase(config *viper.Viper, log *zap.Logger, val *validator.Validate) *AuthUseCase {
	return &AuthUseCase{Config: config, Logger: log, Validator: val}
}

func (c *AuthUseCase) Auth(req *model.AuthRequest) (*model.AuthResponse, error) {
	if err := c.Validator.Struct(req); err != nil {
		c.Logger.Error(errors.ErrInvalidAuthData.Error())
		return nil, errors.ErrInvalidAuthData
	}

	if req.Password == "admin" && req.Username == "admin" {
		token, err := crypt.NewJwtToken([]byte(c.Config.GetString("app.api-key")), req.Username, c.Config.GetInt("token.lifetime"))
		if err != nil {
			c.Logger.Error(err.Error())
			return nil, err
		}

		return &model.AuthResponse{
			Token: token,
		}, nil
	}
	c.Logger.Error(errors.ErrInvalidAuthData.Error())
	return nil, errors.ErrInvalidAuthData
}

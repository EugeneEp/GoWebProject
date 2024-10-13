package usecase

import (
	"GoWebProject/internal/entity"
	"GoWebProject/internal/model"
	"GoWebProject/internal/model/errors"
	"GoWebProject/internal/repository"
	"GoWebProject/pkg/crypt"
	"context"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserUseCase struct {
	DB         *gorm.DB
	Repository *repository.UserRepository
	Logger     *zap.Logger
	Validator  *validator.Validate
}

func NewUserUseCase(db *gorm.DB, repository *repository.UserRepository, log *zap.Logger, val *validator.Validate) *UserUseCase {
	return &UserUseCase{DB: db, Repository: repository, Logger: log, Validator: val}
}

func (c *UserUseCase) GetAll(ctx context.Context) (*[]entity.UserEntity, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	users := &[]entity.UserEntity{}

	if err := c.Repository.GetAll(tx, users); err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	return users, nil
}

func (c *UserUseCase) GetById(ctx context.Context, req model.UserRequest) (*entity.UserEntity, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	user := &entity.UserEntity{}

	if err := c.Repository.GetById(tx, user, req.ID); err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	return user, nil
}

func (c *UserUseCase) Create(ctx context.Context, req *model.CreateUserRequest) (*entity.UserEntity, error) {
	if err := c.Validator.Struct(req); err != nil {
		c.Logger.Error(errors.ErrInvalidUserData.Error())
		return nil, errors.ErrInvalidUserData
	}

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	user := entity.UserEntityInit(req.Email, req.Password, req.Username)

	if err := c.Repository.Create(tx, user); err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	return user, nil
}

func (c *UserUseCase) Update(ctx context.Context, req *model.UpdateUserRequest) (*entity.UserEntity, error) {
	if err := c.Validator.Struct(req); err != nil {
		c.Logger.Error(errors.ErrInvalidUserData.Error())
		return nil, errors.ErrInvalidUserData
	}

	if req.Password == nil && req.Username == nil {
		return nil, nil
	}

	user, err := c.GetById(ctx, model.UserRequest{ID: req.ID})
	if err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	if req.Password != nil {
		user.Password = crypt.SHA256(*req.Password)
	}

	if req.Username != nil {
		user.Username = *req.Username
	}

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Repository.Update(tx, user); err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Logger.Error(err.Error())
		return nil, err
	}

	return user, nil
}

func (c *UserUseCase) Delete(ctx context.Context, req *model.UserRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Repository.Delete(tx, &entity.UserEntity{ID: req.ID}); err != nil {
		c.Logger.Error(err.Error())
		return err
	}

	return tx.Commit().Error
}

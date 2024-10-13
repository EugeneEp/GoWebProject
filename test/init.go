package test

import (
	"GoWebProject/internal/config"
	"GoWebProject/internal/repository"
	"GoWebProject/internal/usecase"
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db          *gorm.DB
	viperConfig *viper.Viper
	validate    *validator.Validate
	log         *zap.Logger
	mock        sqlmock.Sqlmock
	userCase    *usecase.UserUseCase
	authCase    *usecase.AuthUseCase
)

func init() {
	viperConfig = viper.New()

	viperConfig.SetConfigName("config")
	viperConfig.SetConfigType("json")
	viperConfig.AddConfigPath("./../")

	viperConfig.AutomaticEnv()
	if err := viperConfig.ReadInConfig(); err != nil {
		panic(err.Error())
	}

	conf := zap.Config{
		Level:    zap.NewAtomicLevelAt(zap.DebugLevel),
		Encoding: "json",
	}

	log, _ = conf.Build()
	validate = config.NewValidator()

	var mockDb *sql.DB

	mockDb, mock, _ = sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ = gorm.Open(dialector, &gorm.Config{})

	userCase = usecase.NewUserUseCase(db, repository.NewUserRepository(), log, validate)
	authCase = usecase.NewAuthUseCase(viperConfig, log, validate)
}

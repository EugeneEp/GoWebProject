package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGorm(cfg *viper.Viper, log *zap.Logger) *gorm.DB {
	username := cfg.GetString("DB_USER")
	password := cfg.GetString("DB_PASSWORD")
	host := cfg.GetString("DB_IP")
	port := cfg.GetInt("DB_PORT")
	database := cfg.GetString("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", host, username, password, database, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	conn, err := db.DB()
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	if err = init_tables(conn); err != nil {
		log.Error(err.Error())
		panic(err)
	}

	return db
}

func init_tables(conn *sql.DB) error {
	c, err := os.ReadFile("./sql/table.sql")
	if err != nil {
		return err
	}
	sql := string(c)
	_, err = conn.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

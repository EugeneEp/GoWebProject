package entity

import (
	"GoWebProject/pkg/crypt"
	"time"

	"github.com/google/uuid"
)

type UserEntity struct {
	ID        string    `gorm:"column:id;unique;primaryKey;type:uuid;default:uuid_generate_v4()"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password_hash"`
	Timestamp time.Time `gorm:"column:timestamp"`
	Username  string    `gorm:"column:username"`
}

func UserEntityInit(email, password, username string) *UserEntity {
	return &UserEntity{
		ID:        uuid.New().String(),
		Email:     email,
		Password:  crypt.SHA256(password),
		Timestamp: time.Now(),
		Username:  username,
	}
}

func (e *UserEntity) TableName() string {
	return "users"
}

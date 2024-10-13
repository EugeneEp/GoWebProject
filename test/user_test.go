package test

import (
	"GoWebProject/internal/model"
	"GoWebProject/pkg/crypt"
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var user_id = "123"
var current_time = time.Now()

func TestUserCreate(t *testing.T) {
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(user_id))
	mock.ExpectCommit()

	res, err := userCase.Create(context.Background(), &model.CreateUserRequest{
		Email:    "123",
		Password: "123",
		Username: "123",
	})

	assert.Equal(t, user_id, res.ID)
	assert.Nil(t, err)
}

func TestUserUpdate(t *testing.T) {
	newVal := "1234"

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "timestamp", "username"}).AddRow("123", "123", "123", current_time, "123"))
	mock.ExpectCommit()
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE").WithArgs("123", crypt.SHA256(newVal), current_time, "1234", "123").WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	res, err := userCase.Update(context.Background(), &model.UpdateUserRequest{
		UserRequest: model.UserRequest{ID: user_id},
		Username:    &newVal,
		Password:    &newVal,
	})

	assert.Equal(t, newVal, res.Username)
	assert.Equal(t, crypt.SHA256(newVal), res.Password)
	assert.Nil(t, err)
}

func TestUserGetById(t *testing.T) {
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "timestamp", "username"}).AddRow("123", "123", "123", current_time, "123"))
	mock.ExpectCommit()

	res, err := userCase.GetById(context.Background(), model.UserRequest{ID: user_id})

	assert.Equal(t, user_id, res.ID)
	assert.Nil(t, err)
}

func TestUserGetAll(t *testing.T) {
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "timestamp", "username"}).AddRow("123", "123", "123", current_time, "123"))
	mock.ExpectCommit()

	res, err := userCase.GetAll(context.Background())

	assert.Equal(t, 1, len(*res))
	assert.Nil(t, err)
}

func TestUserDelete(t *testing.T) {
	mock.ExpectBegin()
	mock.ExpectExec("DELETE").WithArgs(user_id).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := userCase.Delete(context.Background(), &model.UserRequest{ID: user_id})

	assert.Nil(t, err)
}

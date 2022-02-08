package services

import (
	"github/aryabaguss/bookstore_users-api/domain/users"
	"github/aryabaguss/bookstore_users-api/utils/errors"
	"net/http"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil

	return &user, &errors.RestErr{
		Status: http.StatusInternalServerError,
	}
}

package internal

import (
	"errors"
	"fmt"

	"github.com/mca312/jackbox/server/internal/data"
)

type User struct {
	Email string
	Id    int64
}

func Register(u User) error {
	dataUser := data.GetUserByEmail(u.Email)
	if dataUser != nil {
		return errors.New(fmt.Sprintf("User already exists during registration: %s", u.Email))
	}
	dataUser = &data.DataUser{
		Email: u.Email,
	}
	return data.AddUser(*dataUser)
}

func Login(u User) (User, error) {
	dataUser := data.GetUserByEmail(u.Email)
	if dataUser == nil {
		return User{}, errors.New(fmt.Sprintf("User does not exist during login: %s", u.Email))
	}

	return User{Email: dataUser.Email, Id: dataUser.Id}, nil
}

func GetUser(id int64) (User, error) {
	dataUser := data.GetUserById(id)
	if dataUser == nil {
		return User{}, errors.New(fmt.Sprintf("User does not exist during get user: %d", id))
	}

	return User{Email: dataUser.Email, Id: dataUser.Id}, nil
}

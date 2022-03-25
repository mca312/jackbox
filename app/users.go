package app

import (
	"errors"
	"fmt"

	"github.com/mca312/jackbox/server/app/data"
)

type UserServiceApi interface {
	Register(User) error
	Login(User) (User, error)
	GetUser(int64) (User, error)
}

type UserService struct {
	db data.UserDataBase
}

type User struct {
	Email string
	Id    int64
}

// NewUserService sets up the dependencies we'll need. In our case we need our in memory DB
func NewUserService() *UserService {
	return &UserService{
		db: data.NewMemoryUserDB(),
	}
}

// Register a new user to our DB
func (service *UserService) Register(u User) error {
	dataUser := service.db.GetUserByEmail(u.Email)
	if dataUser != nil {
		return errors.New(fmt.Sprintf("User already exists during registration: %s", u.Email))
	}
	dataUser = &data.DataUser{
		Email: u.Email,
	}
	return service.db.AddUser(*dataUser)
}

// Login a new user and return the User object
func (service *UserService) Login(u User) (User, error) {
	dataUser := service.db.GetUserByEmail(u.Email)
	if dataUser == nil {
		return User{}, errors.New(fmt.Sprintf("User does not exist during login: %s", u.Email))
	}

	return User{Email: dataUser.Email, Id: dataUser.Id}, nil
}

// GetUser by ID
func (service *UserService) GetUser(id int64) (User, error) {
	dataUser := service.db.GetUserById(id)
	if dataUser == nil {
		return User{}, errors.New(fmt.Sprintf("User does not exist during get user: %d", id))
	}

	return User{Email: dataUser.Email, Id: dataUser.Id}, nil
}

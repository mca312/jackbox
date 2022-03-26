package app

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mca312/jackbox/server/app/data"
)

type UserServiceApi interface {
	Register(User) (*User, error)
	Login(User) (User, error)
	GetUser(int64) (User, error)
}

type UserService struct {
	db data.UserDataBase
}

type User struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
	Name  string `json:"name"`
}

// NewUserService sets up the dependencies we'll need. In our case we need our in memory DB
func NewUserService() *UserService {
	return &UserService{
		db: data.NewMemoryUserDB(),
	}
}

// Register a new user to our DB
func (service *UserService) Register(u User) (*User, error) {

	// check email exists and has one @ symbol
	if !IsValidEmail(u.Email) {
		return nil, errors.New("Invalid Email")
	}
	// check name is not empty
	if len(strings.TrimSpace(u.Name)) == 0 {
		return nil, errors.New("Invalid Name")
	}

	dataUser := service.db.GetUserByEmail(u.Email)
	if dataUser != nil {
		return nil, errors.New(fmt.Sprintf("User already exists during registration: %s", u.Email))
	}
	dataUser = &data.DataUser{
		Email: u.Email,
		Name:  u.Name,
	}

	user, err := service.db.AddUser(*dataUser)
	if err != nil {
		return nil, err
	}

	return &User{Email: user.Email, Id: user.Id, Name: user.Name}, nil
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

	return User{Email: dataUser.Email, Id: dataUser.Id, Name: dataUser.Name}, nil
}

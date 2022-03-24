package data

import (
	"errors"
	"fmt"
)

type UserDataBase interface {
	AddUser() error
	GetUserByEmail() *DataUser
	GetUserById() *DataUser
}

type DataUser struct {
	Email string
	Id    int64
}

var testUsers = map[int64]*DataUser{}
var id int64 = 0

// GetUserByEmail simulates DB call
func GetUserByEmail(email string) *DataUser {
	var user *DataUser
	for _, v := range testUsers {
		if v.Email == email {
			user = v
		}
	}

	return user
}

// GetUserById simulates DB call
func GetUserById(id int64) *DataUser {
	return testUsers[id]
}

func AddUser(user DataUser) error {
	exists := GetUserByEmail(user.Email)
	if exists != nil {
		return errors.New(fmt.Sprintf("User exists during Add User: %s", user.Email))
	}

	fmt.Printf("Added user: %s\n", user.Email)
	id = id + 1
	user.Id = id
	testUsers[id] = &user
	return nil
}

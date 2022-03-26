package data

import (
	"errors"
	"fmt"
)

type UserDataBase interface {
	AddUser(DataUser) (*DataUser, error)
	GetUserByEmail(string) *DataUser
	GetUserById(int64) *DataUser
}

type memoryUserDB struct {
	id        int64
	testUsers map[int64]*DataUser
}

func NewMemoryUserDB() *memoryUserDB {
	return &memoryUserDB{
		id:        0,
		testUsers: map[int64]*DataUser{},
	}
}

type DataUser struct {
	Email string
	Id    int64
	Name  string
}

// GetUserByEmail simulates DB call
func (m *memoryUserDB) GetUserByEmail(email string) *DataUser {
	var user *DataUser
	for _, v := range m.testUsers {
		if v.Email == email {
			user = v
		}
	}

	return user
}

// GetUserById simulates DB call
func (m *memoryUserDB) GetUserById(id int64) *DataUser {
	return m.testUsers[id]
}

func (m *memoryUserDB) AddUser(user DataUser) (*DataUser, error) {
	exists := m.GetUserByEmail(user.Email)
	if exists != nil {
		return nil, errors.New(fmt.Sprintf("User exists during Add User: %s", user.Email))
	}

	fmt.Printf("Added user: %s\n", user.Email)
	// increment our Id
	m.id = m.id + 1
	user.Id = m.id
	m.testUsers[m.id] = &user
	return &user, nil
}

package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsers_Register(t *testing.T) {
	t.Run("user registers successfully", func(t *testing.T) {
		validUser := User{Email: "valid@email", Name: "validName"}
		service := NewUserService()

		user, err := service.Register(validUser)

		assert.Nil(t, err)
		assert.Equal(t, user.Email, validUser.Email)
	})

	t.Run("returns error is user already exists", func(t *testing.T) {
		validUser := User{Email: "valid@email", Name: "validName"}
		service := NewUserService()

		// verify user adds correctly
		_, err := service.Register(validUser)
		assert.Nil(t, err)

		// verify adding user with same email returns an error
		_, err = service.Register(validUser)
		assert.NotNil(t, err)
	})
}

func TestUsers_Login(t *testing.T) {
	t.Run("user login successfully", func(t *testing.T) {
		validUser := User{Email: "valid@email", Name: "validName"}
		service := NewUserService()

		// add a user
		_, err := service.Register(validUser)
		assert.Nil(t, err)

		// login with user
		user, err := service.Login(validUser)
		assert.Equal(t, validUser.Email, user.Email)
		assert.Equal(t, int64(1), user.Id)
		assert.Nil(t, err)
	})

	t.Run("returns error when user does not exist", func(t *testing.T) {
		validUser := User{Email: "valid@email", Name: "validName"}
		service := NewUserService()

		// login with user we havn't registered
		_, err := service.Login(validUser)
		assert.NotNil(t, err)
	})
}

func TestUsers_GetUser(t *testing.T) {
	t.Run("user returns successfully", func(t *testing.T) {
		validUser := User{Email: "valid@email", Name: "validName"}
		service := NewUserService()

		// add a user
		_, err := service.Register(validUser)
		assert.Nil(t, err)

		// get user we just registered
		user, err := service.GetUser(1)
		assert.Equal(t, validUser.Email, user.Email)
		assert.Equal(t, int64(1), user.Id)
		assert.Nil(t, err)
	})

	t.Run("returns error when user does not exist", func(t *testing.T) {
		service := NewUserService()

		// get user we haven't registered
		_, err := service.GetUser(1)
		assert.NotNil(t, err)
	})
}

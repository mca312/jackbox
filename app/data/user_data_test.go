package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserData_GetUserByEmail(t *testing.T) {
	t.Run("user returns successfully when exists", func(t *testing.T) {
		validEmail := "test@user"

		db := NewMemoryUserDB()
		db.AddUser(DataUser{Email: validEmail})

		user := db.GetUserByEmail(validEmail)

		assert.Equal(t, validEmail, user.Email)
	})

	t.Run("user returns nil when doesn't exist", func(t *testing.T) {
		validEmail := "test@user"

		db := NewMemoryUserDB()
		db.AddUser(DataUser{Email: validEmail})

		user := db.GetUserByEmail("user@not.exists")

		assert.Nil(t, user)
	})
}

func TestUserData_GetUserById(t *testing.T) {
	t.Run("user returns successfully when exists", func(t *testing.T) {
		validEmail := "test@user"

		db := NewMemoryUserDB()
		db.AddUser(DataUser{Email: validEmail})

		// get user by email first since we don't know the Id
		userForId := db.GetUserByEmail(validEmail)
		user := db.GetUserById(userForId.Id)

		assert.Equal(t, validEmail, user.Email)
		assert.Equal(t, userForId.Id, user.Id)
	})

	t.Run("user returns nil when doesn't exist", func(t *testing.T) {
		validEmail := "test@user"

		db := NewMemoryUserDB()
		db.AddUser(DataUser{Email: validEmail})

		// we know 99999 doesn't exist since we added only 1 user
		user := db.GetUserById(99999)

		assert.Nil(t, user)
	})
}

func TestUserData_AddUser(t *testing.T) {
	t.Run("valid user adds to database", func(t *testing.T) {
		validEmail := "test@user"

		db := NewMemoryUserDB()

		// user does not exist
		user := db.GetUserByEmail(validEmail)
		assert.Nil(t, user)

		// add user
		db.AddUser(DataUser{Email: validEmail})

		// user now exists
		user = db.GetUserByEmail(validEmail)

		assert.Equal(t, validEmail, user.Email)
	})

	t.Run("adding user that already exists returns error", func(t *testing.T) {
		validEmail := "test@user"

		db := NewMemoryUserDB()
		// verify user does not exist
		user := db.GetUserByEmail(validEmail)
		assert.Nil(t, user)

		// add user and verify no error
		user, err := db.AddUser(DataUser{Email: validEmail})
		assert.NotNil(t, user)
		assert.Nil(t, err)

		// add user again and verify error exists
		user, err = db.AddUser(DataUser{Email: validEmail})
		assert.NotNil(t, err)
	})
}

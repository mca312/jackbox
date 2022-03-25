package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mca312/jackbox/server/app"
	"github.com/stretchr/testify/assert"
)

// mockUserService mocks the service and allows the user to specifiy if an error will return
type mockUserService struct {
	returnError bool
}

func (m *mockUserService) Register(u app.User) error {
	var err error
	if m.returnError {
		err = errors.New("test error")
	}
	return err
}
func (m *mockUserService) Login(u app.User) (app.User, error) {
	var err error
	if m.returnError {
		err = errors.New("test error")
	}
	return app.User{}, err
}
func (m *mockUserService) GetUser(id int64) (app.User, error) {
	var err error
	if m.returnError {
		err = errors.New("test error")
	}
	return app.User{}, err
}

func newMockManager(returnError bool) *Manager {
	return &Manager{
		&mockUserService{
			returnError,
		},
	}
}

func TestHandlers_Register(t *testing.T) {
	t.Run("Success returns a 200", func(t *testing.T) {
		manager := newMockManager(false)
		r := setupRouter(manager)
		validUser := UserModel{Email: "valid@email"}
		validReq, _ := json.Marshal(validUser)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/register", strings.NewReader(string(validReq)))
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Invalid model returns a 400", func(t *testing.T) {
		manager := newMockManager(false)
		r := setupRouter(manager)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/register", strings.NewReader("{\"invalidModel\": \"notAProperty\"}"))
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Valid model with invalid email returns a 400", func(t *testing.T) {
		manager := newMockManager(false)
		r := setupRouter(manager)
		invalidEmail := UserModel{Email: "invalidemail"}
		validReq, _ := json.Marshal(invalidEmail)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/register", strings.NewReader(string(validReq)))
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Returns 500 if application service has an error", func(t *testing.T) {
		manager := newMockManager(true)
		r := setupRouter(manager)
		validEmail := UserModel{Email: "valid@email"}
		validReq, _ := json.Marshal(validEmail)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/register", strings.NewReader(string(validReq)))
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

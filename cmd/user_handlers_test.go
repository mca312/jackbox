package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mca312/jackbox/server/app"
	"github.com/stretchr/testify/assert"
)

func TestHandlers_Register(t *testing.T) {
	t.Run("Success returns a 200", func(t *testing.T) {
		manager := NewManager()
		r := setupRouter(manager)
		validUser := app.User{Email: "valid@email", Name: "validName"}
		validReq, _ := json.Marshal(validUser)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/register", strings.NewReader(string(validReq)))
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Invalid model returns a 400", func(t *testing.T) {
		manager := NewManager()
		r := setupRouter(manager)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/register", strings.NewReader("{\"invalidModel\": \"notAProperty\"}"))
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Valid model with invalid email returns a 400", func(t *testing.T) {
		manager := NewManager()
		r := setupRouter(manager)
		invalidEmail := app.User{Email: "invalidemail"}
		validReq, _ := json.Marshal(invalidEmail)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/register", strings.NewReader(string(validReq)))
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

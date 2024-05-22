package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/reyhanrazaby/dating-app/usecase/login"
	"github.com/reyhanrazaby/dating-app/usecase/signup"
	"github.com/stretchr/testify/assert"
)

var r http.Handler

func TestMain(m *testing.M) {
	r = GetRoutes()
	m.Run()
}

func Test_Success(t *testing.T) {
	// Register new user
	signUpRecorder := httptest.NewRecorder()
	signUpBody := []byte(`
		{
			"full_name": "Rey",
			"gender": "M",
			"email": "rey@mail.id",
			"password": "123",
			"date_birth": "22-12-1994",
			"location_lat": -6.8059399341534075,
			"location_lng": 106.95161161606342
		}
	`)
	signUpReq, _ := http.NewRequest("POST", signup.Path, bytes.NewBuffer(signUpBody))
	r.ServeHTTP(signUpRecorder, signUpReq)
	assert.Equal(t, http.StatusOK, signUpRecorder.Code)

	// Login with registered user
	loginRecorder := httptest.NewRecorder()
	loginBody := []byte(`
		{
			"email": "rey@mail.id",
			"password": "123"
		}
	`)
	loginReq, _ := http.NewRequest("POST", login.Path, bytes.NewBuffer(loginBody))
	r.ServeHTTP(loginRecorder, loginReq)
	assert.Equal(t, http.StatusOK, loginRecorder.Code)
}

func Test_WrongPassword(t *testing.T) {
	// Register new user
	signUpRecorder := httptest.NewRecorder()
	signUpBody := []byte(`
		{
			"full_name": "Rey 2",
			"gender": "F",
			"email": "rey2@mail.id",
			"password": "1234",
			"date_birth": "22-12-1994",
			"location_lat": -6.8059399341534075,
			"location_lng": 106.95161161606342
		}
	`)
	signUpReq, _ := http.NewRequest("POST", signup.Path, bytes.NewBuffer(signUpBody))
	r.ServeHTTP(signUpRecorder, signUpReq)
	assert.Equal(t, http.StatusOK, signUpRecorder.Code)

	// Login with wrong password
	loginRecorder := httptest.NewRecorder()
	loginBody := []byte(`
		{
			"email": "rey2@mail.id",
			"password": "5555"
		}
	`)
	loginReq, _ := http.NewRequest("POST", login.Path, bytes.NewBuffer(loginBody))
	r.ServeHTTP(loginRecorder, loginReq)
	assert.Equal(t, http.StatusBadRequest, loginRecorder.Code)
}
